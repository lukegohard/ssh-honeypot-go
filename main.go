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
	hostkey "github.com/Ex0dIa-dev/ssh-honeypot-go/private-host-key"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
)

func init() {
	flag.StringVar(&port, "p", "2222", "enter the port for the honeypot server")
	flag.StringVar(&hostKeyFile, "k", "", "enter the filepath of hostkey file")
	flag.BoolVar(&notifyService, "n", false, "activate notifier service")
}

var port, hostKeyFile string
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

	if hostKeyFile != "" {
		key, err := hostkey.ReadHostKeyFile(hostKeyFile)
		helpers.CheckErr(err)
		s.AddHostKey(key)
	}

	//logging some infos
	log.Printf("[+]Starting Honeypot Server on Address: %v\n", s.Addr)
	if hostKeyFile == "" {
		log.Print("[+]Honeypot HostKey Mode: auto-generated")
	} else {
		log.Printf("[+]Honeypot HostKey Mode: user-input-file")
	}
	log.Printf("[+]Notifier Service Activated: %v", notifyService)
	log.Fatal(s.ListenAndServe())

}

//function called after authentication
func sessionHandler(s ssh.Session) {
	writers.ColorWrite(s, writers.Welcome, colors.Green)
	writers.PrintEnd(s, 1)
}

//function where we collect authentication info(username,password,ip) and log them
func authHandler(ctx ssh.Context, passwd string) bool {
	attempts++
	body := fmt.Sprintf("User: %s,Password: %s, Address: %s", ctx.User(), passwd, ctx.RemoteAddr())
	log.Println(fmt.Sprintf("[%d]%s", attempts, body))

	if notifyService {
		notifier.SendNotify("ssh-honeypot-go", fmt.Sprintf("Connection Attempt: %d", attempts), body)
	}

	return true
}
