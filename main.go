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
		{
			Name:     "info",
			Category: "General",
			Usage:    "Show file information",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "file, f",
					Usage:    "the `filename`",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "crypto, c",
					Usage: "use crypto to calculate the md5, sha1 and, sha256, this flag only works when the file is not directory",
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.InfoOptions{
					FileName: c.String("file"),
					Crypto:   c.Bool("crypto"),
				}
				if err := tools.Info(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Info: %s", err), -1)
				}
				return nil
			},
		},
		{
			Name:     "copy",
			Category: "General",
			Usage:    "copy a file or directory",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "src, s",
					Usage:    "source `file`",
					Required: true,
				},
				cli.StringFlag{
					Name:     "dst, d",
					Usage:    "destination `file`",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "recursive, r",
					Usage: "Recursive copy",
				},
				cli.BoolFlag{
					Name:  "createDir, c",
					Usage: "Create dst directory if it doesn't exist",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "Verbose output",
				},
				cli.BoolFlag{
					Name:  "pbar, pb",
					Usage: "Whether show progress bar, if true, it will calculate the size of directory, it can only be used with no verbose flag",
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.CopyOptions{
					Src:         c.String("src"),
					Dst:         c.String("dst"),
					IsDir:       c.Bool("recursive"),
					CreateDir:   c.Bool("createDir"),
					Verbose:     c.Bool("verbose"),
					ProgressBar: c.Bool("pbar"),
				}
				if err := tools.Copy(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Copy: %s", err), -1)
				}
				return nil
			},
		},
		{
			Name:     "rename",
			Category: "General",
			Usage:    "Rename a file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "src, s",
					Usage:    "source `file`",
					Required: true,
				},
				cli.StringFlag{
					Name:     "dst, d",
					Usage:    "destination `file`",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "keep-original, k",
					Usage: "keep original file",
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.RenameOptions{
					Src:          c.String("src"),
					Dst:          c.String("dst"),
					KeepOriginal: c.Bool("keep-original"),
				}
				if err := tools.Rename(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Rename: %s", err), -1)
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
