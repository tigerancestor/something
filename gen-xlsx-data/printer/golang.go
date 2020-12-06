package printer

import (
	"fmt"
	"os"
	"os/exec"
	"steady/tools/gen-xlsx-data/model"
	"text/template"
)

const goTemplate = `// DO NOT EDIT!!
package cfg{{$.StrStripUnderline .Package}}

{{.Imports}}

// {{.SheetName}}
type {{$.StrFirstToUpper .Name | $.TranslateFirstUpper}}  struct {
{{range .HeaderFields}}{{if eq .FieldType "dateTime" }}
    {{$.TranslateFirstUpper .FieldName}}   string     //{{.Comment}}
    {{$.TranslateFirstUpper .FieldName}}Str2Time   {{$.TranslateGolangFieldType .FieldType}}     //{{.Comment}}字符串
{{else if eq .FieldType "giveItems" }}
    {{$.TranslateFirstUpper .FieldName}}   string     //{{.Comment}}
    {{$.TranslateFirstUpper .FieldName}}Str2GiveItems   {{$.TranslateGolangFieldType .FieldType}}     //{{.Comment}}字符串
{{else if eq .FieldType "giveMoneys" }}
    {{$.TranslateFirstUpper .FieldName}}   string     //{{.Comment}}
    {{$.TranslateFirstUpper .FieldName}}Str2GiveMoneys  {{$.TranslateGolangFieldType .FieldType}}     //{{.Comment}}字符串
{{else}}
{{$.TranslateFirstUpper .FieldName}}   {{$.TranslateGolangFieldType .FieldType}}     //{{.Comment}}{{end}}{{end}}
}
`

type goPrinter struct {
}

func (*goPrinter) Run(g *Globals) bool {
	for _, dataModels := range g.File2DataModel {
		if !genTableGo(g, dataModels) {
			return false
		}
	}

	return true
}

func genTableGo(g *Globals, dataModels []*model.DataModel) bool {
	res := true
	for _, data := range dataModels {
		res = res && genSheetGo(g, data)
	}

	return res
}

func genSheetGo(g *Globals, dataModel *model.DataModel) bool {
	tpl, err := template.New("golang").Funcs(template.FuncMap{
		"TranslateFirstLower":      dataModel.TranslateFirstLower,
		"TranslateFirstUpper":      dataModel.TranslateFirstUpper,
		"TranslateGolangFieldType": dataModel.TranslateGolangFieldType,
		"StrStripUnderline":        dataModel.StrStripUnderline,
	}).Parse(goTemplate)
	if err != nil {
		fmt.Println(err)
		return false
	}

	bf := NewStream()
	importTime := false
	importMsg := false
	for _, t := range dataModel.HeaderFields {
		if !importTime && t.FieldType == "dateTime" {
			dataModel.Imports += `import "time"
`
			importTime = true
		} else if !importMsg && (t.FieldType == "giveItems" || t.FieldType == "giveMoneys") {
			dataModel.Imports += `import "game-stage/msg"
`
			importMsg = true
		}
	}
	err = tpl.Execute(bf.Buffer(), dataModel)
	if err != nil {
		fmt.Println(err)
		return false
	}

	filePath := g.GoOut + "/" + dataModel.Package + "/" + dataModel.Name + ".go"
	err = bf.WriteFile(filePath)

	// 进行格式化
	cmd := exec.Command("gofmt", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	return err == nil
}

func init() {
	RegisterPrinter("golang", &goPrinter{})
}
