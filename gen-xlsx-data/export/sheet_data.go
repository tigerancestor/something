package export

import (
	"encoding/json"
	"fmt"
	"steady/tools/gen-xlsx-data/model"
	"steady/tools/gen-xlsx-data/util"
	"strings"
)

type DataSheet struct {
	*Sheet
	Header *DataHeader
}

func (p *DataSheet) Valid() bool {

	name := strings.TrimSpace(p.Sheet.Name)
	if name != "" && name[0] == '#' {
		return false
	}

	return p.GetCellData(0, 0) != ""
}

type SheetExportInfo struct {
	Name    string
	Package string
}

func (p *DataSheet) Export(file *File, dataModel *model.DataModel, dataSheet *DataSheet) bool {
	// 解析表头
	exportInfo := p.GetCellData(DataSheetHeaderExportInfo, 0) //第0行0列描述导出的相关信息
	sheetExportInfo := &SheetExportInfo{}
	if err := json.Unmarshal([]byte(exportInfo), sheetExportInfo); err != nil {
		panic(err)
	}

	dataSheet.Header.Name = sheetExportInfo.Name
	dataSheet.Header.Package = sheetExportInfo.Package
	fmt.Printf("Name:%s Package:%s \n", sheetExportInfo.Name, sheetExportInfo.Package)

	// 解析字段行
	i := 0
	for {
		fieldName := p.GetCellData(DataSheetHeaderFieldName, i)
		fieldType := p.GetCellData(DataSheetHeaderFieldType, i)
		fieldComment := p.GetCellData(DataSheetHeaderFieldComment, i)
		i++
		if fieldName == "" && fieldType == "" && fieldComment == "" {
			break
		}
		if fieldName == "" {
			continue
		}

		he := &model.DataHeaderElement{
			FieldName: fieldName,
			FieldType: p.GetCellData(DataSheetHeaderFieldType, i-1),
			Comment:   p.GetCellData(DataSheetHeaderFieldComment, i-1),
			Col:       i - 1,
		}
		p.Header.HeaderFields = append(p.Header.HeaderFields, he)
	}

	// 是否继续读行
	var readingLine bool = true

	var meetEmptyLine bool

	var warningAfterEmptyLineDataOnce bool

	// 遍历每一行
	for p.Row = DataSheetHeaderDataBegin; readingLine; p.Row++ {

		// 整行都是空的
		if p.IsFullRowEmpty(p.Row, p.Header.MaxCol()) {

			// 再次碰空行, 表示确实是空的
			if meetEmptyLine {
				break

			} else {
				meetEmptyLine = true
			}

			continue

		} else {

			//已经碰过空行, 这里又碰到数据, 说明有人为隔出的空行, 做warning提醒, 防止数据没导出
			if meetEmptyLine && !warningAfterEmptyLineDataOnce {
				r, _ := p.GetRC()

				fmt.Printf("%s %s|%s(%s)", "空行", p.file.FileName, p.Name, util.ConvR1C1toA1(r, 1))

				warningAfterEmptyLineDataOnce = true
			}

			// 曾经有过空行, 即便现在不是空行也没用, 结束
			if meetEmptyLine {
				break
			}

		}

		line := model.NewLineData()

		// 遍历每一列
		for p.Column = 0; p.Column <= dataSheet.Header.MaxCol(); p.Column++ {
			headEle := dataSheet.Header.Field(p.Column)
			if headEle == nil {
				// 无效表头
				continue
			}

			op := p.processLine(headEle, line)

			if op == lineOpContinue {
				continue
			} else if op == lineOpBreak {
				break
			}

		}

		dataModel.Add(line)
	}

	dataModel.Name = p.Header.Name
	dataModel.Package = p.Header.Package
	dataModel.HeaderFields = p.Header.HeaderFields

	return true
}

const (
	lineOpNone = iota
	lineOpBreak
	lineOpContinue
)

func (p *DataSheet) processLine(headEle *model.DataHeaderElement, line *model.LineData) int {
	//fmt.Printf("processLine fieldName:%s\n", fieldDef.Name)
	// 数据大于列头时, 结束这个列
	if headEle == nil {
		return lineOpBreak
	}

	// #开头表示注释, 跳过
	if strings.Index(headEle.FieldName, "#") == 0 {
		return lineOpContinue
	}

	var rawValue string
	// 浮点数按本来的格式输出
	rawValue = p.GetCellData(p.Row, p.Column)

	r, c := p.GetRC()

	line.Add(&model.FieldValue{
		FieldName:    headEle.FieldName,
		FieldType:    headEle.FieldType,
		FieldComment: headEle.Comment,
		RawValue:     rawValue,
		SheetName:    p.Name,
		FileName:     p.file.FileName,
		R:            r,
		C:            c,
	})

	return lineOpNone
}

func newDataSheet(sheet *Sheet) *DataSheet {
	return &DataSheet{
		Sheet:  sheet,
		Header: &DataHeader{},
	}
}
