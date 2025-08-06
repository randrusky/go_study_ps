package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router)  {
	h := &HomeHandler{
		router: router,
	}

	api := router.Group("/api")
	api.Get("/home", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusOK, "Welcome to the Home Page")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("An error occurred")
}