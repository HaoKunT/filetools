package tools

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

type MoveOptions struct {
	Src    string
	Dst    string
	isMove bool
}

var moveOptions *MoveOptions

// Move source to directory or Rename source to dst, just like `mv`
func Move(options *MoveOptions) error {
	moveOptions = options
	extInfo, err := NewExtInfo(options.Src)
	if err != nil {
		return err
	}
	absDst, err := filepath.Abs(moveOptions.Dst)
	if err != nil {
		return err
	}
	var destPath string
	// judge the Dst is a filen or a directory, if file, it's rename, if directory, it's move
	// if the Dst is end of slash, it means it's a directory, the source will be move into it
	// for safety, first we should judge whether the Dst Exists, if it doesn't exist and end
	// with slash, error, if not end with slash, it should be a file, if it existed and end with slash
	// it's directory, if not end with slash, orverwrite it
	if !IsExisted(moveOptions.Dst) {
		if strings.HasSuffix(moveOptions.Dst, "/") {
			return fmt.Errorf("No such file or directory: destination %s", absDst)
		}
	} else {
		if strings.HasSuffix(moveOptions.Dst, "/") {
			moveOptions.isMove = true
		}
	}
	moveOptions.isMove = strings.HasSuffix(moveOptions.Dst, "/")
	if moveOptions.isMove {
		destPath = filepath.Join(absDst, extInfo.BasePath)
	} else {
		// if the Dst is not end with slash, it means it's a file need to be renamed
		destPath = absDst
	}
	if extInfo.AbsPath == destPath {
		return errors.New("Source and Destination are the same")
	}
	if strings.HasPrefix(destPath, extInfo.AbsPath) {
		return fmt.Errorf("%s is sub-directory of %s", destPath, extInfo.AbsPath)
	}
	err = move(extInfo, destPath)
	if err != nil {
		return err
	}
	fmt.Printf("Move %s -> %s (%s)\n", extInfo.AbsPath, destPath, transSize(extInfo.RealSize))
	return nil
}
