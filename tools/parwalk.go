// This file contain the parralle walk, it works like filepath.Walk
package tools

import (
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

type ParWalker struct {
	os.FileInfo
	AbsPath string
	RelPath string
	Err     error // if the walker meets the error, it's error, nor it will be nil
}

// ParWalk is designed to walk the directory asynchronously.
// In fact the function NewExtInfo will build the directory tree asynchronously. But it need so much memory when the
// directory has many files and directories. When we not need to use the sub-directory information,
// we will use the ParWalk to walk the directory.
// buf argument is the size of channel ParWalker, it will limit the amount of concurrency
func ParWalk(path string, buf int) <-chan ParWalker {
	if buf <= 0 {
		buf = 10
	}
	p, err := ants.NewPool(math.MaxInt32, ants.WithExpiryDuration(1*time.Second), ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	walker := make(chan ParWalker, buf)
	var wg sync.WaitGroup
	wg.Add(1)
	p.Submit(func() {
		defer wg.Done()
		filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
			wg.Add(1)
			p.Submit(func() {
				defer wg.Done()
				if err != nil {
					walker <- ParWalker{
						nil,
						"",
						"",
						err,
					}
					return
				}
				absPath, err1 := filepath.Abs(subpath)
				rpath, err2 := filepath.Rel(path, subpath)
				if err1 != nil || err2 != nil {
					walker <- ParWalker{
						nil,
						"",
						"",
						err,
					}
					return
				}
				walker <- ParWalker{
					info,
					absPath,
					rpath,
					nil,
				}
			})
			return nil
		})
	})
	go func() {
		wg.Wait()
		close(walker)
	}()
	return walker
}
