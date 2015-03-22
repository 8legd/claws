// Command Line Amazon Web Services
package main

import (
	"github.com/8legd/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "claws"
	app.Usage = "Command Line Amazon Web Services"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "Access Key ID",
			EnvVar: "CLAWS_KEY",
		},
		cli.StringFlag{
			Name:   "secret, s",
			Usage:  "Secret Access Key",
			EnvVar: "CLAWS_SECRET",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "s3",
			Usage: "s3 commands",
			Subcommands: []cli.Command{
				{
					Name:  "put",
					Usage: "put local-path bucket@region[:remote-path] \nUpload local-path and store it at the specified remote location",
					Action: func(c *cli.Context) {
						println("TODO: implement s3 put", c.Args().First())
					},
				},
				{
					Name:  "get",
					Usage: "get bucket@region:remote-path [local-path] \nRetrieve the remote-path and store it on the local machine",
					Action: func(c *cli.Context) {
						println("TODO: implement s3 get", c.Args().First())
					},
				},
			},
		},
	}

	app.EnableBashCompletion = true
	app.Run(os.Args)
}
