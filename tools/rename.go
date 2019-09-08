package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

type RenameOptions struct {
	Src          string
	Dst          string
	KeepOriginal bool
}

var renameOptions *RenameOptions

func Rename(options *RenameOptions) error {
	renameOptions = options
	info, err := os.Stat(options.Src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("Not support directory")
	}
	if options.KeepOriginal {
		copyOptions = &CopyOptions{
			Src: options.Src,
			Dst: options.Dst,
		}
		// like copy
		srcAbs, err := filepath.Abs(options.Src)
		if err != nil {
			return err
		}
		dstAbs, err := filepath.Abs(options.Dst)
		if err != nil {
			return err
		}
		if srcAbs == dstAbs {
			return fmt.Errorf("Source and destination are the same file")
		}
		n, err := copyFile(srcAbs, dstAbs)
		if err != nil {
			return err
		}
		fmt.Printf("Rename File %s to %s, Size: %s (%d bytes)\n", options.Src, options.Dst, transSize(n), n)
		return nil
	}
	err = os.Rename(options.Src, options.Dst)
	if err != nil {
		return err
	}
	return nil
}
