package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockEventFactory struct {
	mock.Mock
}

func NewMockEventFactory() *MockEventFactory {
	return &MockEventFactory{}
}

func (m *MockEventFactory) NewEvent(
	ctx echo.Context,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	userID uuid.UUID,
	venueID uuid.UUID,
) (*Event, error) {
	args := m.Called(ctx, title, description, startDate, endDate, userID, venueID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Event), args.Error(1)
}

type MockEventUpdater struct {
	mock.Mock
}

func NewMockEventUpdater() *MockEventUpdater {
	return &MockEventUpdater{}
}

func (m *MockEventUpdater) UpdateEvent(
	ctx echo.Context,
	eventID uuid.UUID,
	title string,
	description string,
	startDate time.Time,
	endDate time.Time,
	createdBy uuid.UUID,
	venueID uuid.UUID,
) (*Event, error) {
	args := m.Called(ctx, eventID, title, description, startDate, endDate, createdBy, venueID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Event), args.Error(1)
}
