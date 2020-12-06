package export

import "steady/tools/gen-xlsx-data/model"

const (
	// 信息所在的行
	DataSheetHeaderExportInfo   = 0 // 字段名(对应proto)
	DataSheetHeaderFieldType    = 1 // 字段类型
	DataSheetHeaderFieldName    = 2 // 字段名
	DataSheetHeaderFieldComment = 3 // 字段注释
	DataSheetHeaderDataBegin    = 4 // 数据开始
)

type DataHeader struct {
	// 按排列的, 字段不重复
	HeaderFields []*model.DataHeaderElement

	// 按字段名分组索引字段, 字段不重复
	HeaderByName map[string]*model.DataHeaderElement

	Name    string
	Package string
}

func (p *DataHeader) Field(col int) *model.DataHeaderElement {
	for _, headerField := range p.HeaderFields {
		if headerField.Col == col {
			return headerField
		}
	}

	return nil
}

func (p *DataHeader) MaxCol() int {
	return p.HeaderFields[len(p.HeaderFields)-1].Col
}

func newDataHeadSheet() *DataHeader {

	return &DataHeader{
		HeaderByName: make(map[string]*model.DataHeaderElement),
	}
}
