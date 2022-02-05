package main

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

import (
	"github.com/Allan-Nava/go-ddd/config"
	"github.com/Allan-Nava/go-ddd/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	// setup fiber stuff
	f := fiber.New()
	f.Use(logger.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS, PUT, PATCH, POST, DELETE",
	}))
	f.Use(recover.New())
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	utils.SetupEnv() //dotenv
	config.SetEnvConfig()
	//
}
