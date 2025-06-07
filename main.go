package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/routes"
)

func main() {
	
	godotenv.Load()
	app := fiber.New()

	routes.RouteAPP(app)

	log.Fatal(app.Listen(":3000"))

}