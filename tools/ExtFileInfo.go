package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type ExtFileInfo struct {
	os.FileInfo
	AbsPath      string
	RootPath     string
	RelativePath string
	// BasePath     string
	RealSize int64
	DirNum   int
	FileNum  int
	Children []*ExtFileInfo
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
	if !exinfo.IsDir() {
		exinfo.RealSize = exinfo.Size()
		exinfo.DirNum = 0
		exinfo.FileNum = 1
		exinfo.Children = nil
		return &exinfo, nil
	}
	exinfo.DirNum = 1
	errs := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go buildTree(&exinfo, errs, &wg)
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
		return nil, fmt.Errorf("NewExtInfo error: %v", errs)
	}
	return &exinfo, nil
}

// build the file tree
func buildTree(extinfo *ExtFileInfo, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
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
		extinfo.Children = append(extinfo.Children, &cextinfo)
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
			go buildTree(&cextinfo, errs, wg)
		}
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
