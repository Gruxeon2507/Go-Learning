package main

import fiber "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Respond with "Hello, World!" on root path, "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//Basic Routing:
	//Function Signature:
	//app.Method(path string, ...func(c *fiber.Ctx) error)
	//app is an instance of Fiber
	// Method is an HTTP request method: GET, PUT, POST, etc.
	// path is a virtual path on the server
	// func(*fiber.Ctx) error is a callback function containing the Context executed when the route is matched

	//Path Parameters:
	app.Get("param/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	//Optional parameter:
	app.Get("oparam/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
		}
		return c.SendString("Hello User")
	})

	//Wildcards
	app.Get("api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	})

	//Static files
	//app.Static(prefix, root string, config ...Static)
	app.Static("static/", "./public")

	//App
	//Return the *App reference so you could easily acces all application settings
	//func (c *Ctx) App() *App
	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	app.Listen(":3000")
}
