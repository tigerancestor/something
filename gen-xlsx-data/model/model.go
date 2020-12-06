package model

import (
	"strings"
)

type FieldValue struct {
	RawValue     string
	Value        string
	R            int
	C            int
	SheetName    string
	FileName     string
	FieldName    string
	FieldType    string
	FieldComment string
}

// 对应record
type LineData struct {
	Values []*FieldValue
}

func (p *LineData) Add(fv *FieldValue) {

	p.Values = append(p.Values, fv)
}

func NewLineData() *LineData {
	return new(LineData)
}

type DataHeaderElement struct {
	FieldName string
	FieldType string
	Comment   string
	Col       int
}

// 对应sheet
type DataModel struct {
	Lines        []*LineData
	HeaderFields []*DataHeaderElement

	// 按字段名分组索引字段, 字段不重复
	HeaderByName map[string]*DataHeaderElement

	Name      string
	Package   string
	SheetName string
	Imports   string
}

func (p *DataModel) Add(data *LineData) {
	p.Lines = append(p.Lines, data)
}

func NewDataModel() *DataModel {
	return new(DataModel)
}

func (p *DataModel) TranslateFirstLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func (p *DataModel) TranslateFirstUpper(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func (p *DataModel) TranslateGolangFieldType(fieldType string) string {
	switch fieldType {
	case "int":
		return "int32"
	case "date", "dateTime":
		return "time.Time"
	case "giveItems":
		return "[]*msg.SGiveItem"
	case "giveMoneys":
		return "[]*msg.SGiveMoney"
	default:
		return fieldType
	}
}

func (p *DataModel) TranslateTsFieldType(fieldType string) string {
	if strings.Contains(fieldType, "int") || strings.Contains(fieldType, "float") {
		return "number"
	} else {
		return "string"
	}
}

func (p *DataModel) StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				// 如果是小写才转大写
				if i == 0 && vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func (p *DataModel) StrStripUnderline(str string) string {
	return strings.ToLower(strings.Replace(str, "_", "", 0))
}
