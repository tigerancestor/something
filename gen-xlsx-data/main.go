package main

import (
	"flag"
	"steady/tools/gen-xlsx-data/export"
	"steady/tools/gen-xlsx-data/printer"
)

var (
	paramJsonOut = flag.String("json_out", "test.json", "output json format (*.json)")
	paramGoOut   = flag.String("go_out", "", "output golang code (*.go)")
	paramTsOut   = flag.String("ts_out", "", "output typescript code (*.ts)")
)

func main() {
	flag.Parse()

	g := printer.NewGlobals()
	for _, v := range flag.Args() {
		g.InputFileList = append(g.InputFileList, v)
	}

	if *paramGoOut != "" {
		g.AddOutputType("golang", *paramGoOut)
		g.GoOut = *paramGoOut
	}
	if *paramJsonOut != "" {
		g.AddOutputType("json", *paramJsonOut)
		g.JsonOut = *paramJsonOut
	}
	if *paramTsOut != "" {
		g.AddOutputType("ts", *paramTsOut)
		g.TsOut = *paramTsOut
	}

	export.Run(g)

}
