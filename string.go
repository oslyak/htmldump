package htmldump

import (
	"fmt"
	"reflect"
)

func isStringOrPointerToString(fieldType reflect.Type) bool {
	if fieldType.Kind() == reflect.Pointer {
		return fieldType.Elem().Kind() == reflect.String
	}

	return fieldType.Kind() == reflect.String
}

// Generate HTML table for a string.
func stringToHTML(doc *htmlDocument, reflectedString reflect.Value) error {
	if reflectedString.Kind() != reflect.String {
		return fmt.Errorf(`[stringToHTML] only accepts string, got %s`, reflectedString.Kind())
	}

	table := new(tableT)
	table.Caption(fmt.Sprintf("String (length: %d) ", len(reflectedString.String()))).
		stringBody(reflectedString).
		toHTML(doc)

	return nil
}

func (table *tableT) stringBody(reflectedString reflect.Value) *tableT {
	str := reflectedString.String()
	var row rowT

	row.addCell(cellT{value: "Value", key: true})
	row.addCell(cellT{value: str})
	table.addBodyRow(row)

	return table
}
