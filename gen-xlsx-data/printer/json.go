package printer

import (
	"errors"
	"fmt"
	"steady/tools/gen-xlsx-data/model"
	"strconv"
	"strings"
)

type jsonPrinter struct {
}

func (*jsonPrinter) Run(g *Globals) bool {
	for _, dataModels := range g.File2DataModel {
		if !genTableJson(g, dataModels) {
			return false
		}

	}

	return true
}

func genTableJson(g *Globals, dataModels []*model.DataModel) bool {
	res := true
	for _, data := range dataModels {
		res = res && genSheetJson(g, data)
	}

	return res
}

func translateValue(fieldName string, fieldType string, rawValue string) interface{} {
	if fieldType == "string" {
		return StringEscape(rawValue)
	} else if fieldType == "dateTime" {
		return StringEscape(rawValue)
	} else if fieldType == "giveItems" || fieldType == "giveMoneys" {
		return StringEscape(rawValue)
	}

	if rawValue == "" {
		return 0
	}

	if value, err := strconv.ParseFloat(rawValue, 64); err != nil {
		panic(errors.New(fmt.Sprintf("fieldName:%s,filedType:%s, rawValue:%s, err:%s", fieldName, fieldType, rawValue, err.Error())))
	} else {
		return int64(value)
	}
}

func genSheetJson(g *Globals, dataModel *model.DataModel) bool {
	fmt.Println(fmt.Sprintf("开始生成：%s/%s.json", dataModel.Package, dataModel.Name))
	bf := NewStream()
	bf.Printf("{\n")

	existId := false
	for _, v := range dataModel.HeaderFields {
		if dataModel.TranslateFirstUpper(v.FieldName) == "Id" {
			existId = true
			break
		}
	}

	// 遍历每一行
	id := 0
	noRepeatedIdMap := map[string]struct{}{}
	for i, r := range dataModel.Lines {
		id++

		// 遍历每一列
		temp := ""
		putId := false
		for j, value := range r.Values {
			if existId {
				idValue := translateValue(value.FieldName, value.FieldType, value.RawValue)
				if strings.ToLower(value.FieldName) == "id" {
					if _, exist := noRepeatedIdMap[fmt.Sprintf("%v", idValue)]; exist {
						panic(fmt.Sprintf("%s/%s.json", dataModel.Package, dataModel.Name) + " repeated id id:" + fmt.Sprintf("%v", idValue))
					}
					noRepeatedIdMap[fmt.Sprintf("%v", idValue)] = struct{}{}
					if value.FieldType == "string" || value.FieldType == "dateTime" {
						bf.Printf("     %s:{ \n", idValue)
					} else {
						bf.Printf("     \"%d\":{ \n", idValue)
					}
				}
			} else {
				if !putId {
					bf.Printf("     \"%d\":{ \n", id)
					putId = true
				}
			}

			if value.RawValue != "" {
				if j == len(r.Values)-1 {
					temp += fmt.Sprintf("         \"%s\": %v \n", dataModel.TranslateFirstUpper(value.FieldName), translateValue(value.FieldName, value.FieldType, value.RawValue))
				} else {
					temp += fmt.Sprintf("         \"%s\": %v, \n", dataModel.TranslateFirstUpper(value.FieldName), translateValue(value.FieldName, value.FieldType, value.RawValue))
				}
			}
		}
		if strings.HasSuffix(temp, ", \n") {
			bf.Printf(temp[:len(temp)-len(", \n")])
			bf.Printf("\n")
		} else {
			bf.Printf(temp)
		}

		if i == len(dataModel.Lines)-1 {
			bf.Printf("     }\n")
		} else {
			bf.Printf("     },\n")
		}
	}

	bf.Printf("}")
	bf.WriteFile(g.JsonOut + "/" + dataModel.Package + "/" + dataModel.Name + ".json")
	fmt.Println(fmt.Sprintf("生成：%s/%s.json成功", dataModel.Package, dataModel.Name))

	return true
}

func StringEscape(s string) string {
	b := make([]byte, 0)
	var index int

	// 表中直接使用换行会干扰最终合并文件格式, 所以转成\n,由pbt文本解析层转回去
	for index < len(s) {
		c := s[index]

		switch c {
		case '"':
			b = append(b, '\\')
			b = append(b, '"')
		case '\n':
			b = append(b, '\\')
			b = append(b, 'n')
		case '\r':
			b = append(b, '\\')
			b = append(b, 'r')
		case '%':
			b = append(b, '%')
			b = append(b, '%')
		case '\\':

			var nextChar byte
			if index+1 < len(s) {
				nextChar = s[index+1]
			}

			b = append(b, '\\')

			switch nextChar {
			case 'n', 'r':
			default:
				b = append(b, c)
			}

		default:
			b = append(b, c)
		}

		index++

	}

	return fmt.Sprintf("\"%s\"", string(b))
}

func init() {
	RegisterPrinter("json", &jsonPrinter{})
}
