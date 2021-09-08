package logging

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/src/helpers"
)

// logRootPath is the root directory which contains all log files
var logRootPath = fmt.Sprintf("%s/logs/", helpers.GetRootPath())

// Log write in the logRootPath file the given data
func Log(user, passwd string, ip net.Addr) {

	if !helpers.DirExists(logRootPath) {
		err := os.Mkdir(logRootPath, os.ModePerm)
		helpers.CheckErr(err)
	}

	logFilename := time.Now().Format("01-02-2006")

	fd, err := os.OpenFile(fmt.Sprintf("%s/%s.%s", logRootPath, logFilename, "log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	helpers.CheckErr(err)
	defer fd.Close()

	logger := log.New(fd, "", log.Ltime)
	logger.Println(fmt.Sprintf("User: %s, Password: %s, IPAddress: %s", user, passwd, ip))
}
