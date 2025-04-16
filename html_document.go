package htmldump

import (
	"fmt"
	"io"
)

type htmlDocument struct {
	writer io.Writer
	body   string
}

func (doc *htmlDocument) save() (int, error) {
	num, err := io.WriteString(doc.writer, doc.body)
	if err != nil {
		werr := fmt.Errorf("[(doc *htmlDocument) save()] writing to io.Writer error: %w", err)
		return 0, werr
	}

	return num, nil
}

// Add HTML table row to the document body.
func (doc *htmlDocument) add(str string) *htmlDocument {
	doc.body += str + "\n"

	return doc
}

func newHTMLDocument(writer io.Writer) *htmlDocument {
	doc := &htmlDocument{
		writer: writer,
		body:   ``,
	}

	doc.body += `<!DOCTYPE html>
<html>
<head> 
  <style>
    .styled-table {
        border-collapse: collapse;
        margin: 25px 0;
        font-size: 0.9em;
        font-family: sans-serif;
        min-width: 400px;
        box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
        border: 2px solid rgb(150, 150, 150);
        border-top: none;
    }

    .styled-table caption {
        text-align: left;
        font-size: 1.5em;
        font-weight: bold;
        padding: 7px;
        background-color: rgb(220, 220, 220);
        border-top: 2px solid rgb(150, 150, 150);
        border-left: 2px solid rgb(150, 150, 150);
        border-right: 2px solid rgb(150, 150, 150);
        border-bottom: none;
    }
        
    .styled-table thead tr {
        white-space: nowrap;
        background-color: #009879;
        color: #ffffff;
        text-align: center;
    }

    .styled-table thead tr th {
        border: 1px solid #006e58;
    }

    .styled-table thead tr:first-of-type {
        border-top: 2px solid #009879;
    }		

    .styled-table thead tr:last-of-type {
        border-bottom: 2px solid #009879;
    }	

    .styled-table tbody tr {
        border-bottom: 1px solid #dddddd;
    }

    .styled-table tbody tr td.key{
        color: #006650;
        font-weight: bold;
    }        

    .styled-table tbody tr td{
        border-right: 1px solid #dddddd;
    }        

    .styled-table tbody tr:nth-child(even) {background: rgb(250, 250, 250)}
    .styled-table tbody tr:nth-child(odd) {background: rgb(230, 230, 230)}

    .styled-table th,
    .styled-table td {
        padding: 4px 7px;
    }
        
    .styled-table tbody tr:hover {
        color: #006650;
    }
  </style>    
</head>

<body>
`

	return doc
}
