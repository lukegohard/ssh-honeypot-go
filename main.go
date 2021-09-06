package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"

	"github.com/gliderlabs/ssh"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
	loggingipaddress "github.com/Ex0dIa-dev/ssh-honeypot-go/logging-ip-address"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/notifier"
	hostkey "github.com/Ex0dIa-dev/ssh-honeypot-go/private-host-key"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
)

func init() {
	flag.StringVar(&port, "p", "2222", "enter the port for the honeypot server")
	flag.StringVar(&hostKeyFile, "k", "", "enter the filepath of hostkey file")
	flag.BoolVar(&notifyService, "n", false, "activate notifier service")
	flag.BoolVar(&logIpAddress, "li", false, "activate ip address logging")
}

var port, hostKeyFile string
var notifyService, logIpAddress bool
var attempts = 0

func main() {

	flag.Parse()

	// setting Log Output to ==> 1) Os.Stdout 2) Syslog
	logwriter, err := syslog.New(syslog.LOG_INFO, os.Args[0])
	helpers.CheckErr(err)
	log.SetOutput(io.MultiWriter(logwriter, os.Stdout))

	s := &ssh.Server{
		Addr:            fmt.Sprintf("0.0.0.0:%s", port),
		Handler:         sessionHandler,
		PasswordHandler: authHandler,
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
	log.Printf("[+]Notifier Service Activated: %v", notifyService)
	log.Printf("[+]Logging IP Address: %v", logIpAddress)
	log.Fatal(s.ListenAndServe())

}

//sessionHandler is called after authentication
func sessionHandler(s ssh.Session) {
	writers.ColorWrite(s, writers.Welcome, colors.Green)
	writers.PrintEnd(s, 1)
}

//authHandler collects authentication info(username,password,ip) and logs them
func authHandler(ctx ssh.Context, passwd string) bool {
	attempts++
	body := fmt.Sprintf("User: %s,Password: %s, Address: %s", ctx.User(), passwd, ctx.RemoteAddr())
	log.Println(fmt.Sprintf("[%d]%s", attempts, body))

	if notifyService {
		notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), body)
	}

	if logIpAddress {
		loggingipaddress.LogIPAddr(ctx.RemoteAddr())
	}

	return true
}
