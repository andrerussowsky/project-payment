package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestSetupRoutesWithAuthentication(t *testing.T) {
    app := fiber.New()

    ctx := context.Background()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        t.Fatal(err)
    }

    collection := client.Database("test")

    app.Post("/login", Login)
    app.Use(AuthenticateRequest())
    SetupRoutes(app, collection)

    // Get authentication token
    req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"john","password":"doe"}`))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    var result map[string]string
    json.NewDecoder(resp.Body).Decode(&result)
    token, exists := result["token"]
    if !exists {
        t.Fatal("No token received")
    }

    // Test /process_payment route
    req = httptest.NewRequest("POST", "/process_payment", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, _ = app.Test(req)
    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

    // Test /get_payment_details/:transaction_id route
    req = httptest.NewRequest("GET", "/get_payment_details/abc", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, _ = app.Test(req)
    assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

    // Test /process_refund route
    req = httptest.NewRequest("POST", "/process_refund", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, _ = app.Test(req)
    assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
