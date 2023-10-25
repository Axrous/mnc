package helper

import (
	"os"

	"github.com/Axrous/mnc/model/web"
	"github.com/sirupsen/logrus"
)

func GenerateLogMiddleware(logRequest web.LogRequest) {

	var log *logrus.Logger
	log = logrus.New()

	logFile := "logger.log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	PanicIfError(err)
	log.SetOutput(file)

	switch {
	case logRequest.StatusCode >= 500:
		log.Error(logRequest)
	case logRequest.StatusCode >= 400:
		log.Warning(logRequest)
	default:
		log.Info(logRequest)
	}
}