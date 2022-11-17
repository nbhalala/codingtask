/*
 * Route: Resolve URL
 * Verify the generated Shorten URL does exist in the database.
 */

package routes

import (
	"github.com/nbhalala/codingtask/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func urlResolve(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Short not found in the Database."})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
