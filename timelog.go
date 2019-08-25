package gotimelog

import (
	"strings"
)

// Timelog represents the contents of a timelog.txt file.
// See: https://gtimelog.org/formats.html#timelog-txt
type Timelog struct {
	Entries []Line
}

func (f *Timelog) Parse(content string) error {
	entries := []Line{}

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		entry := ParseLine(strings.TrimSpace(line))
		entries = append(entries, entry)
	}

	f.Entries = entries
	return nil
}

func (f *Timelog) String() string {
	lines := make([]string, 0, len(f.Entries))

	for _, entry := range f.Entries {
		lines = append(lines, entry.Text())
	}

	return strings.Join(lines, "\n")
}
