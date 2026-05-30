package health

import (
	"context"
	"testing"
)

func TestCheckHandlerReturnServiceStatus(t *testing.T) {
	handler := NewCheckHandler("auth-service")

	res, err := handler.Handle(context.Background(), &CheckRequest{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.Service != "auth-service" {
		t.Fatalf("expected service auth-service, got %q", res.Service)
	}

	if res.Status != "ok" {
		t.Fatalf("expected status ok, got %q", res.Status)
	}

}
