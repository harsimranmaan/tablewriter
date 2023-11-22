package main

import (
	"os"
	"tablewriter/tablewriter"
)

func main() {
	tr := tablewriter.New(os.Stdout, [][]string{
		{"test11", "test1"},
		{"test22", "test12"},
		{"test33", "test13"},
		{"test44444444", "test14"},
	})
	tr.Display()
}
