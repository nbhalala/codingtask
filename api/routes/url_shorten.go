package routes

import(
	"os"
	"strconv"
	"github.com/nbhalala/codingtask/database"
	"github.com/nbhalala/codingtask/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type request struct {
	URL			string		`json:"url"`
	CustomShort	string		`json:"short"`
}

type response struct {
	URL			string		`json:"url"`
	CustomShort	string		`json:"short"`
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

	if body.CustomShort == ""{
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.createClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != ""{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error":"URL custom short is already in use."})
	}

	err = r.set(database.Ctx, id, body.URL).Err()
	if err !+ nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Unable to connect to the Server."})
	}

	resp := response{
		URL:			body.URL,
		CustomShort:	""
	}

	resp.CustomShort = os.Getenv("APP_DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
