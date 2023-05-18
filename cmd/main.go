package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/godgodwinter/go-fiber-0/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor" //digunakan untuk menangani panic pada saat runtime. Panic terjadi ketika terjadi kesalahan yang tidak dapat ditangani pada saat aplikasi berjalan, dan biasanya menyebabkan program berhenti secara tiba-tiba. go recover digunakan untuk menangani panic sehingga aplikasi tidak berhenti dengan sendirinya.
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	strVar := os.Getenv("APP_PORT")
	app_port, err := strconv.Atoi(strVar)

	if err != nil {
		// handle error
		log.Fatal("Error env port tidak ditemukan")
	}

	// !Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		Network:       "tcp4",
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})

	type SomeStruct struct {
		Name string
		Port int
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: fmt.Sprintf("This is go starter on Server port: %d", app_port),
			Port: app_port,
		}

		return c.JSON(data)
	})

	app.Use(recover.New())

	app.Get("/err", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})

	app.Get("/err2", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON("errors")
	})

	//   !routes
	routes.TestingIndex(app)
	// routes.ExampleRouter(app)

	//   !monitor
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// Menentukan alamat dan port
	addr := fmt.Sprintf("0.0.0.0:%d", app_port)

	// Mulai server HTTP
	app.Listen(addr)

}
