package tests

import (
	"testing"
	"time"

	"example.com/learn_web_servers/internal/auth"
	"github.com/google/uuid"
)

func TestFailCreateHash(t *testing.T) {
	password := "12345"
	_, err := auth.HashPassword(password)
	if err != nil {
		t.Errorf("Unable to create hashed PW: %v", err)
	}
}

func TestHashMatch(t *testing.T) {
	password := "12345"
	hashed_password, err := auth.HashPassword(password)
	if err != nil {
		t.Errorf("Unable to create hashed PW: %v", err)
	}
	match, err := auth.CheckPasswordHash(password, hashed_password)
	if err != nil {
		t.Errorf("Unable to run password match")
	}
	if !match {
		t.Errorf("Passwords do not match")
	}
}

func TestBadHashMatch(t *testing.T) {
	password := "12345"
	hashed_password, err := auth.HashPassword(password)
	if err != nil {
		t.Errorf("Unable to create hashed PW: %v", err)
	}
	match, err := auth.CheckPasswordHash("54321", hashed_password)
	if err != nil {
		t.Errorf("Unable to run password match")
	}
	if match {
		t.Errorf("Passwords match when they should not!")
	}
}

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := auth.MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	gotUserID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}

	if gotUserID != userID {
		t.Fatalf("expected user ID %v, got %v", userID, gotUserID)
	}
}

func TestValidateJWTWrongSecret(t *testing.T) {
	userID := uuid.New()

	token, err := auth.MakeJWT(userID, "correct-secret", time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = auth.ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Fatal("expected error for wrong secret")
	}
}

func TestValidateJWTExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := auth.MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = auth.ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected error for expired token")
	}
}
