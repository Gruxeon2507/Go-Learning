package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// // HTTP methods
	// func (app *App) Get(path string, handlers ...Handler) Router
	// func (app *App) Head(path string, handlers ...Handler) Router
	// func (app *App) Post(path string, handlers ...Handler) Router
	// func (app *App) Put(path string, handlers ...Handler) Router
	// func (app *App) Delete(path string, handlers ...Handler) Router
	// func (app *App) Connect(path string, handlers ...Handler) Router
	// func (app *App) Options(path string, handlers ...Handler) Router
	// func (app *App) Trace(path string, handlers ...Handler) Router
	// func (app *App) Patch(path string, handlers ...Handler) Router

	// // Add allows you to specifiy a method as value
	// func (app *App) Add(method, path string, handlers ...Handler) Router

	// // All will register the route on all HTTP methods
	// // Almost the same as app.Use but not bound to prefixes
	// func (app *App) All(path string, handlers ...Handler) Router

	//Simple GET Handler
	app.Get("/api/test", func(c *fiber.Ctx) error {
		return c.SendString("GET Request")
	})

	//Simple POST Handler
	app.Post("/api/test", func(c *fiber.Ctx) error {
		return c.SendString("POST Request")
	})

	//Use can be used for middleware packages and prefix catchers. These routes will match
	//the begining of each path
	//Signature
	//func (app *App) Use(args ...interface{}) Router

	//Match any request
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("First")
		return c.Next()
	})

	//Next() is called,it executes the next method in the stack that matches the current routes
	app.Get("/hi", func(c *fiber.Ctx) error {
		fmt.Println("Second")
		return c.SendString("here")
	})

	//Match any request starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	//Match any request starting with /api or /home
	app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	//Attach multiple handlers
	app.Use("/api", func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "10")
		return c.Next()
	}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	//Paths
	//Route paths, combine with a request method, define the endponts at which request can be made.
	//Route paths can be string or string patterns

	//This route path will match requests to te root route, "/":
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	//This route path will match requets to "/about":
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})

	//This route path will match request to "/random.txt":
	app.Get("/random.txt", func(c *fiber.Ctx) error {
		return c.SendString("random.txt")
	})

	//Parameters
	app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("title"))
		return c.SendString("User: " + c.Params("name") + "\n" + "Title: " + c.Params("title"))
	})

	//Plus - greedy - optional -> required to have atleast one more segment after that
	app.Get("/user/+", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("+"))
	})

	//Willcard - greedy - optional -> can have any number of segment after it
	app.Get("/user/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	//This route path will match request to "/v1/some/resource/name:customVerb", since the parameter character is escaped
	//The \\ make the :customeVerb count as a normal string
	app.Get("/v1/some/resource/name\\:customVerb", func(c *fiber.Ctx) error {
		return c.SendString("Hello:")
	})

	// the hyphen (-) and the dot (.) can be used along with route parameters for useful purposes.
	// http://localhost:3000/plantae/prunus.persica
	app.Get("/plantae/:genus.:species", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
		return nil // prunus.persica
	})

	// http://localhost:3000/flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
		return nil // LAX-SFO
	})

	//Reconize the introductory parameter
	// http://localhost:3000/shop/product/color:blue/size:xs
	app.Get("/shop/product/color::color/size::size", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s:%s\n", c.Params("color"), c.Params("size"))
		return nil // blue:xs
	})

	//useful example:
	// GET /@v1
	// Params: "sign" -> "@", "param" -> "v1"
	app.Get("/:sign:param", func(c *fiber.Ctx) error {
		return c.SendString("@v1")
	})

	// GET /api-v1
	// Params: "name" -> "v1"
	app.Get("/api-:name", func(c *fiber.Ctx) error {
		return c.SendString("/api-v1")
	})

	// GET /customer/v1/cart/proxy
	// Params: "*1" -> "customer/", "*2" -> "/cart"
	app.Get("/*v1*/proxy", func(c *fiber.Ctx) error {
		return c.SendString("/customer/v1/cart/proxy")
	})

	// GET /v1/brand/4/shop/blue/xs
	// Params: "*1" -> "brand/4", "*2" -> "blue/xs"
	app.Get("/v1/*/shop/*", func(c *fiber.Ctx) error {
		return c.SendString("/v1/brand/4/shop/blue/xs")
	})

	//Constraints
	//Route constraints execute when a match has occured to the incoming URL
	//Constraints are listed in: https://docs.gofiber.io/guide/routing

	//Example:
	//Single Constraint: Only allow params test is number with value greater or equal to 5
	app.Get("/:test<min(5)>", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("test"))
	})

	// curl -X GET http://localhost:3000/12
	// 12

	// curl -X GET http://localhost:3000/1
	// Cannot GET /1

	//Multiple Constraints: value must greater than 100 and has at max 5 character
	//Use the ";" for the multipe constraints
	app.Get("/:test<min(100);maxLen(5)>", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("test"))
	})

	// curl -X GET http://localhost:3000/120000
	// Cannot GET /120000

	// curl -X GET http://localhost:3000/1
	// Cannot GET /1

	// curl -X GET http://localhost:3000/250
	// 250

	//Regex Constraint: Date must be in format of yyyy-mm-dd
	//Should use "\\" before routing-specific character when using date time constrain to avoid wrong parsing
	app.Get("/:date<regex(\\d{4}-\\d{2}-\\d{2})}>", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("date"))
	})

	// curl -X GET http://localhost:3000/125
	// Cannot GET /125

	// curl -X GET http://localhost:3000/test
	// Cannot GET /test

	// curl -X GET http://localhost:3000/2022-08-27
	// 2022-08-27

	//Optional Parameter Example:
	app.Get("/:test<int>?", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("test"))
	})
	// curl -X GET http://localhost:3000/42
	// 42
	// curl -X GET http://localhost:3000/
	//
	// curl -X GET http://localhost:3000/7.0
	// Cannot GET /7.0

	//Middleware
	//Middleware in web development refers to a series of functions or code that are executed in a sequential order before or after the main request handler of a web application. Middleware functions can intercept, process, modify, or augment incoming requests and outgoing responses. They play a crucial role in enhancing the functionality, security, and manageability of web applications.
	//As mention above, the Next() is a fiber router fucntion when called, execute the next function that matched the current route
	//Example:
	app.Use(func(c *fiber.Ctx) error {
		// Set a custom header on all responses:
		c.Set("X-Custom-Header", "Hello, World")

		// Go to next middleware:
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//Grouping
	//Using to organize your routes using Group()

	//Define a middleware
	middleware := func(c *fiber.Ctx) error {
		return c.Next()
	}
	//Create group for the api route
	api := app.Group("/api", middleware) // /api

	//Create a sub-group for the /v1 route
	v1 := api.Group("v1", middleware) // api/v1
	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("API v1 List")
	}) // /api/v1/list

	v1.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("API v1 User")
	}) // /api/v1/user

	//Create a sub-group for the /v2 route
	v2 := api.Group("v2", middleware) ///api/v2

	v2.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("API v2 List")
	}) // /api/v2/list

	v2.Get("user", func(c *fiber.Ctx) error {
		return c.SendString("API v2 User")
	}) // /api/v2/user
	app.Listen(":3000")
}
