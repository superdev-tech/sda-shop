package config

import (
	"errors"
	"testing"
)

func TestValidateMissingDatabaseURL(t *testing.T) {
	err := Config{Port: "8080"}.Validate()
	if !errors.Is(err, ErrMissingDatabaseURL) {
		t.Fatalf("want ErrMissingDatabaseURL, got %v", err)
	}
}

func TestValidateOK(t *testing.T) {
	if err := (Config{DatabaseURL: "postgresql://x"}).Validate(); err != nil {
		t.Fatalf("want nil, got %v", err)
	}
}
