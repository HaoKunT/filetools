package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// DeleteOptions is options for delete file or directory
type DeleteOptions struct {
	Path        string
	HardDelete  bool
	ProgressBar bool
}

var deleteOptions *DeleteOptions

// Delete is will delete the file or directory according to DeleteOptions
func Delete(options *DeleteOptions) error {
	deleteOptions = options
	_, err := os.Stat(options.Path)
	if err != nil {
		return err
	}
	// absPath is the Absolute path of the file or directory
	absPath, err := filepath.Abs(options.Path)
	if err != nil {
		return err
	}

	// if soft delete , move the file to the ~/.filetools-trash directory
	if !options.HardDelete {
		homePath, err := Home()
		if err != nil {
			return err
		}
		recyclerBin := filepath.Join(homePath, ".filetools-trash")
		if err := MakeDir(recyclerBin, 0755); err != nil {
			return err
		}
		destinationPathtmp := filepath.Join(recyclerBin, absPath)
		destinationPath := destinationPathtmp
		if runtime.GOOS == "windows" {
			reg := regexp.MustCompile(`\\([A-Z]):\\`)
			destinationPath = reg.ReplaceAllString(destinationPathtmp, `\$1\`)
		}
		destinationPathDir := filepath.Dir(destinationPath)
		if err := MakeDir(destinationPathDir, 0755); err != nil {
			return err
		}
		if err := move(absPath, destinationPath); err != nil {
			return err
		}
		fmt.Printf("Soft delete: move %s -> %s\n", absPath, destinationPath)
	} else {
		if err := os.RemoveAll(absPath); err != nil {
			return err
		}
		fmt.Printf("Hard delete: remove %s\n", absPath)
	}
	return nil
}
