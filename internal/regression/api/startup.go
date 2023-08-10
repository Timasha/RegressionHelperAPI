package api

import (
	"RegressionHelperAPI/internal/regression/api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Startup() {
	app := fiber.New()

	app.Get("/linear", handlers.LinearRegressionHandler)
	app.Get("/nonlinear2", handlers.NonLinear2RegressionHandler)

	log.Fatalln(app.Listen(":8080"))
}
