package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/models"
	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/services"
)

func GetWeather(c *fiber.Ctx) error {
	latParam := c.Query("lat")
	lonParam := c.Query("lon")

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid latitude"})
	}
	lon, err := strconv.ParseFloat(lonParam, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid longitude"})
	}

	weather, err := services.GetWeatherService(lat, lon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	//apply formated repson
	goodFormated := models.FormatWeatherResponse(weather)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": goodFormated})
}