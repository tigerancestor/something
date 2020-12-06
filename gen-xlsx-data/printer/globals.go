package printer

import (
	"steady/tools/gen-xlsx-data/model"
)

type Globals struct {
	InputFileList  []interface{}
	Printers       []*PrinterContext
	File2DataModel map[string][]*model.DataModel
	JsonOut        string
	TsOut          string
	GoOut          string
}

func NewGlobals() *Globals {
	p := &Globals{
		File2DataModel: make(map[string][]*model.DataModel),
	}

	return p
}

func (p *Globals) Print() bool {
	for _, printer := range p.Printers {
		//fmt.Printf("==========开始写入%s文件==========\n", p.name)
		if !printer.Start(p) {
			return false
		}
		//fmt.Printf("==========写入%s文件完成==========\n", p.name)
	}

	return true
}

func (p *Globals) AddOutputType(name string, outfile string) {

	if printer, ok := PrinterByExt[name]; ok {
		p.Printers = append(p.Printers, &PrinterContext{
			p:       printer,
			outFile: outfile,
			name:    name,
		})
	} else {
		panic("output type not found:" + name)
	}
}
