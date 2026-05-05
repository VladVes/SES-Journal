package app

import (

	// "time"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"

	conf "github.com/VladVes/SES-Journal/internal/config"
	. "github.com/VladVes/SES-Journal/internal/logger"
)

func Run() {

	Log.WithFields(logrus.Fields{
		"app":     "SES-Journal",
		"version": "0.0.1",
		"author":  "VladVes",
		"port":    "8080",
		"host":    "localhost",
	}).Info("App is running")

	app := fiber.New(conf.FiberConfig)
	appConf := conf.GetAppConfig()

	// -------------------------------------------------
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	// -------------------------------------------------

	lErr := app.Listen(":" + appConf.Port)
	if lErr != nil {
		Log.WithError(lErr).Fatal("app is not running")
	}

}
