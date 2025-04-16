package htmldump

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"
	"unsafe"
)

const NULL = `NULL`

// ToHTML dumps the specified inputs to an HTML file at the specified path.
// The inputs can be structs, slices, maps, or pointers to them.
func ToHTML(writer io.Writer, inputs ...interface{}) error {
	if len(inputs) == 0 {
		return errors.New(`[ToHTML] requires at least one inputs argument`)
	}

	var (
		err error
		doc = newHTMLDocument(writer)
	)

	for _, input := range inputs {
		reflectedValue := reflect.ValueOf(input)

		switch {
		case isMapOrPointerToMap(reflectedValue):
			err = mapToHTML(doc, reflectedValue)
		case isPointerToSliceOrSlice(reflectedValue):
			err = sliceToHTML(doc, reflectedValue)
		case isStructOrPointerToStruct(reflectedValue.Type()):
			err = structToHTML(doc, input)
		default:
			err = errors.New(`[ToHTML] only accepts: structs, slices, maps, and pointers to them`)
		}

		if err != nil {
			return err
		}
	}

	doc.add("</body>\n</html>")
	_, err = doc.save()

	return err
}

// ToHTMLAndOpen is a convenience function that calls ToHTML and then opens the
// generated HTML file in the default browser.
func ToHTMLAndOpen(path string, inputs ...interface{}) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		werr := fmt.Errorf("[ToHTMLAndOpen] getting absolute path to %s error: %w", path, err)
		panic(werr)
	}

	file, err := os.Create(absolutePath)
	if err != nil {
		werr := fmt.Errorf("[ToHTMLAndOpen] file %s creating error: %w", absolutePath, err)
		panic(werr)
	}

	defer file.Close()

	err = ToHTML(file, inputs...)
	if err != nil {
		panic(err)
	}

	openHTML(path)
}

// openHTML opens the specified HTML file in the default browser.
func openHTML(path string) {
	err := exec.Command(`xdg-open`, path).Start()
	if err != nil {
		panic(err)
	}
}

// Main struct always had to be a pointer.
func getUnexportedField(field reflect.Value) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
}

func convertToPointer(input interface{}) interface{} {
	structType := reflect.TypeOf(input)

	var pointer reflect.Value

	if structType.Kind() == reflect.Pointer {
		return input
	} else {
		inputValue := reflect.ValueOf(input)
		pointer = reflect.New(inputValue.Type())
		pointer.Elem().Set(inputValue)

		return pointer.Elem().Addr().Interface()
	}
}

func getFieldTypeName(fieldType reflect.Type) string {
	if fieldType.Kind() == reflect.Pointer {
		return `*` + fieldType.Elem().Name()
	} else {
		return fieldType.Name()
	}
}

func formatValue(value reflect.Value, format ...string) string {
	valueType := value.Type()
	if valueType.Kind() == reflect.Pointer {
		valueType = valueType.Elem()
		value = value.Elem()

		if !value.IsValid() {
			return ``
		}
	}

	switch valueType.Name() {
	case `Time`:
		t, ok := value.Interface().(time.Time)
		if ok && !t.IsZero() {
			return t.Format(`02.01.2006 15:04:05`)
		}
	case `NullTime`:
		nullTime, ok := value.Interface().(sql.NullTime)
		if !ok || !nullTime.Valid {
			return NULL
		} else {
			return nullTime.Time.Format(`02.01.2006 15:04:05`)
		}
	default:
		return fmt.Sprintf(formatString(format...), value)
	}

	return ``
}

func formatString(format ...string) string {
	if len(format) > 0 {
		return format[0]
	}

	return `%v`
}
