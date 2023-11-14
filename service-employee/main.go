package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Greeting struct to represent the response message
type Greeting struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/api/greet", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		response := Greeting{Message: fmt.Sprintf("Hello, %s!", name)}

		return c.JSON(response)
	})

	// Start the first service on port 3002
	port := 3002
	fmt.Printf("Service employee is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service employee: %v\n", err)
	}

}
