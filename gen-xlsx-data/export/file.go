package export

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"steady/tools/gen-xlsx-data/model"
)

// table信息
type File struct {
	FileName   string
	coreFile   *xlsx.File
	dataSheets []*DataSheet
}

func (p *File) ExportData() (dataModels []*model.DataModel) {
	for _, rawSheet := range p.coreFile.Sheets {
		dSheet := newDataSheet(NewSheet(p, rawSheet))
		if dSheet.GetCellData(0, 0) == "" { // 空表跳过
			continue
		}
		sheetExportInfo := &SheetExportInfo{}
		if err := json.Unmarshal([]byte(dSheet.GetCellData(0, 0)), sheetExportInfo); err != nil {
			continue
		}
		p.dataSheets = append(p.dataSheets, dSheet)
	}

	for _, d := range p.dataSheets {

		fmt.Printf("ExportData %s\n", d.Name)
		dataModel := model.NewDataModel()
		dataModel.SheetName = d.Name
		if !d.Export(p, dataModel, d) {
			return
		}
		dataModels = append(dataModels, dataModel)
	}

	return
}

func NewFile(filename string) *File {
	p := &File{
		FileName: filename,
	}

	var err error
	p.coreFile, err = xlsx.OpenFile(filename)

	if err != nil {
		fmt.Sprintf("无法找到文件 %s", filename)
		return nil
	}

	return p
}
