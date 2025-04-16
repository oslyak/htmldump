package htmldump_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/oslyak/htmldump"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

type order struct {
	Column string
}

type filter struct {
	IncludeRemoved bool
	Limit          uint64
	Offset         uint64
	orderPtr       *order
	order
}

type filterAccounts struct {
	Name         string
	OwnerID      int64
	ExchangeID   int64
	today        time.Time
	yesterdayPtr *time.Time
	EmptyDatePtr *time.Time
	EmptyPointer *filter
	EmptyFilter  filter

	filter
}

func newFilterAccounts() filterAccounts {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	return filterAccounts{
		Name:         gofakeit.Word(),
		OwnerID:      int64(gofakeit.Uint8()),
		ExchangeID:   int64(gofakeit.Uint8()),
		today:        now,
		yesterdayPtr: &yesterday,
		filter: filter{
			IncludeRemoved: gofakeit.Bool(),
			Limit:          uint64(gofakeit.Uint16()),
			Offset:         uint64(gofakeit.Uint16()),
			order: order{
				Column: gofakeit.Word(),
			},
			orderPtr: &order{Column: `pointer`},
		},
	}
}

func TestDumpSlice(t *testing.T) {
	t.Parallel()
	slice := []string{`a`, `b`, `c`}

	buffer := bytes.NewBuffer([]byte{})
	err := htmldump.ToHTML(buffer, slice)
	require.NoError(t, err)

	table := extractHTMLTable(t, buffer.String())
	table = removeStyle(t, table)

	expect := `<caption>[]string (length: 3)</caption>`
	require.Contains(t, table, expect)

	expect = `<tr><th>index</th><th>value</th></tr><tr><th>int</th><th>string</th></tr>`
	require.Contains(t, table, expect)

	expect = `<tr><td>0</td><td>a</td></tr><tr><td>1</td><td>b</td></tr><tr><td>2</td><td>c</td></tr>`
	require.Contains(t, table, expect)

	filter := newFilterAccounts()

	slice1 := []*filterAccounts{nil, &filter, &filter, &filter}

	// dump.ToHTMLAndOpen(`/tmp/dump_slice_test.html`, slice1)

	buffer = bytes.NewBuffer([]byte{})
	err = htmldump.ToHTML(buffer, slice1)
	require.NoError(t, err)

	table = extractHTMLTable(t, buffer.String())
	table = removeStyle(t, table)

	expect = `<caption>[]*dump_test.filterAccounts`
	require.Contains(t, table, expect)

	expect = `<caption>[]*dump_test.filterAccounts`
	require.Contains(t, table, expect)

	strBuilder := strings.Builder{}
	strBuilder.WriteString(`<tr><th>index</th><th>Name</th><th>OwnerID</th><th>ExchangeID</th>`)
	strBuilder.WriteString(`<th>today</th><th>yesterdayPtr</th>`)
	strBuilder.WriteString(`<th>EmptyDatePtr</th><th colspan="5">EmptyPointer(*filter)</th>`)
	strBuilder.WriteString(`<th colspan="5">EmptyFilter(filter)</th><th colspan="5">filter</th></tr>`)

	require.Contains(t, table, strBuilder.String())
}
