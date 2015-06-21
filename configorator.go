package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "configorator"
	app.Usage = "make cloud-configs beautiful"
	app.Version = "0.0.1"
	app.Author = "Lee Chang <l33tchang@gmail.com>"
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run web",
			Subcommands: []cli.Command{
				{
					Name:   "web",
					Action: runWeb,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "template",
							Value: "",
							Usage: "the directory where the template files are located",
						},
						cli.StringFlag{
							Name:  "port",
							Value: "8000",
							Usage: "the port that the server will listen on",
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func runWeb(ctx *cli.Context) {
	//fmt.Println("This is the template: ", ctx.String("template"))
	webListener(ctx.String("template"), ctx.String("port"))
}
