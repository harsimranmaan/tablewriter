package tablewriter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDisplay(t *testing.T) {
	tests := []struct {
		input       [][]string
		expectedErr error
		expectedOut string
	}{
		{
			[][]string{
				{"test", "test1"},
				{"test2", "test3"},
			}, nil,
			`
| test  | test1 |
| test2 | test3 |
`,
		},
		{
			[][]string{
				{"this is a test", "a"},
				{"test2", "test3"},
			}, nil,
			`
| this is a test | a     |
| test2          | test3 |
`,
		},

		{
			[][]string{
				{"this is a test"},
				{"test2", "test3"},
			}, errColsNotMatching,
			``,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var b bytes.Buffer
			tr := New(&b, test.input)
			err := tr.Display()
			if test.expectedErr != err {
				t.Errorf("%s != %s", test.expectedErr, err)
			}
			if test.expectedOut != b.String() {
				t.Errorf("%s != %s", test.expectedOut, b.String())
			}

		})
	}
}
