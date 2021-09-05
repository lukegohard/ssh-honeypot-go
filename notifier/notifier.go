package notifier

import (
	"os/exec"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
)

//send a notification using "notify-send" tool
func SendNotify(appName, title, body string) {
	args := []string{}
	args = append(args, "-a", appName)
	args = append(args, title)
	args = append(args, body)

	cmd := exec.Command("notify-send", args...)
	helpers.CheckErr(cmd.Run())
}
