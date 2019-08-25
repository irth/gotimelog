package gotimelog

import (
	"bufio"
	"fmt"
	"io"
)

// Timelog represents the contents of a timelog.txt file.
// See: https://gtimelog.org/formats.html#timelog-txt
type Timelog struct {
	Lines LineSlice
}

func (f *Timelog) Load(r io.Reader) error {
	br := bufio.NewScanner(r)
	lines := []Line{}

	for {
		if ok := br.Scan(); !ok {
			err := br.Err()
			if err == nil { // EOF
				break
			}
			return err
		}

		entry := ParseLine(br.Text())
		lines = append(lines, entry)
	}

	f.Lines = lines
	return nil
}

func (f *Timelog) Save(w io.Writer) error {
	for _, line := range f.Lines {
		_, err := fmt.Fprintln(w, line.Text())
		if err != nil {
			return err
		}
	}
	return nil
}
