package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(app *fiber.App, collection *mongo.Database) {
    // Process payment route
    app.Post("/process_payment", func(c *fiber.Ctx) error {
        return ProcessPayment(c, collection)
    })
    // Get payment details route
    app.Get("/get_payment_details/:transaction_id", func(c *fiber.Ctx) error {
        return GetPaymentDetails(c, collection)
    })
    // Process refund route
    app.Post("/process_refund", func(c *fiber.Ctx) error {
        return ProcessRefund(c, collection)
    })
}
