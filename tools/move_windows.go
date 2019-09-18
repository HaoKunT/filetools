package tools

import (
	"errors"
	"os"
	"strings"
	"syscall"
)

// moveFile funtion is working with windows to move file across volumn
func moveFile(oldpath, newpath string) error { //跨卷移动
	from, err := syscall.UTF16PtrFromString(oldpath)
	if err != nil {
		return err
	}
	to, err := syscall.UTF16PtrFromString(newpath)
	if err != nil {
		return err
	}
	return syscall.MoveFile(from, to) //windows API
}

// move function is working with windows to move a directory across volumn
func move(oldpath, newpath string) error {
	oldpathInfo, err := NewExtInfo(oldpath)
	if err != nil {
		return err
	}
	// if file larger than 4g, it should be hard deleted
	if oldpathInfo.RealSize > 4294967296 {
		return errors.New("the file is too large, please remove it directly, use -hd")
	}
	if !oldpathInfo.IsDir() {
		return moveFile(oldpath, newpath)
	}
	if err := MakeDir(newpath, oldpathInfo.Mode()); err != nil {
		return err
	}
	allFiles := oldpathInfo.Traversal()
	for _, f := range allFiles {
		newfilepath := strings.Replace(f.AbsPath, oldpath, newpath, -1)
		if f.IsDir() {
			if err := MakeDir(newfilepath, f.Mode()); err != nil {
				return err
			}
			if err := move(f.AbsPath, newfilepath); err != nil {
				return err
			}
		}
		moveFile(f.AbsPath, newfilepath)
	}
	if err := os.RemoveAll(oldpathInfo.AbsPath); err != nil {
		return err
	}
	return nil
}
