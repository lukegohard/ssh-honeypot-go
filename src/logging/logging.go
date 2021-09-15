package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/src/helpers"
	"github.com/sirupsen/logrus"
)

// Log write in the logRootPath file the given data
func Log(user, passwd, ip, status string) {

	var logRootPath string

	// if this env exists, honeypot is running on docker, and path is "/app/logs/"
	if os.Getenv("HONEYPOT_LOGSPATH") == "" {
		logRootPath = fmt.Sprintf("%s/logs/", helpers.GetRootPath())
	} else {
		logRootPath = os.Getenv("HONEYPOT_LOGSPATH")
	}

	if logRootPath == fmt.Sprintf("%s/logs/", helpers.GetRootPath()) && !helpers.DirExists(logRootPath) {
		err := os.Mkdir(logRootPath, os.ModePerm)
		helpers.CheckErr(err)
	}

	logFilename := time.Now().Format("01-02-2006")

	fd, err := os.OpenFile(fmt.Sprintf("%s%s.%s", logRootPath, logFilename, "log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	helpers.CheckErr(err)
	defer fd.Close()

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "15:04:05",
	})
	logrus.SetOutput(fd)
	logrus.WithFields(logrus.Fields{
		"ipaddr":   ip,
		"username": user,
		"password": passwd,
		"status":   status,
	}).Info("attempt")
}
