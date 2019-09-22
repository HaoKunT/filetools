// +build !windows

package tools

import (
	"os"
	"path/filepath"
)

func move(oldpathInfo *ExtFileInfo, newpath string) error {
	if err := os.Rename(oldpathInfo.AbsPath, newpath); err != nil {
		return err
	}
	return nil
}

// remove the file or directory, soft remove, it will move the file to trash folder
func remove(srcInfo *ExtFileInfo) (newpath string, err error) {
	homePath, err := Home()
	if err != nil {
		return "", err
	}
	recyclerBin := filepath.Join(homePath, ".ft-trash")
	if err := MakeDir(recyclerBin, 0755); err != nil {
		return "", err
	}

	destinationPath := filepath.Join(recyclerBin, srcInfo.AbsPath)
	destinationPathDir := filepath.Dir(destinationPath)
	if err := MakeDir(destinationPathDir, 0755); err != nil {
		return "", err
	}
	if err := move(srcInfo, destinationPath); err != nil {
		return "", err
	}
	return destinationPath, nil
}
