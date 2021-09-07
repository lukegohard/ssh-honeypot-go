package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"log/syslog"
	"os"
	"time"

	"github.com/gliderlabs/ssh"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/fakeshell"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
	loggingipaddress "github.com/Ex0dIa-dev/ssh-honeypot-go/logging-ip-address"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/notifier"
	hostkey "github.com/Ex0dIa-dev/ssh-honeypot-go/private-host-key"
)

func init() {
	flag.StringVar(&port, "p", "2222", "enter the port for the honeypot server")
	flag.StringVar(&hostKeyFile, "k", "", "enter the filepath of hostkey file")

	flag.BoolVar(&notifyServiceActivated, "n", false, "activate notifier service")
	flag.BoolVar(&logIPAddressActivated, "li", false, "activate ip address logging")
	flag.BoolVar(&logAllAttempts, "la", false, "logging all attempts, failed too")

}

var port, hostKeyFile string
var notifyServiceActivated, logIPAddressActivated, logAllAttempts bool
var attempts = 0

var config Config

// Config contains the json "struct" of config.json file
type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	// IdleTimeoutSeconds int    `json:"idletimeoutseconds"`
}

func main() {

	flag.Parse()

	// setting Log Output to ==> 1) Os.Stdout 2) Syslog
	logwriter, err := syslog.New(syslog.LOG_INFO, os.Args[0])
	helpers.CheckErr(err)
	log.SetOutput(io.MultiWriter(logwriter, os.Stdout))

	// reading config file
	configBytes, err := ioutil.ReadFile("./config.json")
	helpers.CheckErr(err)
	err = json.Unmarshal(configBytes, &config)
	helpers.CheckErr(err)

	s := &ssh.Server{
		Addr:            fmt.Sprintf("0.0.0.0:%s", port),
		Handler:         sessionHandler,
		PasswordHandler: authHandler,
		IdleTimeout:     45 * time.Second,
	}

	// if hostKeyFile is empty,the key will be auto-generated
	// else key will be read from file
	if hostKeyFile != "" {
		key, err := hostkey.ReadHostKeyFile(hostKeyFile)
		helpers.CheckErr(err)
		s.AddHostKey(key)
	}

	log.Printf("[+]Starting Honeypot Server on Address: %v\n", s.Addr)
	if hostKeyFile == "" {
		log.Print("[+]Honeypot HostKey Mode: auto-generated")
	} else {
		log.Printf("[+]Honeypot HostKey Mode: user-input-file")
	}
	log.Printf("[+]Notifier Service Activated: %v", notifyServiceActivated)
	log.Printf("[+]Logging IP Address: %v", logIPAddressActivated)
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
	body := fmt.Sprintf("User: %s,Password: %s, Address: %s, Status: ", ctx.User(), passwd, ctx.RemoteAddr())

	if ctx.User() != config.User || passwd != config.Password {

		if logAllAttempts {
			log.Println(fmt.Sprintf("[%d]%s%s", attempts, body, "failed"))
		}

		if notifyServiceActivated {
			notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), body)
		}

		if logIPAddressActivated {
			loggingipaddress.LogIPAddr(ctx.RemoteAddr())
		}

		return false
	}

	log.Println(fmt.Sprintf("[%d]%s%s", attempts, body, "connected"))

	if notifyServiceActivated {
		notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), body)
	}

	if logIPAddressActivated {
		loggingipaddress.LogIPAddr(ctx.RemoteAddr())
	}

	return true
}
