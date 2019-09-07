package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"
)

type CopyOptions struct {
	Src         string
	Dst         string
	IsDir       bool
	CreateDir   bool
	Verbose     bool
	ProgressBar bool
}

var copyOptions *CopyOptions
var pgbSize *pb.ProgressBar
var pgbNum *pb.ProgressBar
var dirCopy bool

func Copy(options *CopyOptions) error {
	copyOptions = options

	info, err := os.Stat(options.Src)
	if err != nil {
		return err
	}

	if info.IsDir() && !options.IsDir {
		return errors.New("Source is a directory, please use -r flag")
	}
	if info.IsDir() {
		dirCopy = true
		if copyOptions.Verbose && copyOptions.ProgressBar {
			return errors.New("The verbose flag can't be used with progress bar")
		}
		dstInfo, err := os.Stat(options.Dst)
		if err != nil {
			return err
		}
		if !dstInfo.IsDir() {
			return fmt.Errorf("Destination %s is not a directory", options.Dst)
		}
		srcAbs, err := filepath.Abs(info.Name())
		if err != nil {
			return err
		}
		srcDirName := filepath.Base(srcAbs)
		dstDir := filepath.Join(copyOptions.Dst, srcDirName)
		dstAbs, err := filepath.Abs(dstDir)
		if err != nil {
			return err
		}
		if dstAbs == srcAbs {
			return fmt.Errorf("Source and destination are the same")
		}

		if strings.HasPrefix(dstAbs, srcAbs) {
			return fmt.Errorf("%s is sub-directory of source %s", dstDir, info.Name())
		}
		if !IsExisted(dstAbs) && !copyOptions.CreateDir {
			return fmt.Errorf("%s: No such file or directory", dstAbs)
		}
		if err := MakeDir(dstAbs, info.Mode()); err != nil {
			return err
		}
		n, err := copyDir(srcAbs, dstAbs)
		if err != nil {
			return err
		}
		fmt.Printf("Copy %s to %s, Size: %s(%d bytes)\n", options.Src, options.Dst, transSize(n), n)
	} else {
		n, err := copyFile(options.Src, options.Dst)
		if err != nil {
			return err
		}
		fmt.Printf("Copy %s to %s, Size: %s(%d bytes)\n", options.Src, options.Dst, transSize(n), n)
	}
	return nil
}

func copyFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	if copyOptions.ProgressBar && !dirCopy {
		pgbSize = pb.New64(fi.Size()).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 100)
		pgbSize.ShowSpeed = true
		pgbSize.ShowCounters = true
		pgbSize.ShowTimeLeft = true
		pgbSize.Start()
		reader := pgbSize.NewProxyReader(srcFile)
		n, err := io.Copy(desFile, reader)
		pgbSize.Finish()
		if err != nil {
			return 0, err
		}
		return n, nil
	} else if copyOptions.ProgressBar && dirCopy {
		reader := pgbSize.NewProxyReader(srcFile)
		return io.Copy(desFile, reader)
	}
	return io.Copy(desFile, srcFile)
}

func copyDir(srcDir, dstDir string) (int64, error) {
	var allsize int64
	if copyOptions.ProgressBar {
		err := calculateDir(srcDir)
		if err != nil {
			return 0, err
		}
		pgbSize = pb.New64(srcSize).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 100)
		pgbNum = pb.New(srcFileNum)
		pgbSize.ShowSpeed = true
		pgbSize.ShowCounters = true
		pgbSize.ShowTimeLeft = true
		pgbNum.ShowSpeed = true
		pgbNum.ShowCounters = true
		pgbNum.ShowTimeLeft = true
		pgbSize.Start()
		pgbNum.Start()
	}
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == srcDir {
			return nil
		}
		dstNewPath := strings.Replace(path, srcDir, dstDir, -1)

		if info.IsDir() {
			if err := MakeDir(dstNewPath, info.Mode()); err != nil {
				return err
			}
		} else {
			n, err := copyFile(path, dstNewPath)
			if err != nil {
				return err
			}
			allsize += n
			if copyOptions.Verbose {
				fmt.Printf("%s -> %s : %s\n", path, dstNewPath, transSize(n))
			}
			if copyOptions.ProgressBar {
				pgbNum.Increment()
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	if copyOptions.ProgressBar {
		pgbSize.Finish()
		pgbNum.Finish()
	}
	return allsize, nil
}

func IsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string, perm os.FileMode) error {
	if !IsExisted(dir) {
		if err := os.MkdirAll(dir, perm); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}
