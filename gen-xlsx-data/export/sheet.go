package export

import (
	"github.com/tealeg/xlsx"
	"strings"
)

// 描述一个表单
type Sheet struct {
	*xlsx.Sheet

	Row int // 当前行

	Column int // 当前列

	file *File // 指向父级

}

// 取行列信息
func (p *Sheet) GetRC() (int, int) {

	return p.Row + 1, p.Column + 1

}

// 获取单元格 cursor=行,  index=列
func (p *Sheet) GetCellData(cursor, index int) string {

	if cursor >= len(p.Rows) {
		return ""
	}

	r := p.Rows[cursor]
	for len(r.Cells) <= index {
		r.AddCell()
	}

	return strings.TrimSpace(r.Cells[index].Value)
}

func (p *Sheet) GetCellDataAsNumeric(cursor, index int) string {

	if cursor >= len(p.Rows) {
		return ""
	}

	r := p.Rows[cursor]
	for len(r.Cells) <= index {
		r.AddCell()
	}

	gn, err := r.Cells[index].FormattedValue()

	if err != nil {
		return ""
	}

	return gn
}

// 设置单元格
func (p *Sheet) SetCellData(cursor, index int, data string) {

	p.Cell(cursor, index).Value = data
}

// 整行都是空的
func (p *Sheet) IsFullRowEmpty(row, maxCol int) bool {

	for col := 0; col <= maxCol; col++ {

		data := p.GetCellData(row, col)

		if data != "" {
			return false
		}
	}

	return true
}

func NewSheet(file *File, sheet *xlsx.Sheet) *Sheet {
	p := &Sheet{
		file:  file,
		Sheet: sheet,
	}

	return p
}
