package export

import (
	"steady/tools/gen-xlsx-data/printer"
	"strings"
	"sync"
)

func getFileList(g *printer.Globals) (ret []string) {
	// 合并类型
	for _, in := range g.InputFileList {

		inputFile := in.(string)

		mergeFileList := strings.Split(inputFile, "+")

		for _, fileName := range mergeFileList {
			ret = append(ret, fileName)
		}
	}

	return
}

func cacheFile(g *printer.Globals) (fileObjByName map[string]*File) {

	var fileObjByNameGuard sync.Mutex
	fileObjByName = map[string]*File{}

	//fmt.Println("==========cacheFile==========")

	fileList := getFileList(g)

	var task sync.WaitGroup
	task.Add(len(fileList))

	for _, filename := range fileList {

		go func(xlsxFileName string) {

			//fmt.Println(filepath.Base(xlsxFileName))
			file := NewFile(xlsxFileName)

			fileObjByNameGuard.Lock()
			fileObjByName[xlsxFileName] = file
			fileObjByNameGuard.Unlock()

			task.Done()

		}(filename)

	}

	task.Wait()

	return
}
