package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type InfoOptions struct {
	FileName string
	Crypto   bool
}

var infoOptions *InfoOptions

func Info(options *InfoOptions) error {
	infoOptions = options

	info, err := os.Stat(options.FileName)
	if err != nil {
		return err
	}

	absName, err := filepath.Abs(options.FileName)
	if err != nil {
		return err
	}

	fmt.Println("General:")
	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Absolute path: %s\n", absName)

	if info.IsDir() {
		fmt.Printf("Type: directory\n")
		if err := calculateDir(infoOptions.FileName); err != nil {
			return err
		}
	} else {
		kind, err := filetype.MatchFile(info.Name())
		if err != nil {
			return err
		}
		fmt.Printf("Type: %s\n", kind.Extension)
		if kind != filetype.Unknown {
			fmt.Printf("MIME-type: %s\n", kind.MIME.Value)
		}
		fmt.Printf("Size: %s (%d bytes)\n", transSize(info.Size()), info.Size())
	}
	fmt.Printf("Mode: %s\n", info.Mode())
	fmt.Printf("ModTime: %s\n", info.ModTime())
	fmt.Printf("Is Regular: %v\n", info.Mode().IsRegular())

	if infoOptions.Crypto && !info.IsDir() {
		// crypto
		fmt.Println()
		fmt.Println("Crypto")
		// md5
		md5V, err := md5f(absName)
		if err != nil {
			return err
		}
		fmt.Printf("MD5: %s\n", md5V)

		// sha1
		sha1V, err := sha1f(absName)
		if err != nil {
			return err
		}
		fmt.Printf("SHA1: %s\n", sha1V)

		// sha256
		sha256V, err := sha256f(absName)
		if err != nil {
			return err
		}
		fmt.Printf("SHA256: %s\n", sha256V)
	}

	return nil
}

func calculateDir(root string) error {
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
		srcFileNum++
		srcSize += info.Size()

		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("Include: %d files, %d directories\n", srcFileNum, srcDirNum)
	fmt.Printf("All Size: %s\n", transSize(srcSize))
	return nil
}

func md5f(fName string) (string, error) {
	f, e := os.Open(fName)
	if e != nil {
		return "", e
	}
	defer f.Close()
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return "", e
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func sha1f(fName string) (string, error) {
	f, e := os.Open(fName)
	if e != nil {
		return "", e
	}
	defer f.Close()
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return "", e
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func sha256f(fName string) (string, error) {
	f, e := os.Open(fName)
	if e != nil {
		return "", e
	}
	defer f.Close()
	h := sha256.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return "", e
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
