package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/ticket"
	tStatus "event_system/internal/domain/ticketstatus"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	tests := []struct {
		name    string
		t       *ticket.Ticket
		wantErr bool
	}{
		{
			name: "Create ticket successfully",
			t: ticket.New(
				"test_qr_code",
				fixtures.TestRegistrations["registration1"].ID,
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			// start transaction
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("failed to begin transaction: %v", err)
			}

			defer func() {
				// rollback transaction
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("failed to rollback transaction: %v", err)
				}
			}()

			tr := NewTicketRepository()

			repoErr := tr.Create(ctx, tt.t)
			if tt.wantErr {
				assert.Error(t, repoErr)
				return
			}

			assert.NoError(t, err)

			var ticketModel models.Ticket
			if err := tx.First(&ticketModel, "qr_code = ?", tt.t.QRCode()).Error; err != nil {
				assert.Fail(t, "failed to get ticket model")
			}

			assert.Equal(t, tt.t.ID(), ticketModel.ID)
			assert.Equal(t, tt.t.QRCode(), ticketModel.QRCode)
			assert.Equal(t, tt.t.TicketStatus().Value().String(), ticketModel.Status)
			assert.Equal(t, tt.t.IssuedAt().Unix(), ticketModel.IssuedAt.Unix())
			assert.Equal(t, tt.t.RegistrationID(), ticketModel.RegistrationID)
		})
	}
}

func TestFetchTicketByQRCode(t *testing.T) {
	tests := []struct {
		name    string
		qrCode  string
		wantErr bool
	}{
		{
			name:    "Fetch ticket by qr code successfully",
			qrCode:  fixtures.TestTickets["ticket1"].QRCode,
			wantErr: false,
		},
		{
			name:    "Fetch ticket by non-existent qr code",
			qrCode:  "non_existent_qr_code",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewTicketRepository()

			ticket, err := r.FetchByQRCode(ctx, tt.qrCode)

			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)

			// Not found
			if ticket == nil {
				return
			}

			assert.Equal(t, fixtures.TestTickets["ticket1"].ID, ticket.ID())
			assert.Equal(t, fixtures.TestTickets["ticket1"].QRCode, ticket.QRCode())
			assert.Equal(t, fixtures.TestTickets["ticket1"].Status, ticket.TicketStatus().Value().String())
			assert.Equal(t, fixtures.TestTickets["ticket1"].IssuedAt.Unix(), ticket.IssuedAt().Unix())
			assert.Equal(t, fixtures.TestTickets["ticket1"].UsedAt.Unix(), ticket.UsedAt().Unix())
			assert.Equal(t, fixtures.TestTickets["ticket1"].RegistrationID, ticket.RegistrationID())
		})
	}
}

func TestUpdateTicket(t *testing.T) {
	tests := []struct {
		name    string
		t       *ticket.Ticket
		wantErr bool
	}{
		{
			name: "Update ticket successfully",
			t: ticket.Reconstruct(
				fixtures.TestTickets["ticket1"].ID,
				fixtures.TestTickets["ticket1"].QRCode,
				fixtures.TestTickets["ticket1"].IssuedAt,
				time.Now(),
				tStatus.Reconstruct(tStatus.Used),
				fixtures.TestTickets["ticket1"].RegistrationID,
			),
			wantErr: false,
		},
		{
			name: "Update ticket with non-existent ticket",
			t: ticket.Reconstruct(
				uuid.New(), // non-existent ticket
				fixtures.TestTickets["ticket1"].QRCode,
				fixtures.TestTickets["ticket1"].IssuedAt,
				time.Now(),
				tStatus.Reconstruct(tStatus.Used),
				fixtures.TestRegistrations["registration1"].ID,
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			// start transaction
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("failed to begin transaction: %v", err)
			}

			defer func() {
				// rollback transaction
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("failed to rollback transaction: %v", err)
				}
			}()

			tr := NewTicketRepository()

			repoErr := tr.Update(ctx, tt.t)
			if tt.wantErr {
				assert.Error(t, repoErr)
				return
			}

			assert.NoError(t, err)

			var ticketModel models.Ticket
			if err := tx.Where("id = ?", tt.t.ID()).First(&ticketModel).Error; err != nil {
				t.Fatalf("failed to get ticket model: %v", err)
			}

			assert.Equal(t, tt.t.ID(), ticketModel.ID)
			assert.Equal(t, tt.t.QRCode(), ticketModel.QRCode)
			assert.Equal(t, tt.t.TicketStatus().Value().String(), ticketModel.Status)
			assert.Equal(t, tt.t.IssuedAt().Unix(), ticketModel.IssuedAt.Unix())
			assert.Equal(t, tt.t.UsedAt().Unix(), ticketModel.UsedAt.Unix())
			assert.Equal(t, tt.t.RegistrationID(), ticketModel.RegistrationID)
		})
	}
}
