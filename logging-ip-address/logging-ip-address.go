package loggingipaddress

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
)

const logRootPath = "./logs/"
const logIPPath = "./logs/ip-address/"

// LogIPAddr write in the logIPPath file the given ip
func LogIPAddr(ip net.Addr) {

	if !helpers.DirExists(logRootPath) {
		err := os.Mkdir(logRootPath, os.ModePerm)
		helpers.CheckErr(err)
	}

	if !helpers.DirExists(logIPPath) {
		err := os.Mkdir(logIPPath, os.ModePerm)
		helpers.CheckErr(err)
	}

	logFilename := time.Now().Format("01-02-2006")

	fd, err := os.OpenFile(fmt.Sprintf("%s/%s.%s", logIPPath, logFilename, "log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	helpers.CheckErr(err)
	defer fd.Close()

	logger := log.New(fd, "", log.Ltime)
	logger.Println(ip)
}
