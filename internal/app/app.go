package app

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	conf "github.com/VladVes/SES-Journal/internal/config"
	dataRepo "github.com/VladVes/SES-Journal/internal/repository"

	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/router"
)

func Run() {
	appConf := conf.GetAppConfig()

	logger.Init(appConf.LogLevel)

	logger.Log.WithFields(logrus.Fields{
		"app":     "SES-Journal",
		"version": "0.0.1",
		"author":  "VladVes",
		"port":    appConf.Port,
	}).Info("App is starting")

	rt := router.New()
	handler := rt.Register()

	server := &http.Server{
		Addr:         ":" + appConf.Port,
		Handler:      handler,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	_, err := dataRepo.DBConnet(appConf.Dsn, appConf.Env)

	if err != nil {
		logger.Log.WithError(err).Fatal("DB connection failed")
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Log.WithError(err).Fatal("App failed to start")
	}

}
