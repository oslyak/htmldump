package htmldump_test

import (
	"regexp"
	"strings"
	"testing"
)

func extractHTMLTable(t *testing.T, doc string) string {
	t.Helper()

	being := strings.Index(doc, `<table`)
	end := strings.Index(doc, `</table>`) + len(`</table>`)
	table := doc[being:end]

	whiteSpaces := regexp.MustCompile(`>\s+<`)
	table = whiteSpaces.ReplaceAllString(table, "><")

	return table
}

func removeStyle(t *testing.T, table string) string {
	t.Helper()

	style := regexp.MustCompile(`class="[^"]+"|style="[^"]+"`)
	table = style.ReplaceAllString(table, ``)

	td := regexp.MustCompile(`<td\s+>`)
	table = td.ReplaceAllString(table, `<td>`)

	th := regexp.MustCompile(`<th\s+>`)
	table = th.ReplaceAllString(table, `<th>`)

	// tableRE := regexp.MustCompile(`<table\s+>`)
	// table = tableRE.ReplaceAllString(table, `<table>`)

	return table
}
