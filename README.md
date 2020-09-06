# Inertia.js Golang Adapter (Fiber)

This is a Inertia.js server-side adapter based on inertia-laravel, but for Fiber Framework.

## Installation 

#### What do you need ?

- Fiber Middleware
```go
	app := fiber.New()
    // AssetsPath is the path of your assets
	app.Use(inertia.New(inertia.Config{
		AssetsPath: "./public",
	}))
```
- Use `Render` method

```go
	inertia.Render(c,
		"App",
		inertia.Map{
		"Hello" : "World",
		},
	)
```
#### Install Client Side

Use [official documentation](https://inertiajs.com/client-side-setup) to install client side.

#### Example:

```go
package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"github.com/theArtechnology/fiber-inertia/inertia"
	"log"
)

func main() {
	engine := html.New("./public", ".html")
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Static("/assets", "./public/build")

	app.Use(inertia.New(inertia.Config{
		AssetsPath: "./public",
	}))

	app.Get("/hi", hello)
	app.Get("/bonjour", world)
	fmt.Println("Server started")
	log.Fatal(app.Listen(3001))
}

func hello (c *fiber.Ctx) {
	inertia.Render(c,
		"Main", // Will render component named as Main
		inertia.Map{
		    "Hi-EN" : "Hello World",
		},
	)
}

func world (c *fiber.Ctx) {
	inertia.Render(c,
		"sub/Users", // Will render component in a subdirectory
		inertia.Map{
		    "Hi-FR" : "Bonjour tout le monde",
		},
	)
}

```