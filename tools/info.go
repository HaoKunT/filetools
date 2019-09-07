package tools

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type InfoOptions struct {
	FileName string
}

var infoOptions *InfoOptions

func Info(options *InfoOptions) error {
	infoOptions = options

	info, err := os.Stat(options.FileName)
	if err != nil {
		return err
	}

	absName, err := filepath.Abs(info.Name())
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Absolute path: %s\n", absName)

	if info.IsDir() {
		fmt.Printf("Type: directory\n")
		if err := calculateDir(infoOptions.FileName); err != nil {
			return err
		}
	} else {
		kind, err := filetype.MatchFile(info.Name())
		if err != nil {
			return err
		}
		fmt.Printf("Type: %s\n", kind.Extension)
		if kind != filetype.Unknown {
			fmt.Printf("MIME-type: %s\n", kind.MIME.Value)
		}
		fmt.Printf("Size: %d\n", info.Size())
	}
	fmt.Printf("Mode: %s\n", info.Mode())
	fmt.Printf("ModTime: %s\n", info.ModTime())
	fmt.Printf("Is Regular: %v\n", info.Mode().IsRegular())
	return nil
}

func calculateDir(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			// fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			// return filepath.SkipDir
			srcDirNum++
			return nil
		}
		srcFileNum++
		srcSize += info.Size()

		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("Include: %d files, %d directories\n", srcFileNum, srcDirNum)
	fmt.Printf("All Size: %s\n", transSize(srcSize))
	return nil
}
