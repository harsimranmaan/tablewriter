package tablewriter

import (
	"errors"
	"io"
)

type Table struct {
	w     io.Writer
	table [][]string
}

var errColsNotMatching = errors.New("columns don't match")

// New creates an instance of tablewriter
func New(w io.Writer, table [][]string) *Table {
	return &Table{w: w, table: table}
}

// Display writes the table to the writer
func (t *Table) Display() (err error) {
	_, err = t.w.Write([]byte("test"))
	return err
}
