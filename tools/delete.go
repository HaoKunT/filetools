package tools

import (
	"fmt"
	"os"
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
	srcInfo, err := NewExtInfo(deleteOptions.Path)
	if err != nil {
		return err
	}

	// if soft delete , move the file to the ~/.filetools-trash directory
	if !options.HardDelete {
		dstpath, err := remove(srcInfo)
		if err != nil {
			return err
		}
		fmt.Printf("Soft delete: move %s -> %s (%s)\n", srcInfo.AbsPath, dstpath, transSize(srcInfo.RealSize))
	} else {
		if err := os.RemoveAll(srcInfo.AbsPath); err != nil {
			return err
		}
		fmt.Printf("Hard delete: remove %s\n", srcInfo.AbsPath)
	}
	return nil
}
