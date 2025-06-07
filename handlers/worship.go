package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/services"
)

func GetWorship(c *fiber.Ctx) error {
	laStr := c.Query("lat")
	lonStr := c.Query("lon")

	lat, err1 := strconv.ParseFloat(laStr, 64)
	lon, err2 := strconv.ParseFloat(lonStr, 64)

	if err1 != nil || err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid la dan lon",
		})
	}

	data, err := services.GetWorshipService(lat, lon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}
