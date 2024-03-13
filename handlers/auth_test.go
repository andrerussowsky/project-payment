package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    app := fiber.New()

    app.Post("/login", Login)

    // Test with correct credentials
    req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"john","password":"doe"}`))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // Test with incorrect credentials
    req = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"wrong","password":"wrong"}`))
    req.Header.Set("Content-Type", "application/json")
    resp, _ = app.Test(req)

    assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
