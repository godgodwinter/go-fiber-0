package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/godgodwinter/go-fiber-0/app/controllers"
	"github.com/godgodwinter/go-fiber-0/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ExampleStruct struct {
	Name string
	Desc string
}

func TestingIndex(app *fiber.App) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app_version := os.Getenv("APP_VERSION")
	api_version := app.Group("/api/" + app_version)
	app_routes := api_version.Group("/testing")
	app_routes.Use(middleware.AuthMiddleware)
	// Menambahkan middleware secara berurutan !contoh penerapan beberapa middle ware
	// app_routes.Use(middleware.CheckAuthMiddleware)
	// app_routes.Use(middleware.CheckAccessMiddleware)
	app_routes.Get("/index", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})
	app_routes.Get("/nama/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	app_routes.Get("/json/:name", func(c *fiber.Ctx) error {
		data := ExampleStruct{
			Name: c.Params("name"),
			Desc: "Ini desc",
		}
		return c.JSON(data)
	})
	// app_routes.Get("/json/:name", func UserRoute(c *fiber.App) error {
	// 	app.Post("/user", controllers.CreateUser) //add this
	// })

	app_routes.Get("/example/:userId", controllers.GetAUser)

}

// controllers.GetAUser(c * fiber.App)
// func ExampleRouter(app *fiber.App) {
// 	app.Get("/example", controllers.GetAUser) //add this
// }
