package htmldump_test

import (
	"bytes"

	"strings"
	"testing"

	"github.com/oslyak/htmldump"

	"github.com/stretchr/testify/require"
)

func TestDumpMap(t *testing.T) {
	mmap := map[string]string{
		`a`: `a1`,
		`b`: `b2`,
		`c`: `c3`,
	}

	filter := newFilterAccounts()

	filters := map[int64]*filterAccounts{
		0: nil,
		1: &filter,
		2: nil,
		3: &filter,
		4: &filter,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := htmldump.ToHTML(buffer, mmap)
	require.NoError(t, err)

	table := extractHTMLTable(t, buffer.String())
	table = removeStyle(t, table)

	expect := `<caption>map[string]string (length: 3)</caption>`
	require.Contains(t, table, expect)

	// dump.ToHTMLAndOpen(`/tmp/dump_map_test.html`, filters)

	buffer = bytes.NewBuffer([]byte{})
	err = htmldump.ToHTML(buffer, filters)
	require.NoError(t, err)

	table = extractHTMLTable(t, buffer.String())
	table = removeStyle(t, table)

	expect = `<caption>map[int64]*dump_test.filterAccounts (length: 5)</caption>`
	require.Contains(t, table, expect)

	trBuilder := strings.Builder{}
	trBuilder.WriteString(`<tr>`)
	trBuilder.WriteString(`<th>int64</th>`)
	trBuilder.WriteString(`<th>string</th>`)
	trBuilder.WriteString(`<th>int64</th>`)
	trBuilder.WriteString(`<th>int64</th>`)
	trBuilder.WriteString(`<th>Time</th>`)
	trBuilder.WriteString(`<th>*Time</th>`)
	trBuilder.WriteString(`<th>*Time</th>`)
	trBuilder.WriteString(`<th>IncludeRemoved(bool)</th>`)
	trBuilder.WriteString(`<th>Limit(uint64)</th>`)
	trBuilder.WriteString(`<th>Offset(uint64)</th>`)
	trBuilder.WriteString(`<th>orderPtr(*order)</th>`)
	trBuilder.WriteString(`<th>order</th>`)
	trBuilder.WriteString(`<th>IncludeRemoved(bool)</th>`)
	trBuilder.WriteString(`<th>Limit(uint64)</th>`)
	trBuilder.WriteString(`<th>Offset(uint64)</th>`)
	trBuilder.WriteString(`<th>orderPtr(*order)</th>`)
	trBuilder.WriteString(`<th>order</th>`)
	trBuilder.WriteString(`<th>IncludeRemoved(bool)</th>`)
	trBuilder.WriteString(`<th>Limit(uint64)</th>`)
	trBuilder.WriteString(`<th>Offset(uint64)</th>`)
	trBuilder.WriteString(`<th>orderPtr(*order)</th>`)
	trBuilder.WriteString(`<th>order</th>`)
	trBuilder.WriteString(`</tr>`)

	expect = trBuilder.String()
	require.Contains(t, table, expect)
}
