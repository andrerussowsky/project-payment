package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

// BankSimulatorResponse is the response from the bank simulator
type BankSimulatorResponse struct {
    Status string `json:"status"`
}

// ProcessPaymentSimulator simulates processing a payment
func processPaymentSimulator(c *fiber.Ctx) error {
    response := BankSimulatorResponse{Status: "Approved"}
    return c.JSON(response)
}

// RefundPaymentSimulator simulates refunding a payment
func refundPaymentSimulator(c *fiber.Ctx) error {
    response := BankSimulatorResponse{Status: "Approved"}
    return c.JSON(response)
}

func main() {
    // Create a new Fiber application
    app := fiber.New()

    // Process payment route
    app.Post("/process_payment", processPaymentSimulator)

    // Refund payment route
    app.Post("/refund_payment", refundPaymentSimulator)

    // Start the server on port 4000
    log.Fatal(app.Listen(":4000"))
}
