// +build windows

package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unsafe"

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
func move(oldpathInfo *ExtFileInfo, newpath string) error {
	if !oldpathInfo.IsDir() {
		return moveFile(oldpathInfo.AbsPath, newpath)
	}
	if err := MakeDir(newpath, oldpathInfo.Mode()); err != nil {
		return err
	}
	allFiles := oldpathInfo.Traversal()
	for _, f := range allFiles {
		newfilepath := strings.Replace(f.AbsPath, oldpathInfo.AbsPath, newpath, -1)
		if f.IsDir() {
			if err := MakeDir(newfilepath, f.Mode()); err != nil {
				return err
			}
			if err := move(f, newfilepath); err != nil {
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

// remove the file or directory, soft remove, it will move the file to trash folder
func remove(srcInfo *ExtFileInfo) (newpath string, err error) {
	// first get the free space of C:
	// C: need have at least 10G free space
	if freeSpace := mustGetFreeSpace("C:"); freeSpace < (10*1024*1024*1024 + srcInfo.RealSize) {
		return "", fmt.Errorf("Not enough free space for C:// (%s free now), maybe you can hard delete it.", transSize(freeSpace))
	}
	homePath, err := Home()
	if err != nil {
		return "", err
	}
	recyclerBin := filepath.Join(homePath, ".ft-trash")
	if err := MakeDir(recyclerBin, 0755); err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`\\([A-Z]):\\`)
	destinationPath := reg.ReplaceAllString(filepath.Join(recyclerBin, srcInfo.AbsPath), `\$1\`)
	destinationPathDir := filepath.Dir(destinationPath)
	if err := MakeDir(destinationPathDir, 0755); err != nil {
		return "", err
	}
	if err := move(srcInfo, destinationPath); err != nil {
		return "", err
	}
	return destinationPath, nil
}

// must get the free space of given volumn
func mustGetFreeSpace(volumn string) int64 {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	_, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumn))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	if err != syscall.Errno(0) {
		panic(err)
	}
	return lpFreeBytesAvailable
}
