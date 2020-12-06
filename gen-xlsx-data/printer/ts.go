package printer

import (
	"fmt"
	"steady/tools/gen-xlsx-data/model"
	"text/template"
)

const tsTemplate = `// DO NOT EDIT!!

// {{.SheetName}}
class Config{{$.StrFirstToUpper .Name | $.TranslateFirstUpper}}  {
{{range .HeaderFields}}
    /**{{.Comment}}*/
	public {{$.TranslateFirstUpper .FieldName}} : {{$.TranslateTsFieldType .FieldType}};      
{{end}}
}
`

type tsPrinter struct {
}

func (*tsPrinter) Run(g *Globals) bool {
	for _, dataModels := range g.File2DataModel {
		if !genTableTs(g, dataModels) {
			return false
		}

	}

	return true
}

func genTableTs(g *Globals, dataModels []*model.DataModel) bool {
	res := true
	for _, data := range dataModels {
		res = res && genSheetTs(g, data)
	}

	return res
}

func genSheetTs(g *Globals, dataModel *model.DataModel) bool {
	tpl, err := template.New("ts").Funcs(template.FuncMap{
		"TranslateFirstLower":  dataModel.TranslateFirstLower,
		"TranslateFirstUpper":  dataModel.TranslateFirstUpper,
		"TranslateTsFieldType": dataModel.TranslateTsFieldType,
		"StrFirstToUpper":      dataModel.StrFirstToUpper,
	}).Parse(tsTemplate)
	if err != nil {
		fmt.Println(err)
		return false
	}

	bf := NewStream()

	err = tpl.Execute(bf.Buffer(), dataModel)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = bf.WriteFile(g.TsOut + "/" + dataModel.Package + "/" + dataModel.Name + ".ts")

	return err == nil
}

func init() {
	RegisterPrinter("ts", &tsPrinter{})
}
