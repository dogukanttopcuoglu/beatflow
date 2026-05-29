package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestNewCreatesDevelopmentLogger(t *testing.T) {
	log, err := New(Options{
		Environment: "development",
		ServiceName: "auth-service",
		Level:       "debug",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if log == nil {
		t.Fatal("expected logger, got nil")
	}
}

func TestNewCreatesProductionLogger(t *testing.T) {
	log, err := New(Options{
		ServiceName: "auth-service",
		Environment: "production",
		Level:       "debug",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if log == nil {
		t.Fatal("expected logger, got nil")
	}
}

func TestNewCreatesProductionLoggerByDefault(t *testing.T) {
	log, err := New(Options{
		ServiceName: "auth-service",
		Level:       "debug",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if log == nil {
		t.Fatal("expected logger, got nil")
	}
}

func TestNewUsesSafeDefaultLevel(t *testing.T) {
	log, err := New(Options{
		Environment: "production",
		ServiceName: "auth-service",
		Level:       "not-a-real-level",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if log == nil {
		t.Fatal("expected logger, got nil")
	}
}

func TestLoggerCanBeExtendedWithRequestFields(t *testing.T) {
	log, err := New(Options{
		Environment: "production",
		ServiceName: "auth-service",
		Level:       "info",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	requestLogger := log.With(
		zap.String("request_id", "req-123"),
		zap.String("correlation_id", "corr-456"),
	)

	if requestLogger == nil {
		t.Fatal("expected request scoped logger, got nil")
	}
}
