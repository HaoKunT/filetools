package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/haokunt/filetools/tools"
	"github.com/urfave/cli"
)

var version = "MISSING build version [git hash]"

func main() {
	app := cli.NewApp()
	app.Name = "filetools"
	app.Usage = "some tools about file"
	app.Version = fmt.Sprintf("Git:[%s] (%s)", strings.ToUpper(version), runtime.Version())
	app.Commands = []cli.Command{
		{
			Name:     "compare",
			Category: "General",
			Usage:    "compare two files or directory",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "src, s",
					Usage: "source `file`",
					Value: "./",
				},
				cli.StringFlag{
					Name:  "dst, d",
					Usage: "destination `file`",
					Value: "./",
				},
				cli.BoolFlag{
					Name:  "content, c",
					Usage: "whether compare content or not",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "print verbose output",
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.CompareOptions{
					Src:     c.String("src"),
					Dst:     c.String("dst"),
					Content: c.Bool("content"),
					Verbose: c.Bool("verbose"),
				}
				if err := tools.Compare(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Compare: %s", err), -1)
				}
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
