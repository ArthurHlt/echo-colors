package main

import (
	"github.com/mitchellh/colorstring"
	"os"
	"strings"
	"github.com/codegangsta/cli"
)

var notrailing bool

func main() {
	app := cli.NewApp()
	app.Name = "echoc"
	app.Usage = "Echo in colors some text easily"
	app.Version = "1.0.0"
	app.UsageText = "echoc \"[red] this is red [yellow]and yellow [_red_]with red background [bold] and bold [reset] and no colors\""
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "n",
			Destination: &notrailing,
			Usage: "Optional. Do not print the trailing newline character",
		},

	}
	app.Action = colorIt
	app.Run(os.Args)

}
func colorIt(c *cli.Context) error {
	text := strings.Join(c.Args(), " ")
	if notrailing {
		colorstring.Print(text)
	} else {
		colorstring.Println(text)
	}
	return nil
}