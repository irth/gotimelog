package gotimelog

import (
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// Timelog represents the contents of a timelog.txt file.
// See: https://gtimelog.org/formats.html#timelog-txt
type Timelog struct {
	Entries []Line
}

func (f *Timelog) LoadFile(path string) error {
	rawContents, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "loading timelog.txt")
	}
	return f.Load(string(rawContents))
}

func (f *Timelog) Load(content string) error {
	entries := []Line{}

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		entry := ParseLine(strings.TrimSpace(line))
		entries = append(entries, entry)
	}

	f.Entries = entries
	return nil
}
