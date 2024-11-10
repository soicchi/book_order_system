package logging

import (
	"context"
	"log/slog"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (m *MockLogger) Debug(msg string, attrs ...any) {
	m.Called(msg, attrs)
}

func (m *MockLogger) Info(msg string, attrs ...any) {
	m.Called(msg, attrs)
}

func (m *MockLogger) Warn(msg string, attrs ...any) {
	m.Called(msg, attrs)
}

func (m *MockLogger) Error(msg string, attrs ...any) {
	m.Called(msg, attrs)
}

func (m *MockLogger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	m.Called(ctx, level, msg, attrs)
}
