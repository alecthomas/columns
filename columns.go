// Package columns aligns columnar data to a writer.
package columns

import (
	"fmt"
	"io"
	"strings"
)

type Alignment int

const (
	Left Alignment = iota
	Right
)

// A Column of data to output.
type Column struct {
	// Alignment of text within the column.
	Align Alignment
	// MinWidth is the minimum width of the column. Space padding will be added.
	MinWidth int
	// MaxWidth is the maximum width of the column. The column will be truncated to this value.
	MaxWidth int
	// Column data. %v is used to display the value.
	Column []interface{}
}

var (
	alignmentMap = []string{"-", ""}
	newline      = []byte("\n")
)

// Format columns and write to w.
//
// Width is the desired output width. Spacing is the number of spaces separating columns.
func Format(w io.Writer, width, spacing int, columns []*Column) error {
	sp := strings.Repeat(" ", spacing)
	// Compute widths
	sum := 0
	widths := make([]int, len(columns))
	for ci, c := range columns {
		w := c.MinWidth
		for _, r := range c.Column {
			l := len(fmt.Sprintf("%v", r))
			if l > w {
				if c.MaxWidth > 0 && l > c.MaxWidth {
					l = c.MaxWidth
				}
				w = l
			}
		}
		sum += w + spacing
		widths[ci] = w
	}
	sum -= spacing

	for sum > width {
		widest := -1
		// Find widest column...
		for ci, c := range columns {
			if widths[ci] == c.MinWidth {
				continue
			}
			if widest == -1 || widths[ci] > widths[widest] {
				widest = ci
			}
		}
		if widest == -1 {
			break
		}
		// ...and narrow it
		widths[widest]--
		sum--
	}

	for ri := 0; ri < len(columns[0].Column); ri++ {
		row := []string{}
		for ci, c := range columns {
			f := "%" + alignmentMap[c.Align] + "*v"
			w := widths[ci]
			s := fmt.Sprintf(f, w, c.Column[ri])
			if len(s) > w {
				s = s[:w]
			}
			row = append(row, s)
		}
		_, err := w.Write([]byte(strings.Join(row, sp)))
		if err != nil {
			return err
		}
		_, err = w.Write(newline)
		if err != nil {
			return err
		}
	}
	return nil
}
