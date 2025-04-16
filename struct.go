package htmldump

import (
	"fmt"
	"reflect"
)

func structToHTML(doc *htmlDocument, input interface{}) error {
	structType := reflect.TypeOf(input)
	if !isStructOrPointerToStruct(structType) {
		return fmt.Errorf(`[structToHTML] only accepts struct or pointer to struct, got %s`, structType.Kind())
	}

	var caption string

	if structType.Kind() == reflect.Pointer {
		caption = `*struct ` + structType.Elem().Name()
	} else {
		caption = `struct ` + structType.Name()
	}

	table := new(tableT).
		Caption(caption).
		structHeader().
		structBody(input, 0)

	table.toHTML(doc)

	return nil
}

func (table *tableT) structHeader() *tableT {
	row := new(rowT)
	row.cells = append(row.cells, cellT{value: `Field`})
	row.cells = append(row.cells, cellT{value: `Type`})
	row.cells = append(row.cells, cellT{value: `Value`})

	table.addHeaderRow(*row)

	return table
}

func (table *tableT) structBody(input interface{}, level int) *tableT {
	const spacer = 12

	pointer := convertToPointer(input)
	structType := reflect.TypeOf(pointer).Elem()
	structValue := reflect.ValueOf(pointer).Elem()

	for idx := 0; idx < structType.NumField(); idx++ {
		row := new(rowT)

		structField := structType.Field(idx)
		fieldName := structField.Name
		fieldTypeName := getFieldTypeName(structField.Type)
		fieldValue := getUnexportedField(structValue.Field(idx))

		nameStyle := styleT{paddingLeft: spacer * level}
		style := styleT{}

		if isStructOrPointerToStruct(fieldValue.Type()) && !isSkippedType(fieldValue.Type()) {
			style.background = getBackground(level)
			nameStyle.background = style.background

			row.addCellStr(fieldName, nameStyle).
				addCellStr(fieldTypeName, style).
				addCellStr(structFieldValue(structField, fieldValue), style)

			table.addBodyRow(*row)

			if fieldValue.IsValid() && !fieldValue.IsZero() {
				table.structBody(fieldValue.Interface(), level+1)
			}

			continue
		}

		style.background = getBackground(level)
		nameStyle.background = style.background
		row.addCellStr(fieldName, nameStyle).
			addCellStr(fieldTypeName, style).
			addCellStr(formatValue(fieldValue), style)

		table.addBodyRow(*row)
	}

	return table
}

func structFieldValue(structField reflect.StructField, fieldValue reflect.Value) string {
	switch {
	case structField.Type.Kind() == reflect.Pointer && fieldValue.IsNil():
		return NULL
	case structField.Type.Kind() != reflect.Pointer && fieldValue.IsZero():
		return `{}`
	default:
		return ``
	}
}

func getBackground(level int) string {
	gradient := []string{``, `#AFFFFF`, `#6DEFFF`, `#47D3FF`, `#00B7EB`}
	index := level

	if index > len(gradient)-1 {
		index = len(gradient) - 1
	}

	return gradient[index]
}
