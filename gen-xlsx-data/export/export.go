package export

import (
	"fmt"
	"path/filepath"
	"steady/tools/gen-xlsx-data/printer"
)

func Run(g *printer.Globals) bool {
	cacheFiles := cacheFile(g)

	fmt.Println("==========开始导出==========")

	for _, file := range cacheFiles {
		fmt.Println(filepath.Base(file.FileName))
		dataModels := file.ExportData()
		if len(dataModels) == 0 {
			return false
		}

		g.File2DataModel[file.FileName] = dataModels
	}

	// 根据各种导出类型, 调用各导出器导出
	return g.Print()
}
