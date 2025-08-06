package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

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
	log.Info().
		Bool("isAdmin", true).
		Str("email", "andrur@ukr.net").
		Int("id", 10).
		Msg("Info")
	return c.SendString("An error occurred")
}