package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/handlers"
)

func RouteAPP(app *fiber.App){
	api := app.Group("/api")

	api.Get("/weather/current", handlers.GetWeather)
	api.Get("/worship/current", handlers.GetWorship)
}