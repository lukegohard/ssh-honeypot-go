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
	"github.com/Ex0dIa-dev/ssh-honeypot-go/notifier"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
)

func init() {
	flag.StringVar(&port, "p", "2222", "enter the port for the honeypot server")
	flag.BoolVar(&notifyService, "n", false, "activate notifier service")
}

var port string
var notifyService bool
var attempts = 0

func main() {

	flag.Parse()

	logwriter, err := syslog.New(syslog.LOG_INFO, os.Args[0])
	helpers.CheckErr(err)

	log.SetOutput(io.MultiWriter(logwriter, os.Stdout))

	s := &ssh.Server{
		Addr:            fmt.Sprintf("0.0.0.0:%s", port),
		Handler:         sessionHandler,
		PasswordHandler: authHandler,
	}

	log.Printf("[+]Starting Honeypot Server on Address: %v\n", s.Addr)
	log.Printf("[+]Notifier Service Activated: %v", notifyService)
	log.Fatal(s.ListenAndServe())
}

func sessionHandler(s ssh.Session) {
	writers.ColorWrite(s, writers.Welcome, colors.Green)
	writers.PrintEnd(s, 1)
}

func authHandler(ctx ssh.Context, passwd string) bool {
	attempts++
	body := fmt.Sprintf("User: %s,Password: %s, Address: %s", ctx.User(), passwd, ctx.RemoteAddr())
	log.Println(fmt.Sprintf("[%d]%s", attempts, body))

	if notifyService {
		notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), body)
	}

	return true
}
