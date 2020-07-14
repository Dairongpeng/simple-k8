package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"simple-k8/base"
	"simple-k8/log"
)

func main() {
	fmt.Println("Simple-K8 " + VERSION)
	base.ConfigureProductVersion(VERSION)

	app := cli.NewApp()
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config path",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug info",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		log.SetDebug(ctx.Bool("debug"))
		if err := ParseConfig(ctx.String("config")); err != nil {
			return err
		}

		return base.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exit with failure: %v\n", err)
	}
}
