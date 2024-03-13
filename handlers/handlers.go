package handlers

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(app *fiber.App, collection *mongo.Database) {
    // Process payment route
    app.Post("/process_payment", AuditTrail(collection.Collection("audit"), "process_payment"), func(c *fiber.Ctx) error {
        return ProcessPayment(c, collection)
    })
    // Get payment details route
    app.Get("/get_payment_details/:transaction_id", AuditTrail(collection.Collection("audit"), "get_payment_details"), func(c *fiber.Ctx) error {
        return GetPaymentDetails(c, collection)
    })
    // Process refund route
    app.Post("/process_refund", AuditTrail(collection.Collection("audit"), "process_refund"), func(c *fiber.Ctx) error {
        return ProcessRefund(c, collection)
    })
}

func AuditTrail(collection *mongo.Collection, action string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Log the request
        requestBody := string(c.Body())
        if c.Method() == "GET" {
            requestBody = c.Params("transaction_id")
        }
        auditTrail(c, collection, action+"_request", requestBody)

        // Call the next middleware in the stack
        err := c.Next()

        // Log the response
        auditTrail(c, collection, action+"_response", fmt.Sprintf("StatusCode:%v Body:%v", c.Response().Header.StatusCode(), string(c.Response().Body())))

        return err
    }
}

// Function to record an audit trail
func auditTrail(c *fiber.Ctx, auditCollection *mongo.Collection, action string, details interface{}) {
    // Extract user from JWT
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)

    // Create a new audit record
    auditRecord := bson.M{
        "action":   action,
        "details":  details,
        "datetime": time.Now(),
        "user":     name,
    }

    // Insert the audit record
    _, err := auditCollection.InsertOne(context.Background(), auditRecord)
    if err != nil {
        log.Println("Failed to record audit trail:", err)
    }
}
