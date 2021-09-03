package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
	"github.com/gliderlabs/ssh"
)

func init() {
	flag.StringVar(&port, "p", "2222", "enter the port for the honeypot server")

}

var port string
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
	log.Fatal(s.ListenAndServe())
}

func sessionHandler(s ssh.Session) {
	writers.ColorWrite(s, writers.Welcome, colors.Green)
	writers.PrintEnd(s, 1)
}

func authHandler(ctx ssh.Context, passwd string) bool {
	attempts++
	log.Printf("[%d]User: %s,Password: %s, Address: %s", attempts, ctx.User(), passwd, ctx.RemoteAddr())
	return true

}
