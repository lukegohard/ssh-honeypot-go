package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/src/fakeshell"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/src/helpers"
	logging "github.com/Ex0dIa-dev/ssh-honeypot-go/src/logging"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/src/notifier"
	privatehostkey "github.com/Ex0dIa-dev/ssh-honeypot-go/src/private-host-key"
)

func init() {
	flag.StringVar(&port, "port", "2222", "enter the port for the honeypot server")

	flag.BoolVar(&notifyServiceActivated, "notify", false, "activate notifier service")
	flag.BoolVar(&logActivated, "log", false, "activate ip address logging")
	flag.BoolVar(&logAllAttempts, "log-all", false, "logging all attempts, failed too")

}

var port string
var notifyServiceActivated, logActivated, logAllAttempts bool
var attempts = 0

var config helpers.Config

func main() {

	flag.Parse()

	config = helpers.ParseConfigFile()

	s := &ssh.Server{
		Addr:            fmt.Sprintf("0.0.0.0:%s", port),
		Handler:         sessionHandler,
		PasswordHandler: authHandler,
		IdleTimeout:     45 * time.Second,
	}

	var keyFilePath string

	// if this env exists, honeypot is running on docker, and path is "/app/config/hostkey_rsa"
	if os.Getenv("HONEYPOT_HOSTKEYFILE") == "" {
		keyFilePath = fmt.Sprintf("%s/config/hostkey_rsa", helpers.GetRootPath())
	} else {
		keyFilePath = os.Getenv("HONEYPOT_HOSTKEYFILE")
	}

	keyFileBool := helpers.FileExists(keyFilePath)

	// keyFileBool is true(file exists), the key will be read from file
	// else will be auto-generated
	if keyFileBool {
		key, err := privatehostkey.ReadHostKeyFile(keyFilePath)
		helpers.CheckErr(err)
		s.AddHostKey(key)
	}

	log.Printf("[+]Starting Honeypot Server on Address: %v\n", s.Addr)
	if keyFileBool {
		log.Printf("[+]Honeypot HostKey Mode: user-input-file")
	} else {
		log.Print("[+]Honeypot HostKey Mode: auto-generated")
	}
	log.Printf("[+]Notifier Service Activated: %v", notifyServiceActivated)
	log.Printf("[+]Logging IP Address: %v", logActivated)
	log.Printf("[+]Logging All Attempts: %v", logAllAttempts)
	fmt.Println()
	log.Fatal(s.ListenAndServe())

}

// sessionHandler is called after authentication
func sessionHandler(s ssh.Session) {
	fakeshell.FakeShell(s)
}

// authHandler collects authentication info(username,password,ip) and logs them
func authHandler(ctx ssh.Context, passwd string) bool {
	attempts++

	var clientIP string
	if strings.ContainsRune(ctx.RemoteAddr().String(), ':') {
		clientIP, _, _ = net.SplitHostPort(ctx.RemoteAddr().String())
	} else {
		clientIP = ctx.RemoteAddr().String()
	}

	body := fmt.Sprintf("User: %s,Password: %s, Address: %s, Status: ", ctx.User(), passwd, clientIP)

	if ctx.User() != config.Auth.User || passwd != config.Auth.Password {

		status := "failed"

		if logAllAttempts {
			log.Println(fmt.Sprintf("[%d]%s%s", attempts, body, status))
		}

		if notifyServiceActivated {
			notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), fmt.Sprintf("body%s", status))
		}

		if logActivated {
			logging.Log(ctx.User(), passwd, clientIP, status)
		}

		return false
	}

	status := "connected"

	log.Println(fmt.Sprintf("[%d]%s%s", attempts, body, status))

	if notifyServiceActivated {
		notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), fmt.Sprintf("body%s", status))
	}

	if logActivated {
		logging.Log(ctx.User(), passwd, clientIP, status)
	}

	return true
}
