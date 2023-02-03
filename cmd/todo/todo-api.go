package main

/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */

import (
	"fmt"
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"

	"github.com/Allan-Nava/go-ddd/config"
	"github.com/Allan-Nava/go-ddd/database"
	"github.com/Allan-Nava/go-ddd/todo"
	"github.com/Allan-Nava/go-ddd/utils"
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
	// need to fix
	//database connections
	dbConn := database.InitDB()
	//stores
	todoStore := todo.MySqlTodoStore{DB: dbConn}
	//fmt.Printf("%v ", todoStore)
	//services
	//todoServ := todo.TodoService{Store: &todoStore}
	todoServ := todo.NewService(&todoStore)
	//fmt.Printf("todoServ: %v ", todoServ)
	//handlers
	todoHand := todo.TodoHandler{Service: &todoServ}

	//health check endpoint
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	//customer api
	todos := f.Group("/todos")
	todos.Get("/", nil, todoHand.GetAll)

	//swagger setup
	f.Get("/swagger/*", swagger.HandlerDefault) // default
	//
	f.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	fmt.Println("\nTodo CUSTOMER API")
	fmt.Printf("\nENV: %s", config.CONFIGURATION.AppEnv)
	fmt.Printf("\nRUNNING MODE: %s", config.CONFIGURATION.RunningMode)
	f.Listen("0.0.0.0:8080")
}
