package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()


	app.Get("/:foo", func(c*fiber.Ctx) error {
		result := c.Params("foo")
		return c.SendString(result)
	})

	app.Listen(":3000")
}