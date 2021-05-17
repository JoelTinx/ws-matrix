package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetWelcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{ "message": "Hello!, I'm Joel.\nMay be you want to go to /rotate" })
}