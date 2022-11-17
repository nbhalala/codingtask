package routes

import(
	"os"
	"github.com/nbhalala/codingtask/database"
	"github.com/nbhalala/codingtask/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type request struct {
	URL		string		`json:"url"`
	ShortURL	string		`json:"short"`
}

type response struct {
	URL		string		`json:"url"`
	ShortURL	string		`json:"short"`
}

func urlShorten(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON."})
	}

	//Verify the URL

	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL."})
	}

	//Verify Domain Error

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Invalid Domain."})
	}

	var id string

	id = uuid.New().String()[:6]

	// Add verification : Shortne URL is already exists/in-use
	// TBD	
	
	// Add verification : Server Connection
	// TBD
		
	// API Response
	resp := response{
		URL:		body.URL,
		ShortURL:	"",
	}

	resp.ShortURL = os.Getenv("APP_DOMAIN") + "/" + id

	// StatusOK: 200
	return c.Status(fiber.StatusOK).JSON(resp)
}
