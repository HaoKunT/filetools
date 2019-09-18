// +build !windows

package tools

import "os"

func move(oldpath, newpath string) error {
	if err := os.Rename(oldpath, newpath); err != nil {
		return err
	}
	return nil
}
