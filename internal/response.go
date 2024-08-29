package internal

import "github.com/gofiber/fiber/v2"

type TheResponse struct {
	StatusCode  int         `json:"code"`
	StatusError bool        `json:"error"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func ReturnTheResponse(c *fiber.Ctx, se bool, sc int, m string, dt interface{}) error {
	tr := TheResponse{sc, se, m, dt}

	return c.Status(sc).JSON(tr)
}
