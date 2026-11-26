package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

type fakeDB struct{ err error }

func (f fakeDB) Ping(context.Context) error { return f.err }

func TestHealthz(t *testing.T) {
	app := New("http://localhost:3000", fakeDB{})
	res, err := app.Test(httptest.NewRequest("GET", "/healthz", nil))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("want 200, got %d", res.StatusCode)
	}
}

func TestHealthzDatabaseDown(t *testing.T) {
	app := New("http://localhost:3000", fakeDB{err: errors.New("down")})
	res, _ := app.Test(httptest.NewRequest("GET", "/healthz", nil))
	if res.StatusCode != 503 {
		t.Fatalf("want 503, got %d", res.StatusCode)
	}
}

func TestUnknownAPIReturns404JSON(t *testing.T) {
	app := New("http://localhost:3000", fakeDB{})
	req := httptest.NewRequest("GET", "/api/products/featured", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 404 {
		t.Fatalf("want 404, got %d", res.StatusCode)
	}
	if got := res.Header.Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
		t.Fatalf("want CORS header for frontend origin, got %q", got)
	}
	var body ErrorBody
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Error.Code != "not_found" {
		t.Fatalf("want code not_found, got %q", body.Error.Code)
	}
}

func TestPreflightForJSONIsNotAllowedYet(t *testing.T) {
	app := New("http://localhost:3000", fakeDB{})
	req := httptest.NewRequest("OPTIONS", "/api/auth/login", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "POST")
	req.Header.Set("Access-Control-Request-Headers", "content-type")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if got := res.Header.Get("Access-Control-Allow-Headers"); got != "Accept" {
		t.Fatalf("want only Accept allowed until module 4D, got %q", got)
	}
}
