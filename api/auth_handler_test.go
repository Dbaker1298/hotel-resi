package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Dbaker1298/hotel-resi/db/fixtures"
	"github.com/gofiber/fiber/v2"
)

func TestAuthenticateWrongPassword(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	fixtures.AddUser(tdb.Store, "james", "bond", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@bond.com",
		Password: "supersecurepassnotcorrect",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var genResp genericResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		t.Fatal(err)
	}

	if genResp.Type != "error" {
		t.Fatalf("expected type to be <error>, got %s", genResp.Type)
	}

	if genResp.Msg != "invalid credentials" {
		t.Fatalf("expected msg to be <invalid credentials>, got %s", genResp.Msg)
	}
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	// insertedUser := insertTestUser(t, tdb.User)
	insertedUser := fixtures.AddUser(tdb.Store, "james", "bond", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@bond.com",
		Password: "james_bond",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var authResponse AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		t.Fatal(err)
	}

	if authResponse.Token == "" {
		t.Fatal("expected the JWT token to be returned in the response")
	}

	// Set the encrypted password to empty string so we can compare the structs
	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insertedUser, authResponse.User) {
		fmt.Println("insertedUser", insertedUser)
		fmt.Println("authResponse.User", authResponse.User)
		t.Fatalf("expected user to be %v, got %v", insertedUser, authResponse.User)
	}
}
