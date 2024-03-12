package handlers

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "payment/models"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// ProcessPayment processes a payment
func ProcessPayment(c *fiber.Ctx, collection *mongo.Database) error {
    // Parse the request body
    p := new(models.PaymentRequest)
    if err := c.BodyParser(p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Log the payment request
    auditTrail(collection.Collection("audit"), "process_payment", p)

    // Encode the payment request as JSON
    paymentJson, err := json.Marshal(p)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Send the payment request to the bank
    resp, err := http.Post("http://localhost:4000/process_payment", "application/json", bytes.NewBuffer(paymentJson))
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Decode the bank's response
    var bankResponse models.BankResponse
    err = json.NewDecoder(resp.Body).Decode(&bankResponse)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Check if the payment was approved
    if bankResponse.Status != "Approved" {
        return c.Status(400).SendString("Payment was not approved")
    }

    // Create a new payment object
    payment := models.Payment{
        MerchantID:    p.MerchantID,
        Amount:        p.Amount,
        CardNumber:    p.CardNumber,
        CardExpiry:    p.CardExpiry,
        CVV:           p.CVV,
        TransactionID: uuid.NewString(),
        Status:        models.PaymentApproved,
    }

    // Save the payment details to the database
    _, err = collection.Collection("payment").InsertOne(c.Context(), payment)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Return the payment details
    return c.JSON(payment)
}

// GetPaymentDetails retrieves the details of a payment
func GetPaymentDetails(c *fiber.Ctx, collection *mongo.Database) error {
    // Get the transaction ID from the URL
    transactionID := c.Params("transaction_id")
    fmt.Println(transactionID)

    // Log the request to get payment details
    auditTrail(collection.Collection("audit"), "get_payment_details", transactionID)

    // Find the payment details in the database
    var p models.Payment
    err := collection.Collection("payment").FindOne(c.Context(), bson.M{"transactionid": transactionID}).Decode(&p)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Return the payment details
    return c.JSON(p)
}

func ProcessRefund(c *fiber.Ctx, collection *mongo.Database) error {
    // Parse the request body
    r := new(models.RefundRequest)
    if err := c.BodyParser(r); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Log the refund request
    auditTrail(collection.Collection("audit"), "process_refund", r)

    // Find the payment details in the database
    var p models.Payment
    err := collection.Collection("payment").FindOne(c.Context(), bson.M{"transactionid": r.TransactionID}).Decode(&p)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    if p.Status != models.PaymentApproved {
        return c.Status(400).SendString("Payment was not approved")
    }

    // Encode the refund request as JSON
    refundJson, err := json.Marshal(p)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Send the refund request to the bank
    resp, err := http.Post("http://localhost:4000/refund_payment", "application/json", bytes.NewBuffer(refundJson))
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Decode the bank's response
    var bankResponse models.BankResponse
    err = json.NewDecoder(resp.Body).Decode(&bankResponse)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Check if the refund was approved
    if bankResponse.Status != "Approved" {
        return c.Status(400).SendString("Refund was not approved")
    }

    // Update the payment status to refunded
    _, err = collection.Collection("payment").UpdateOne(c.Context(), bson.M{"transactionid": r.TransactionID}, bson.M{"$set": bson.M{"status": models.PaymentRefunded}})
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    // Return the refund details
    return c.SendStatus(200)
}

// Function to record an audit trail
func auditTrail(auditCollection *mongo.Collection, action string, details interface{}) {
    // Create a new audit record
    auditRecord := bson.M{
        "action":   action,
        "details":  details,
        "datetime": time.Now(),
    }

    // Insert the audit record
    _, err := auditCollection.InsertOne(context.Background(), auditRecord)
    if err != nil {
        log.Println("Failed to record audit trail:", err)
    }
}
