/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: template
 * @Version: 1.0.0
 * @Date: 2023/3/13 17:37
 */

package sql2struct

import (
	"fmt"
	"github.com/CKzcb/goTour/internal/word"
	"html/template"
	"os"
)

const structTp1 = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{ $length := len .Comment }} {{ if gt $length 0 }}//
{{.Comment}} {{else}}// {{.Name}} {{end}}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}
	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTp1}
}

func (s *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf(" "+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (s *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl, _ := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}), nil).Parse(s.structTpl)

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
