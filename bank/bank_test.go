package main

import (
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func TestProcessPaymentSimulator(t *testing.T) {
    app := fiber.New()

    app.Post("/process_payment", processPaymentSimulator)

    req := httptest.NewRequest("POST", "/process_payment", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)

    body, _ := io.ReadAll(resp.Body)
    assert.Equal(t, `{"status":"Approved"}`, string(body))
}

func TestRefundPaymentSimulator(t *testing.T) {
    app := fiber.New()

    app.Post("/refund_payment", refundPaymentSimulator)

    req := httptest.NewRequest("POST", "/refund_payment", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)

    body, _ := io.ReadAll(resp.Body)
    assert.Equal(t, `{"status":"Approved"}`, string(body))
}
