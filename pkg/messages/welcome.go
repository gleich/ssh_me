package messages

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/ssh_me/pkg/commands"
	"github.com/gliderlabs/ssh"
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Output a welcome message to the user
func OutputWelcome(s ssh.Session) {
	fmt.Fprintln(s, color.Color(233, 39, 65).Sprint(`
┌┬┐┌─┐┌┬┐┌┬┐  ┌─┐┬  ┌─┐┬┌─┐┬ ┬
│││├─┤ │  │   │ ┬│  ├┤ ││  ├─┤
┴ ┴┴ ┴ ┴  ┴   └─┘┴─┘└─┘┴└─┘┴ ┴
`))

	runes := []rune("👋 Hello, welcome to my ssh server!\nThis server acts like a terminal but with the following commands:\n" + commands.HelpMessage)
	for _, char := range runes {
		fmt.Fprint(s, string(char))
		time.Sleep(time.Millisecond * 20)
	}
}