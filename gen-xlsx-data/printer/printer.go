package printer

type PrinterContext struct {
	outFile string
	p       Printer
	name    string
}

func (p *PrinterContext) Start(g *Globals) bool {

	return p.p.Run(g)
}

type Printer interface {
	Run(g *Globals) bool
}

var PrinterByExt = make(map[string]Printer)

func RegisterPrinter(ext string, p Printer) {

	if _, ok := PrinterByExt[ext]; ok {
		panic("duplicate printer")
	}

	PrinterByExt[ext] = p
}
