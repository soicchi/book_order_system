package repository

import (
	"errors"
	"fmt"

	"event_system/internal/domain/ticket"
	tStatus "event_system/internal/domain/ticketstatus"
	errs "event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TicketRepository struct{}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (tr *TicketRepository) Create(ctx echo.Context, t *ticket.Ticket) error {
	db := database.GetDB(ctx)

	ticketModel := models.Ticket{
		ID:             t.ID(),
		QRCode:         t.QRCode(),
		Status:         t.TicketStatus().Value().String(),
		IssuedAt:       t.IssuedAt(),
		RegistrationID: t.RegistrationID(),
	}

	err := db.Create(&ticketModel).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.New(
			fmt.Errorf("qr code %s already exists: %w", t.QRCode(), err),
			errs.AlreadyExistError,
			errs.WithField("QRCode"),
		)
	}

	if err != nil {
		return errs.New(
			fmt.Errorf("failed to create ticket: %w", err),
			errs.UnexpectedError,
		)
	}

	return nil
}

func (tr *TicketRepository) FetchByQRCode(ctx echo.Context, qrCode string) (*ticket.Ticket, error) {
	db := database.GetDB(ctx)

	var ticketModel models.Ticket
	err := db.First(&ticketModel, "qr_code = ?", qrCode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch ticket by qr code: %w", err),
			errs.UnexpectedError,
		)
	}

	t := ticket.Reconstruct(
		ticketModel.ID,
		ticketModel.QRCode,
		ticketModel.IssuedAt,
		ticketModel.UsedAt,
		tStatus.Reconstruct(tStatus.FromString(ticketModel.Status)),
		ticketModel.RegistrationID,
	)

	return t, nil
}

func (tr *TicketRepository) Update(ctx echo.Context, t *ticket.Ticket) error {
	db := database.GetDB(ctx)

	result := db.Model(&models.Ticket{}).Where("id = ?", t.ID()).Updates(models.Ticket{
		UsedAt: t.UsedAt(),
		Status: t.TicketStatus().Value().String(),
	})
	if result.Error != nil {
		return errs.New(
			fmt.Errorf("failed to update ticket: %w", result.Error),
			errs.UnexpectedError,
		)
	}

	if result.RowsAffected == 0 {
		return errs.New(
			fmt.Errorf("ticket with id %s not found", t.ID()),
			errs.NotFoundError,
			errs.WithField("TicketID"),
		)
	}

	return nil
}
