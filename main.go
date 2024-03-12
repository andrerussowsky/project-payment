package main

import (
    "context"
    "log"
    "payment/handlers"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Create a new context
    ctx := context.Background()

    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Set collections
    collection := client.Database("test")

    // Create a new Fiber instance
    app := fiber.New()

    // Use the logger middleware
    app.Use(logger.New())

    // Login route (unauthenticated)
    app.Post("/login", handlers.Login)

    // Middleware to authenticate requests
    app.Use(handlers.AuthenticateRequest())

    // Setup routes (authenticated)
    handlers.SetupRoutes(app, collection)

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
