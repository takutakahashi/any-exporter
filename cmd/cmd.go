package main

import (
	"log"
	"os"

	"github.com/takutakahashi/any-exporter/pkg/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "any-exporter"
	app.Usage = "prometheus exporter for anything"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config.yaml filepath",
		},
	}
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	configPath := c.String("config")
	if configPath == "" {
		cli.ShowAppHelp(c)
		return nil
	}
	s, err := server.New(configPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	go s.StartWorker()
	s.StartServer()
	return nil
}
