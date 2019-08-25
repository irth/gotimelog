package gotimelog

import (
	"bufio"
	"fmt"
	"io"
)

// Timelog represents the contents of a timelog.txt file.
// See: https://gtimelog.org/formats.html#timelog-txt
type Timelog struct {
	Entries []Line
}

func (f *Timelog) Load(r io.Reader) error {
	br := bufio.NewScanner(r)
	entries := []Line{}

	for {
		if ok := br.Scan(); !ok {
			err := br.Err()
			if err == nil { // EOF
				break
			}
			return err
		}

		entry := ParseLine(br.Text())
		entries = append(entries, entry)
	}

	f.Entries = entries
	return nil
}

func (f *Timelog) Save(w io.Writer) error {
	for _, entry := range f.Entries {
		_, err := fmt.Fprintln(w, entry.Text())
		if err != nil {
			return err
		}
	}
	return nil
}
