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

type Column struct {
	Align    Alignment
	MinWidth int
	MaxWidth int
	// Fit       func(s string, width int) []string
	Column []interface{}
}

var (
	alignmentMap = []string{"-", ""}
	newline      = []byte("\n")
)

func Format(w io.Writer, width int, spacing int, columns []*Column) error {
	sp := strings.Repeat(" ", spacing)
	// Compute widths
	for _, c := range columns {
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
		c.MinWidth = w
	}

	for ri := 0; ri < len(columns[0].Column); ri++ {
		row := []string{}
		for _, c := range columns {
			f := "%" + alignmentMap[c.Align] + "*v"
			s := fmt.Sprintf(f, c.MinWidth, c.Column[ri])
			if c.MaxWidth > 0 && len(s) > c.MaxWidth {
				s = s[:c.MaxWidth]
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
