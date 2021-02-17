package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"

	gettext "github.com/chai2010/gettext-go"
	"github.com/haokunt/filetools/tools"
	"github.com/urfave/cli"

	_ "embed"
)

var version = "MISSING build version [git hash]"

//go:embed local.zip
var localZipBytes []byte

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	gettext.BindLocale(gettext.New("filetools", "local.zip", localZipBytes))
	gettext.SetDomain("filetools")

	app := cli.NewApp()
	app.Name = "filetools"
	app.Usage = gettext.Gettext("some tools about file")
	app.Version = fmt.Sprintf("Git:[%s] (%s)", strings.ToUpper(version), runtime.Version())
	app.Commands = []cli.Command{
		{
			Name:     "compare",
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("compare two files or directory"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "src, s",
					Usage: gettext.Gettext("source `file`"),
					Value: "./",
				},
				cli.StringFlag{
					Name:  "dst, d",
					Usage: gettext.Gettext("destination `file`"),
					Value: "./",
				},
				cli.BoolFlag{
					Name:  "content, c",
					Usage: gettext.Gettext("whether compare content or not"),
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: gettext.Gettext("print verbose output"),
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
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("Show file information"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "file, f",
					Usage:    gettext.Gettext("the `filename`"),
					Required: true,
				},
				cli.BoolFlag{
					Name:  "crypto, c",
					Usage: gettext.Gettext("use crypto to calculate the md5, sha1 and, sha256, this flag only works when the file is not directory"),
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
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("copy a file or directory"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "src, s",
					Usage:    gettext.Gettext("source `file`"),
					Required: true,
				},
				cli.StringFlag{
					Name:     "dst, d",
					Usage:    gettext.Gettext("destination `file`"),
					Required: true,
				},
				cli.BoolFlag{
					Name:  "recursive, r",
					Usage: gettext.Gettext("Recursive copy"),
				},
				cli.BoolFlag{
					Name:  "createDir, c",
					Usage: gettext.Gettext("Create dst directory if it doesn't exist"),
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: gettext.Gettext("print verbose output"),
				},
				cli.BoolFlag{
					Name:  "pbar, pb",
					Usage: gettext.Gettext("Whether show progress bar, if true, it will calculate the size of directory, it can only be used with no verbose flag"),
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
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("Rename a file"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "src, s",
					Usage:    gettext.Gettext("source `file`"),
					Required: true,
				},
				cli.StringFlag{
					Name:     "dst, d",
					Usage:    gettext.Gettext("destination `file`"),
					Required: true,
				},
				cli.BoolFlag{
					Name:  "keep-original, k",
					Usage: gettext.Gettext("keep original file"),
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
		{
			Name:     "list",
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("list all files in a directory"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "directory, d",
					Usage:    gettext.Gettext("the directory `path`"),
					Required: true,
				},
				cli.StringSliceFlag{
					Name: "sortby, s",
					Usage: gettext.Gettext("sort the files by some fields, now supported: name, size, " +
						"you can use it like this, --sortby name:desc --sortby size:asc, " +
						"desc means descending, asc means ascending, desc and asc can be write in short name 'a' and 'd'." +
						"the sort field appears first will be the main field, and so on. " +
						"if sort, the command will calculate first, so the performance will be slower."),
				},
				cli.IntFlag{
					Name:  "limit, l",
					Usage: gettext.Gettext("limit the `number` of files to show"),
					Value: math.MaxInt32,
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.ListOptions{
					Path:   c.String("directory"),
					SortBy: c.StringSlice("sortby"),
					Limit:  c.Int("limit"),
				}
				if err := tools.List(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error List: %s", err), -1)
				}
				return nil
			},
		},
		{
			Name:     "delete",
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("delete a file or directory, it will move file to the ~/.filetools-trash default"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "file, f",
					Usage:    gettext.Gettext("the `file` or directory"),
					Required: true,
				},
				cli.BoolFlag{
					Name:  "hard, hd",
					Usage: gettext.Gettext("hard delete, it will delete the directory directly"),
				},
			},
			Action: func(c *cli.Context) error {
				Options := tools.DeleteOptions{
					Path:       c.String("file"),
					HardDelete: c.Bool("hard"),
				}
				if err := tools.Delete(&Options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Delete: %s", err), -1)
				}
				return nil
			},
		},
		{
			Name:     "move",
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("move a file or directory, it works like `mv`"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "source, s",
					Usage:    gettext.Gettext("source `file`"),
					Required: true,
				},
				cli.StringFlag{
					Name:     "destination, d",
					Usage:    gettext.Gettext("destination `file`"),
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				options := tools.MoveOptions{
					Src: c.String("source"),
					Dst: c.String("destination"),
				}
				if err := tools.Move(&options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Move: %s", err), -1)
				}
				return nil
			},
		},
		{
			Name:     "server",
			Category: gettext.Gettext("General"),
			Usage:    gettext.Gettext("Using a Simple HTTP Server to download or upload files"),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "directory, d",
					Usage: gettext.Gettext("the directory to be served"),
					Value: pwd,
				},
				cli.StringFlag{
					Name:  "host",
					Usage: gettext.Gettext("the host you want to use"),
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: gettext.Gettext("port"),
					Value: 9000,
				},
				cli.StringFlag{
					Name:  "user, u",
					Usage: gettext.Gettext("username of the server"),
				},
				cli.StringFlag{
					Name:  "password, pw",
					Usage: gettext.Gettext("password of the server, if you not specify a password and username, The server will not use basic authentication"),
				},
				cli.StringFlag{
					Name:  "upload-dir, o",
					Usage: gettext.Gettext("upload directory, default is the same as root directory"),
				},
			},
			Action: func(c *cli.Context) error {
				var uploadDir = c.String("upload-dir")
				if uploadDir == "" {
					uploadDir = c.String("directory")
				}
				options := tools.ServerOptions{
					Host:      c.String("host"),
					Port:      c.Int("port"),
					RootPath:  c.String("directory"),
					UploadDir: uploadDir,
					User:      c.String("user"),
					Password:  c.String("password"),
				}
				if err := tools.Server(&options); err != nil {
					return cli.NewExitError(fmt.Errorf("Error Server: %s", err), -1)
				}
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
