package fakeshell

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers"
	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

const cmdsFilePath = "./fakeshell/cmds.txt"

// FakeShell create a fake shell to waste attacker's time
// Read command, and "execute" them
func FakeShell(s ssh.Session) {

	bytes, err := ioutil.ReadFile(cmdsFilePath)
	helpers.CheckErr(err)
	commandsList := strings.Split(string(bytes), "\n")

	term := term.NewTerminal(s, fmt.Sprintf(
		"%s%s%s@%s%s%s>$%s ",
		colors.Yellow,
		s.User(),
		colors.Green,
		colors.Blue,
		s.LocalAddr(),
		colors.Green,
		colors.Reset,
	))

	for {

		ln, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		if ln == "exit" {
			break
		}

		commandAndArgs := strings.Split(ln, " ")
		command := commandAndArgs[0]
		unknown := true
		for _, c := range commandsList {
			if c == command {
				unknown = false
				break
			}
		}
		if unknown {
			writers.ColorWrite(term, fmt.Sprintf("unknown command: %s", command), colors.Red)
			writers.PrintEnd(term, 1)
		}
	}

	_, err = term.Write([]byte(colors.Reset))
	if err != nil {
		if err == io.EOF {
			s.Close()
			return
		} else {
			panic(err)
		}
	}
	s.Close()
}
