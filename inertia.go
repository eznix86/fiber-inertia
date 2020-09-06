package inertia

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	// AssetsPath defines the path of assets, if something is changed in the target path,
	// inertiaJS will be notified
	AssetsPath string
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter 	func(*fiber.Ctx) bool
}


func New (config ...Config) func(*fiber.Ctx) {

	if len(config) == 0 {
		panic("No config found")
	}

	return func(c *fiber.Ctx) {

		if len(config[0].AssetsPath) == 0 {
			panic("AssetsPath is required")
		}

		hash := hashDir(config[0].AssetsPath)

		if c.Method() == "GET" && c.XHR() && c.Get("X-Inertia-Version", "1") != hash {
			c.Set("X-Inertia-Location", c.Path())
			_ = c.Status(fiber.StatusConflict).JSON(Map{})
			return
		}

		c.Set("X-Inertia-Version", hash)

		c.Next()
	}
}

type Map map[string]interface{}


func Render(c *fiber.Ctx, component string, props Map) {

	props = PartialReload(c, component, props)

	Display(c, component, props)
}

func Display(c *fiber.Ctx, component string, props Map) {
	data := map[string]interface{}{
		"component": component,
		"props":     props,
		"url":       c.OriginalURL(),
		"version":   c.Get("X-Inertia-Version", ""),
	}


	renderJSON, err := strconv.ParseBool(c.Get("X-Inertia", "false"))

	if err != nil {
		log.Fatal("X-Inertia not parsable")
	}

	if renderJSON && c.XHR() {
		JsonResponse(c, data)
		return
	}

	HtmlResponse(c, data)
}

func HtmlResponse(c *fiber.Ctx, data Map) {
	componentDataByte, _ := json.Marshal(data)

	_ = c.Render("index", fiber.Map{
		"Page": string(componentDataByte),
	})
}

func JsonResponse(c *fiber.Ctx, page Map) {
	jsonByte, _ := json.Marshal(page)
	_ = c.Status(fiber.StatusOK).JSON(string(jsonByte))
}

func PartialReload(c *fiber.Ctx, component string, props Map) Map {
	if c.Get("X-Inertia-Partial-Component", "/") == component {
		var newProps = make(Map)
		partials := strings.Split(c.Get("X-Inertia-Partial-Data", ""), ",")
		for key, _ := range props {
			for _, partial := range partials {
				if key == partial {
					newProps[partial] = props[key]
				}
			}
		}
	}
	return props
}
