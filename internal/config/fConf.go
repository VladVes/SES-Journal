package config

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

var FiberConfig = fiber.Config{
	Immutable:      true,
	ReadBufferSize: 1024 * 10,
	ReadTimeout:    3 * time.Second,
	WriteTimeout:   3 * time.Second,
}
