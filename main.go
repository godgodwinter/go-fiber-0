package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

// !Custom config
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    Network: "tcp4",
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1",
})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("home tes")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})
	
	type SomeStruct struct {
		Name string
		Age  uint8
	  }
	  
	  app.Get("/json", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
		  Name: "Grame",
		  Age:  20,
		}
	  
		return c.JSON(data)
	  })
	  
	  app.Use(recover.New())

	  app.Get("/err", func(c *fiber.Ctx) error {
		  panic("This panic is caught by fiber")
	  })

	  app.Get("/err2",func(c *fiber.Ctx) error{
		return c.Status(fiber.StatusBadRequest).JSON("errors")
	  })

	//   !monitor
    app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

    // Menentukan alamat dan port
    addr := "0.0.0.0:8003"
	
    // Mulai server HTTP
    app.Listen(addr)

}