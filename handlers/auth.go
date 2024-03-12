package handlers

import (
    "payment/models"
    "time"

    jwtware "github.com/gofiber/contrib/jwt"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

// Secret key used for JWT signing
var jwtSecret = []byte("your-secret-key")

// Function to login and return a JWT
func Login(c *fiber.Ctx) error {
    // Parse the request body
    p := new(models.LoginRequest)
    if err := c.BodyParser(p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Throws Unauthorized error
    if p.Username != "john" || p.Password != "doe" {
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    // Create the Claims
    claims := jwt.MapClaims{
        "name":  "John Doe",
        "admin": true,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Generate encoded token and send it as response.
    t, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(fiber.Map{"token": t})
}

// Middleware to authenticate requests
func AuthenticateRequest() func(*fiber.Ctx) error {
    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
    })
}
