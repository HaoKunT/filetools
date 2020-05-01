package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/liuyongshuai/goUtils"
)

type CompareOptions struct {
	Src     string
	Dst     string
	Content bool
}

var compareOptions *CompareOptions

var srcIndex = make(map[string]os.FileInfo)
var dstIndex = make(map[string]os.FileInfo)

var srcFileNum int
var srcDirNum int
var srcSize int64

var newFileNum int
var deleteFileNum int
var modifyFileNum int

func transSize(size int64) string {
	var unit = []string{"Byte", "KB", "MB", "GB", "TB", "PB"}
	var out = float64(size)
	index := 0
	for {
		if out < 1024 || index == len(unit)-1 {
			break
		}
		out /= 1024
		index++
	}
	if index == 0 {
		return fmt.Sprintf("%d %s", int(out), "Bytes")
	}
	return fmt.Sprintf("%.2f %s", out, unit[index])
}

func generateIndex(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			// fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			// return filepath.SkipDir
			srcDirNum++
			return nil
		}
		rpath, _ := filepath.Rel(root, path)
		srcFileNum++
		srcSize += info.Size()

		srcIndex[rpath] = info

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
		return
	}
	fmt.Printf(goUtils.LightRed("%d files, %d directories\n"), srcFileNum, srcDirNum)
	fmt.Printf(goUtils.LightRed("All Size: %s\n\n"), transSize(srcSize))
}

func comparefile(spath, dpath string) bool {
	sFile, err := os.Open(spath)
	if err != nil {
		return false
	}
	dFile, err := os.Open(dpath)
	if err != nil {
		return false
	}
	b := comparebyte(sFile, dFile)
	sFile.Close()
	dFile.Close()
	return b
}

//下面可以代替md5比较.
func comparebyte(sfile *os.File, dfile *os.File) bool {
	var sbyte []byte = make([]byte, 512)
	var dbyte []byte = make([]byte, 512)
	var serr, derr error
	for {
		_, serr = sfile.Read(sbyte)
		_, derr = dfile.Read(dbyte)
		if serr != nil || derr != nil {
			if serr != derr {
				return false
			}
			if serr == io.EOF {
				break
			}
		}
		if bytes.Equal(sbyte, dbyte) {
			continue
		}
		return false
	}
	return true
}

func Compare(Options *CompareOptions) error {
	compareOptions = Options
	fmt.Println("Caculating src directory...")
	generateIndex(compareOptions.Src)
	err := filepath.Walk(compareOptions.Dst, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		rpath, _ := filepath.Rel(compareOptions.Dst, path)
		if srcinfo, ok := srcIndex[rpath]; !ok {
			if IsVerbose {
				fmt.Printf(goUtils.Green("new file: %q\n"), rpath)
			}
			newFileNum++
		} else if !compareOptions.Content {
			if srcinfo.Size() != info.Size() {
				if IsVerbose {
					fmt.Printf(goUtils.Yellow("modified file: %q\n"), rpath)
				}
				modifyFileNum++
			}
		} else {
			if !comparefile(filepath.Join(compareOptions.Src, rpath), filepath.Join(compareOptions.Dst, rpath)) {
				if IsVerbose {
					fmt.Printf(goUtils.Yellow("modified file: %q\n"), rpath)
				}
				modifyFileNum++
			}
		}
		delete(srcIndex, rpath)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", compareOptions.Dst, err)
		return err
	}
	for rpath := range srcIndex {
		if IsVerbose {
			fmt.Printf(goUtils.Red("delete file: %q\n"), rpath)
		}
		deleteFileNum++
	}
	fmt.Printf(goUtils.Green("\n\nnew file: %d\n"), newFileNum)
	fmt.Printf(goUtils.Yellow("modified file: %d\n"), modifyFileNum)
	fmt.Printf(goUtils.Red("delete file: %d\n"), deleteFileNum)
	return nil
}
