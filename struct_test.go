package htmldump_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"htmldump"

	"github.com/stretchr/testify/require"
)

func TestDumpStruct(t *testing.T) {
	t.Parallel()

	filter := newFilterAccounts()

	// dump.ToHTMLAndOpen(`/tmp/dump_struct_test.html`, filter, &filter)

	buffer := bytes.NewBuffer([]byte{})
	err := htmldump.ToHTML(buffer, filter)
	require.NoError(t, err)

	table := extractHTMLTable(t, buffer.String())
	table = removeStyle(t, table)

	require.Contains(t, table, `<caption>struct filterAccounts</caption>`)
	require.Contains(t, table, `<tr><th>Field</th><th>Type</th><th>Value</th></tr>`)

	expect := fmt.Sprintf("<tr><td>OwnerID</td><td>int64</td><td>%d</td></tr><tr>", filter.OwnerID)
	require.Contains(t, table, expect)

	expect = fmt.Sprintf("<tr><td>ExchangeID</td><td>int64</td><td>%d</td></tr><tr>", filter.ExchangeID)
	require.Contains(t, table, expect)

	expect = fmt.Sprintf(`<tr><td>EmptyPointer</td><td>*filter</td><td>%s</td></tr>`, htmldump.NULL)
	require.Contains(t, table, expect)

	strBuilder := strings.Builder{}
	strBuilder.WriteString(`<tr><td>order</td><td>order</td><td></td></tr>`)
	strBuilder.WriteString(`<tr><td>Column</td><td>string</td><td>`)
	strBuilder.WriteString(filter.filter.order.Column)
	strBuilder.WriteString("</td></tr>")

	require.Contains(t, table, strBuilder.String())
}
