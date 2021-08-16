package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"
)

type CopyOptions struct {
	Src          string
	Dst          string
	IsDir        bool
	CreateDir    bool
	Verbose      bool
	ProgressBar  bool
	ParallelNums int
}

var copyOptions *CopyOptions
var pgbSize *pb.ProgressBar
var pgbNum *pb.ProgressBar
var pgbPool *pb.Pool
var dirCopy bool

func Copy(options *CopyOptions) error {
	copyOptions = options

	info, err := NewExtInfo(copyOptions.Src)
	if err != nil {
		return err
	}

	// first the directory judge
	if info.IsDir() && !options.IsDir {
		return errors.New("Source is a directory, please use -r flag")
	}
	// if the directoy
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
		dstDir := filepath.Join(copyOptions.Dst, info.BasePath)
		dstAbs, err := filepath.Abs(dstDir)
		if err != nil {
			return err
		}
		if dstAbs == info.AbsPath {
			return fmt.Errorf("Source and destination are the same")
		}

		// TODO: /home/xxx/filetools and /home/xxx/filetools-copy
		if strings.HasPrefix(dstAbs, info.AbsPath) {
			return fmt.Errorf("%s is sub-directory of source %s", dstAbs, info.AbsPath)
		}
		if !IsExisted(dstAbs) && !copyOptions.CreateDir {
			return fmt.Errorf("%s: No such file or directory", dstAbs)
		}
		if err := MakeDir(dstAbs, info.Mode()); err != nil {
			return err
		}
		fmt.Printf("Include %d files, %d directories\n", info.FileNum, info.DirNum)
		fmt.Printf("Size %s (%d bytes)\n", transSize(info.RealSize), info.RealSize)
		n, err := copyDir(info, dstAbs)
		if err != nil {
			return err
		}
		fmt.Printf("Copy %s to %s, Size: %s(%d bytes)\n", options.Src, options.Dst, transSize(n), n)
	} else {
		srcName := filepath.Base(options.Src)
		dstName := filepath.Join(options.Dst, srcName)
		srcAbs, err := filepath.Abs(options.Src)
		if err != nil {
			return err
		}
		dstAbs, err := filepath.Abs(dstName)
		if err != nil {
			return err
		}
		if dstAbs == srcAbs {
			return fmt.Errorf("Source and destination are the same file")
		}
		n, err := copyFile(srcAbs, dstAbs)
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

func copyDir(srcDir *ExtFileInfo, dstDir string) (int64, error) {
	var allsize int64
	var err error
	if copyOptions.ProgressBar {
		pgbSize = pb.New64(srcDir.RealSize).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 100)
		pgbNum = pb.New(srcDir.FileNum)
		pgbSize.ShowSpeed = true
		pgbSize.ShowCounters = true
		pgbSize.ShowTimeLeft = true
		pgbNum.ShowSpeed = true
		pgbNum.ShowCounters = true
		pgbNum.ShowTimeLeft = true

		pgbPool, err = pb.StartPool(pgbSize, pgbNum)
		if err != nil {
			return 0, err
		}
	}
	errC := make(chan error, 10)
	var wg sync.WaitGroup
	var errsList []error
	go func() {
		for err := range errC {
			errsList = append(errsList, err)
		}
	}()

	for info := range srcDir.TraversalChan(copyOptions.ParallelNums) {
		if info.AbsPath == srcDir.AbsPath {
			continue
		}
		dstNewPath := strings.Replace(info.AbsPath, srcDir.AbsPath, dstDir, -1)

		if info.IsDir() {
			if err := MakeDir(dstNewPath, info.Mode()); err != nil {
				errC <- err
			}
		} else {
			wg.Add(1)
			go func(i *ExtFileInfo, newPath string) {
				Concurrency <- struct{}{}
				defer func() {
					wg.Done()
					<-Concurrency
				}()
				n, err := copyFile(i.AbsPath, newPath)
				if err != nil {
					errC <- err
				}
				atomic.AddInt64(&allsize, n)
				if copyOptions.Verbose {
					fmt.Printf("%s -> %s : %s\n", i.AbsPath, newPath, transSize(n))
				}
				if copyOptions.ProgressBar {
					pgbNum.Increment()
				}
			}(info, dstNewPath)
		}
	}

	// err = filepath.Walk(srcDir.AbsPath, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		errC <- err
	// 		return err
	// 	}
	// 	if path == srcDir.AbsPath {
	// 		return nil
	// 	}
	// 	dstNewPath := strings.Replace(path, srcDir.AbsPath, dstDir, -1)

	// 	if info.IsDir() {
	// 		if err := MakeDir(dstNewPath, info.Mode()); err != nil {
	// 			errC <- err
	// 			return err
	// 		}
	// 	} else {
	// 		wg.Add(1)
	// 		go func() {
	// 			defer wg.Done()
	// 			n, err := copyFile(path, dstNewPath)
	// 			if err != nil {
	// 				errC <- err
	// 			}
	// 			allsize = atomic.AddInt64(&allsize, n)
	// 			if copyOptions.Verbose {
	// 				fmt.Printf("%s -> %s : %s\n", path, dstNewPath, transSize(n))
	// 			}
	// 			if copyOptions.ProgressBar {
	// 				pgbNum.Increment()
	// 			}
	// 		}()
	// 	}
	// 	return nil
	// })
	wg.Wait()
	close(errC)
	if copyOptions.ProgressBar {
		pgbSize.Finish()
		pgbNum.Finish()
		pgbPool.Stop()
	}
	if len(errsList) != 0 {
		return 0, fmt.Errorf("copy dir error: %v", errsList)
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
