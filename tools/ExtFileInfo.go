package tools

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/Re-volution/sizestruct"
	"github.com/panjf2000/ants/v2"
)

type ExtFileInfo struct {
	os.FileInfo
	AbsPath      string
	RootPath     string
	RelativePath string
	BasePath     string
	RealSize     int64
	DirNum       int
	FileNum      int
	Parent       *ExtFileInfo
	Children     []*ExtFileInfo
	lock         sync.RWMutex
}

// NewExtInfo build file tree, the path is root path
func NewExtInfo(path string) (*ExtFileInfo, error) {
	var exinfo ExtFileInfo
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	exinfo.FileInfo = info

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	exinfo.RootPath = absPath
	exinfo.AbsPath = absPath
	exinfo.RelativePath = "."
	exinfo.BasePath = filepath.Base(absPath)
	exinfo.Parent = nil
	if !exinfo.IsDir() {
		exinfo.RealSize = exinfo.Size()
		exinfo.DirNum = 0
		exinfo.FileNum = 1
		exinfo.Children = nil
		return &exinfo, nil
	}
	p, err := ants.NewPool(math.MaxInt32, ants.WithExpiryDuration(1*time.Second), ants.WithNonblocking(true))
	if err != nil {
		return nil, err
	}
	defer p.Release()
	errs := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	err = p.Submit(func() {
		buildTree(&exinfo, errs, &wg, p)
		wg.Done()
	})
	if err != nil {
		errs <- err
	}
	// go func() {
	// 	timer := time.NewTimer(1 * time.Second)
	// 	for {
	// 		select {
	// 		case <-timer.C:
	// 			fmt.Printf("Cap: %d, Running: %d, Free: %d, wg: %v\n", p.Cap(), p.Running(), p.Free(), wg)
	// 			timer.Reset(1 * time.Second)
	// 		}
	// 	}
	// }()
	var isError bool
	var errsList []error
	go func() {
		wg.Wait()
		close(errs)
	}()
	for err := range errs {
		isError = true
		errsList = append(errsList, err)
	}
	if isError {
		return nil, fmt.Errorf("NewExtInfo error: %v", errsList)
	}
	p.Release()
	return &exinfo, nil
}

// build the file tree
// the errs will transform the error to root, wg also wait the build, parentSync is wait for the dir to calculate the size
func buildTree(extinfo *ExtFileInfo, errs chan<- error, wg *sync.WaitGroup, p *ants.Pool) {
	// Concurrency <- struct{}{}
	extinfo.lock.Lock()
	defer func() {
		extinfo.lock.Unlock()
		// <-Concurrency
	}()
	filelist, err := ioutil.ReadDir(extinfo.AbsPath)
	if err != nil {
		errs <- err
	}
	for _, file := range filelist {
		var cextinfo ExtFileInfo
		cextinfo.FileInfo = file
		cextinfo.AbsPath = filepath.Join(extinfo.AbsPath, file.Name())
		cextinfo.RootPath = extinfo.RootPath
		cextinfo.RelativePath = filepath.Join(extinfo.RelativePath, file.Name())
		cextinfo.BasePath = filepath.Base(cextinfo.AbsPath)
		extinfo.Children = append(extinfo.Children, &cextinfo)
		cextinfo.Parent = extinfo
		if !cextinfo.IsDir() {
			cextinfo.RealSize = extinfo.Size()
			cextinfo.DirNum = 0
			cextinfo.FileNum = 1
			cextinfo.Children = nil
			extinfo.RealSize += cextinfo.Size()
			extinfo.FileNum++
		} else {
			extinfo.DirNum++
			wg.Add(1)
			err = p.Submit(func() {
				buildTree(&cextinfo, errs, wg, p)
				wg.Done()
			})
			if err != nil {
				errs <- err
			}
		}
	}
	if extinfo.Parent != nil {
		extinfo.Parent.lock.Lock()
		changeHandler(extinfo.Parent, extinfo)
		extinfo.Parent.lock.Unlock()
	}
}

func changeHandler(pinfo, info *ExtFileInfo) {
	pinfo.RealSize += info.RealSize
	pinfo.FileNum += info.FileNum
	pinfo.DirNum += info.DirNum
	if pinfo.Parent != nil {
		pinfo.Parent.lock.Lock()
		changeHandler(pinfo.Parent, info)
		pinfo.Parent.lock.Unlock()
	}
}

// GetAllFiles get all files in directory, all path will be abs Path
func (extinfo *ExtFileInfo) GetAllFiles() []string {
	infos := extinfo.Traversal()
	var allFiles []string
	for _, fi := range infos {
		if !fi.IsDir() {
			allFiles = append(allFiles, fi.AbsPath)
		}
	}
	return allFiles
}

// GetAllFilesDirs get all files and directories in directory, all path will be abs Path
func (extinfo *ExtFileInfo) GetAllFilesDirs() []string {
	infos := extinfo.Traversal()
	var allFiles []string
	for _, fi := range infos {
		allFiles = append(allFiles, fi.AbsPath)
	}
	return allFiles
}

// Traversal the directory
func (extinfo *ExtFileInfo) Traversal() []*ExtFileInfo {
	var infos []*ExtFileInfo
	for _, info := range extinfo.Children {
		infos = append(infos, info)
		if info.IsDir() {
			infos = append(infos, info.Traversal()...)
		}
	}
	return infos
}

// TraversalChan means Traversal the directory (channel)
// buf means the channel buffer size
func (extinfo *ExtFileInfo) TraversalChan(buf int) <-chan *ExtFileInfo {
	if buf < 0 {
		buf = 0
	}
	ch := make(chan *ExtFileInfo, buf)
	go traversalChan(extinfo, ch, true)
	return ch
}

func traversalChan(extinfo *ExtFileInfo, ch chan<- *ExtFileInfo, isclose bool) {
	for _, info := range extinfo.Children {
		ch <- info
		if info.IsDir() {
			traversalChan(info, ch, false)
		}
	}
	if isclose {
		close(ch)
	}
}

// RealSizeOfExtFileinfo returns the size of the struct
func RealSizeOfExtFileinfo(extinfo *ExtFileInfo) int64 {
	return int64(sizestruct.SizeOf(extinfo))
}

func filedeep(absPath string) int {
	if runtime.GOOS == "windows" {
		return len(strings.Split(absPath, "\\"))
	}
	return len(strings.Split(absPath, "/"))
}
