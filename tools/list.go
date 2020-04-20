package tools

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/liuyongshuai/goUtils"
)

type ListOptions struct {
	Path   string
	SortBy []string
	Limit  int
}

type osSort []ExtFileInfo

type sortTagFunc func(a, b os.FileInfo) status

var sortMapFuncs map[string]sortTagFunc

var listOptions *ListOptions
var listIndex = make(osSort, 0)
var sortByParams [][2]string
var fileNum int
var showFileNum int

func initList() {
	sortMapFuncs = map[string]sortTagFunc{
		"name": sortByNameFunc,
		"size": sortBySizeFunc,
	}
}

func sortByNameFunc(a, b os.FileInfo) status {
	if a.Name() < b.Name() {
		return ascending
	} else if a.Name() > b.Name() {
		return descending
	} else {
		return Equal
	}
}

func sortBySizeFunc(a, b os.FileInfo) status {
	if a.Size() < b.Size() {
		return ascending
	} else if a.Size() > b.Size() {
		return descending
	} else {
		return Equal
	}
}

func List(options *ListOptions) error {
	listOptions = options
	initList()
	pathName := listOptions.Path
	info, err := os.Stat(pathName)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("The path %s is not a directory", pathName)
	}
	if listOptions.Limit < 0 {
		return fmt.Errorf("The limit %d can't be negative", listOptions.Limit)
	}
	var isSort bool

	if len(listOptions.SortBy) != 0 {
		isSort = true
		sortByParams, err = parseSortBy()
		if err != nil {
			return err
		}
	}

	for walker := range ParWalk(pathName, 1000) {
		if walker.Err != nil {
			fmt.Println(goUtils.Red(fmt.Sprintf("Acess to the file/directory %s error: %s\n", walker.AbsPath, walker.Err)))
			continue
		}
		if !walker.IsDir() {
			if !isSort {
				if showFileNum >= listOptions.Limit {
					break
				}

				fmt.Printf("%s  %s  %10s  %s\n", walker.Mode(), walker.ModTime().Format("2006-01-02 15:04:05"), transSize(walker.Size()), walker.RelPath)
				showFileNum++
			} else {
				listIndex = append(listIndex, ExtFileInfo{
					FileInfo:     walker.FileInfo,
					RelativePath: walker.RelPath,
					// BasePath:     path,
				},
				)
				if fileNum > listOptions.Limit {
					sort.Sort(listIndex)
					listIndex = listIndex[:listOptions.Limit]
				}

			}
			fileNum++
		}
	}

	// err = filepath.Walk(pathName, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !info.IsDir() {
	// 		if !isSort {
	// 			if showFileNum >= listOptions.Limit {
	// 				return breakWalkError
	// 			}
	// 			rpath, err := filepath.Rel(pathName, path)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			fmt.Printf("%s  %s  %10s  %s\n", info.Mode(), info.ModTime().Format("2006-01-02 15:04:05"), transSize(info.Size()), rpath)
	// 			showFileNum++
	// 		} else {
	// 			absPath, err := filepath.Abs(path)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			listIndex = append(listIndex, ExtFileInfo{
	// 				FileInfo:     info,
	// 				AbsPath:      absPath,
	// 				RelativePath: path,
	// 				// BasePath:     path,
	// 			},
	// 			)
	// 		}
	// 		fileNum++
	// 	}
	// 	return nil
	// })
	// if err == breakWalkError {
	// 	return nil
	// }
	// if err != nil {
	// 	return err
	// }

	if isSort {
		sort.Sort(listIndex)
		for _, oS := range listIndex {
			fmt.Printf("%s  %s  %10s  %s\n", oS.Mode(), oS.ModTime().Format("2006-01-02 15:04:05"), transSize(oS.Size()), oS.RelativePath)
		}
	}
	return nil
}

func parseSortBy() ([][2]string, error) {
	reg, err := regexp.Compile(`^(?:name|size)(?:\:(?:desc|asc|a|d))?$`)
	if err != nil {
		return nil, err
	}
	for _, nameAndFlags := range listOptions.SortBy {
		if !reg.MatchString(nameAndFlags) {
			return nil, fmt.Errorf("Invalid --sortBy value: %s", nameAndFlags)
		}
		nameAndFlagsSlices := strings.Split(nameAndFlags, ":")
		var nameAndFlagsSlices2 [2]string
		nameAndFlagsSlices2[0] = nameAndFlagsSlices[0]
		if len(nameAndFlagsSlices) == 1 {
			nameAndFlagsSlices2[1] = "asc"
		} else {
			nameAndFlagsSlices2[1] = nameAndFlagsSlices[1]
		}
		sortByParams = append(sortByParams, nameAndFlagsSlices2)
	}
	return sortByParams, nil
}

func (oS osSort) Len() int {
	return len(oS)
}

func (oS osSort) Less(i, j int) bool {
	for _, nameAndFlags := range sortByParams {
		name := nameAndFlags[0]
		order := nameAndFlags[1]
		stat := sortMapFuncs[name](oS[i], oS[j])
		if stat == Equal {
			continue
		}
		if strings.HasPrefix(order, "a") {
			return judgeReturn(stat)
		} else if strings.HasPrefix(order, "d") {
			return !judgeReturn(stat)
		} else {
			panic(fmt.Sprintf("The sort flag error: the flag %s error", nameAndFlags))
		}
	}
	return true
}

func judgeReturn(s status) bool {
	if s != ascending {
		return false
	}
	return true
}

func (oS osSort) Swap(i, j int) {
	oS[i], oS[j] = oS[j], oS[i]
}
