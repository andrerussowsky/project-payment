package models

var (
    // PaymentApproved is the status for an approved payment
    PaymentApproved = "Approved"
    // PaymentDeclined is the status for a declined payment
    PaymentDeclined = "Declined"
    // PaymentRefunded is the status for a refunded payment
    PaymentRefunded = "Refunded"
)

// PaymentRequest is the request to process a payment
type PaymentRequest struct {
    MerchantID string  `json:"merchant_id" validate:"required"`
    Amount     float64 `json:"amount" validate:"required,gte=0"`
    CardNumber string  `json:"card_number" validate:"required,len=16"`
    CardExpiry string  `json:"card_expiry" validate:"required"`
    CVV        string  `json:"cvv" validate:"required,len=3"`
}

// Payment is the details of a payment
type Payment struct {
    MerchantID    string  `json:"merchant_id"`
    Amount        float64 `json:"amount"`
    CardNumber    string  `json:"card_number"`
    CardExpiry    string  `json:"card_expiry"`
    CVV           string  `json:"cvv"`
    TransactionID string  `json:"transaction_id"`
    Status        string  `json:"status"`
}

// RefundRequest is the request to process a refund
type RefundRequest struct {
    TransactionID string `json:"transaction_id" validate:"required"`
}
