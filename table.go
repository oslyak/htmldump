package htmldump

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type styleT struct {
	paddingLeft  int
	paddingRight int
	background   string
	custom       string
}

func (style *styleT) toHTML() string {
	html := ` style="`
	empty := true

	if style.paddingLeft > 0 {
		html += `padding-left: ` + strconv.Itoa(style.paddingLeft) + `px;`
		empty = false
	}

	if style.paddingRight > 0 {
		html += `padding-right: ` + strconv.Itoa(style.paddingRight) + `px;`
		empty = false
	}

	if len(style.background) > 0 {
		html += `background: ` + style.background + `;`
		empty = false
	}

	if len(style.custom) > 0 {
		html += style.custom
		empty = false
	}

	html += `" `

	if empty {
		html = ``
	}

	return html
}

type rowT struct {
	cells []cellT
}

type cellT struct {
	colspan int
	key     bool
	value   string
	styleT
}

func (cell *cellT) toHTML(tag string) string {
	var result strings.Builder

	style := cell.styleT.toHTML()

	result.WriteString(`<` + tag)

	if cell.key {
		result.WriteString(` class="key"`)
	}

	if cell.colspan > 1 {
		result.WriteString(` colspan="` + strconv.Itoa(cell.colspan) + `"`)
	}

	result.WriteString(style)
	result.WriteString(`>`)
	result.WriteString(cell.value)
	result.WriteString(`</`)
	result.WriteString(tag)
	result.WriteString(`>`)
	result.WriteString("\n")

	return result.String()
}

type tableT struct {
	caption string
	header  headerT
	body    bodyT
	columns int
}

type (
	headerT []rowT
	bodyT   []rowT
)

func (row *rowT) addCell(cell cellT) {
	row.cells = append(row.cells, cell)
}

func (row *rowT) addCellStr(value string, styles ...styleT) *rowT {
	style := styleT{}
	if len(styles) > 0 {
		style = styles[0]
	}

	row.cells = append(row.cells, cellT{value: value, styleT: style})

	return row
}

func (header *headerT) toHTML() string {
	var html string

	for _, row := range *header {
		html += "      <tr>\n"
		for _, cell := range row.cells {
			html += `        ` + cell.toHTML(`th`)
		}

		html += "      </tr>\n"
	}

	return strings.TrimSuffix(html, "\n")
}

func (body *bodyT) toHTML() string {
	var html string

	for _, row := range *body {
		html += "      <tr>\n"
		for _, cell := range row.cells {
			html += `        ` + cell.toHTML(`td`)
		}

		html += "      </tr>\n"
	}

	return strings.TrimSuffix(html, "\n")
}

func (table *tableT) addBodyRow(row rowT) {
	table.body = append(table.body, row)
}

func (table *tableT) addHeaderRow(row rowT) {
	table.header = append(table.header, row)
}

func (table *tableT) Caption(caption string) *tableT {
	table.caption = caption
	return table
}

func (table *tableT) toHTML(doc *htmlDocument) {
	doc.add(`  <table class="styled-table">`)

	if len(table.caption) > 0 {
		doc.add(`    <caption>` + table.caption + `</caption>`)
	}

	doc.add(`    <thead>`).
		add(table.header.toHTML()).
		add(`    </thead>`)
	doc.add(`    <tbody>`).
		add(table.body.toHTML()).
		add(`    </tbody>`)
	doc.add(`  </table>`)
}

func (table *tableT) headerRow(valueType reflect.Type, captions, types *rowT) {
	if valueType.Kind() == reflect.Pointer {
		valueType = valueType.Elem()
	}

	if valueType.Kind() == reflect.Struct {
		structType := valueType
		for idx := 0; idx < structType.NumField(); idx++ {
			field := structType.Field(idx)

			if isStructOrPointerToStruct(field.Type) && !isSkippedType(field.Type) {
				captions.addCell(cellT{
					value:   structFieldCaption(field),
					colspan: structNumbeOfFields(field.Type),
				})

				embeddedStructHeader(field.Type, types)

				continue
			}

			captions.addCellStr(field.Name)

			fieldType := field.Type
			if field.Type.Kind() == reflect.Pointer {
				fieldType = field.Type.Elem()
				types.addCellStr(`*` + fieldType.Name())
			} else {
				types.addCellStr(fieldType.Name())
			}
		}
	} else {
		captions.addCellStr(`value`)
		types.addCellStr(valueType.Name())
	}

	table.addHeaderRow(*captions)
	table.addHeaderRow(*types)

	table.columns = len(types.cells)
}

func embeddedStructHeader(structType reflect.Type, types *rowT) {
	if structType.Kind() == reflect.Pointer {
		structType = structType.Elem()
	}

	for idx := 0; idx < structType.NumField(); idx++ {
		structField := structType.Field(idx)
		fieldType := structField.Type
		fieldTypeName := fieldType.Name()

		if fieldType.Kind() == reflect.Pointer {
			fieldType = fieldType.Elem()
			fieldTypeName = `*` + fieldType.Name()
		}

		if structField.Name == fieldTypeName {
			types.addCellStr(fieldTypeName)

			continue
		}

		types.addCellStr(fmt.Sprintf(`%s(%s)`, structField.Name, fieldTypeName))
	}
}

func structNumbeOfFields(structType reflect.Type) int {
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	return structType.NumField()
}

func structFieldCaption(field reflect.StructField) string {
	fieldTypeName := field.Type.Name()

	if field.Type.Kind() == reflect.Pointer {
		field.Type = field.Type.Elem()
		fieldTypeName = `*` + field.Type.Name()
	}

	if field.Name == fieldTypeName {
		return field.Name
	}

	return fmt.Sprintf(`%s(%s)`, field.Name, fieldTypeName)
}
