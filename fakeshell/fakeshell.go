package fakeshell

import (
	"fmt"
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
	commands_list := strings.Split(string(bytes), "\n")

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
		helpers.CheckErr(err)
		if ln == "exit" {
			break
		}

		command_and_args := strings.Split(ln, " ")
		command := command_and_args[0]
		unknown := true
		for _, c := range commands_list {
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
	helpers.CheckErr(err)
	s.Close()
}
