package main

import (
	"log"
	"study/go-fiber/config"
	"study/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()	
	dbConf := config.NewDatabaseConfig()
    log.Println((dbConf)) 
	app := fiber.New()
	home.NewHandler(app)
	app.Listen(":3000")

}