/*
 * Route: Shorten URL
 * Generate an UUID to create a Shorten URL from the original URL
 * Also, have additional verifications (e.g. Invalid URL, Invalid Domain, JSON Parse Error, etc)
 */

package routes

import (
	"os"
	"github.com/nbhalala/codingtask/database"
	"github.com/nbhalala/codingtask/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// API REQUEST
type request struct {
	URL		string		`json:"url"`
	ShortURL	string		`json:"short"`
}

// API RESPONSE
type response struct {
	URL		string		`json:"url"`
	ShortURL	string		`json:"short"`
}

func urlShorten(c *fiber.Ctx) error {

	body := new(request)

	//Verify the JSON Parse
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON."})
	}

	//Verify the URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL."})
	}

	//Verify the Domain
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Invalid Domain."})
	}

	// Generate Unique ID
	var id string
	id = uuid.New().String()[:6]
	
	r := database.CreateClient(0)
	defer r.Close()

	// Add verification : Shortne URL is already exists/in-use
	// TBD	
	
	// Add verification : Server Connection
	err := r.Set(database.Ctx, id, body.URL, 24*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Unable to connect to server"})
	}
		
	// API Response
	resp := response{
		URL:		body.URL,
		ShortURL:	"",
	}

	resp.ShortURL = os.Getenv("APP_DOMAIN") + "/" + id

	// StatusOK: 200
	return c.Status(fiber.StatusOK).JSON(resp)
}
