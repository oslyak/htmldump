package htmldump

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Generate HTML table for a slice, with type and values.
func sliceToHTML(doc *htmlDocument, reflectedSlice reflect.Value) error {
	if !isPointerToSliceOrSlice(reflectedSlice) {
		return fmt.Errorf(`[sliceToHTML] only accepts slice or pointer to it, got %s`, reflectedSlice.Kind())
	}

	caption, err := sliceCaption(reflectedSlice)
	if err != nil {
		return err
	}

	table := new(tableT)
	table.Caption(caption).
		sliceHeader(reflectedSlice).
		sliceBody(reflectedSlice).
		toHTML(doc)

	return nil
}

// Returns the caption of HTML table for a slice values type.
func sliceCaption(reflectedSlice reflect.Value) (string, error) {
	typeName, err := SliceTypeName(reflectedSlice.Type())
	if err != nil {
		return ``, err
	}

	return fmt.Sprintf(`%s (length: %d)`, typeName, reflectedSlice.Len()), nil
}

// Generate HTML table header for a slice with slice key and slice value types and struct fields.
func (table *tableT) sliceHeader(reflectedSlice reflect.Value) *tableT {
	var captions, types rowT

	captions.addCell(cellT{value: `index`, key: true})
	types.addCell(cellT{value: `int`, key: true})

	table.headerRow(reflectedSlice.Type().Elem(), &captions, &types)

	return table
}

func (table *tableT) sliceBody(reflectedSlice reflect.Value) *tableT {
	for index := 0; index < reflectedSlice.Len(); index++ {
		var row rowT

		row.addCell(cellT{value: strconv.Itoa(index), key: true})

		item := reflectedSlice.Index(index)
		if item.Kind() == reflect.Pointer {
			item = item.Elem()
		}

		switch {
		case !item.IsValid():
			row.addCell(cellT{value: NULL, colspan: table.columns - 1})
		case item.Kind() == reflect.Struct:
			structRow(&row, item)
		default:
			row.addCellStr(formatValue(item))
		}

		table.addBodyRow(row)
	}

	return table
}

func structRow(row *rowT, item reflect.Value) {
	for idx := 0; idx < item.NumField(); idx++ {
		field := item.Field(idx)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()

		if isStructOrPointerToStruct(field.Type()) && !isSkippedType(field.Type()) {
			if field.Kind() == reflect.Pointer {
				field = field.Elem()
			}

			if field.IsValid() {
				embeddedStructBody(row, field)
			} else {
				itemType := item.Type().Field(idx).Type
				row.addCell(cellT{value: NULL, colspan: structNumbeOfFields(itemType)})
			}

			continue
		}

		row.addCellStr(formatValue(field))
	}
}

func embeddedStructBody(row *rowT, embeddedStruct reflect.Value) {
	for idx := 0; idx < embeddedStruct.NumField(); idx++ {
		field := embeddedStruct.Field(idx)
		structField := embeddedStruct.Type().Field(idx)

		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				row.addCellStr(NULL)
				continue
			}

			field = field.Elem()
		}

		if isStructOrPointerToStruct(structField.Type) && !isSkippedType(structField.Type) {
			if field.IsValid() {
				row.addCellStr(formatValue(field, `%+v`))
			} else {
				row.addCellStr(NULL)
			}

			continue
		}

		row.addCellStr(formatValue(field))
	}
}

func isStructOrPointerToStruct(fieldType reflect.Type) bool {
	if fieldType.Kind() == reflect.Pointer {
		return fieldType.Elem().Kind() == reflect.Struct
	}

	return fieldType.Kind() == reflect.Struct
}

func isSkippedType(fieldType reflect.Type) bool {
	const typesToSkipp = `Time NullTime`

	if fieldType.Kind() == reflect.Pointer {
		fieldType = fieldType.Elem()
	}

	return strings.Contains(typesToSkipp, fieldType.Name())
}

func isPointerToSliceOrSlice(fieldType reflect.Value) bool {
	if fieldType.Kind() == reflect.Pointer {
		return fieldType.Elem().Kind() == reflect.Slice
	}

	return fieldType.Kind() == reflect.Slice
}
