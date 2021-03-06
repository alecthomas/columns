package columns

import (
	"bytes"
	"testing"

	"github.com/alecthomas/assert"
)

func TestFormatAlignRight(t *testing.T) {
	w := bytes.NewBuffer(nil)
	err := Format(w, 80, 2, []*Column{
		{Align: Right, Column: []interface{}{1, 20}},
		{Align: Right, Column: []interface{}{123, 2}},
	})
	assert.NoError(t, err)
	assert.Equal(t, " 1  123\n20    2\n", w.String())
}

func TestFormatAlignLeft(t *testing.T) {
	w := bytes.NewBuffer(nil)
	err := Format(w, 80, 2, []*Column{
		{Align: Left, Column: []interface{}{1, 20}},
		{Align: Left, Column: []interface{}{123, 2}},
	})
	assert.NoError(t, err)
	assert.Equal(t, "1   123\n20  2\n", w.String())
}

func TestFormatMinWidth(t *testing.T) {
	w := bytes.NewBuffer(nil)
	err := Format(w, 80, 2, []*Column{
		{MinWidth: 5, Column: []interface{}{1, 2}},
		{Column: []interface{}{123, 2}},
	})
	assert.NoError(t, err)
	assert.Equal(t, "1      123\n2      2\n", w.String())
}

func TestFormatWidth(t *testing.T) {
	w := bytes.NewBuffer(nil)
	err := Format(w, 4, 1, []*Column{
		{Column: []interface{}{1, 2}},
		{Column: []interface{}{123, 2}},
	})
	assert.NoError(t, err)
	assert.Equal(t, "1 12\n2 2\n", w.String())
}
