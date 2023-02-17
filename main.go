package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/Calculadora-201901073", func (c *fiber.Ctx) error {

		type Request struct {
			Num1 int
			Num2 int
		}

		var body Request
		c.BodyParser(&body)

		suma := body.Num1 + body.Num2

		return c.JSON(fiber.Map{
			"suma": suma,
		})
	})

	app.Listen(":3000")
}