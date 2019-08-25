package gotimelog

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const dateFormat = "2006-01-02 15:04"

type File struct {
	Entries []*Entry
}

type Entry struct {
	Timestamp time.Time
	Title     string
}

func (f *File) LoadFile(path string) error {
	rawContents, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "loading timelog.txt")
	}
	return f.Load(string(rawContents))
}

func (f *File) Load(content string) error {
	entries := []*Entry{}

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		entry, ok := ParseEntry(strings.TrimSpace(line))
		if !ok {
			continue
		}
		entries = append(entries, entry)
	}

	f.Entries = entries
	return nil
}

func ParseEntry(line string) (*Entry, bool) {
	if len(line) < len(dateFormat)+2 {
		// valid entries contain a timestamp and a title separated by ": "
		return nil, false
	}

	if line[0] == '#' {
		return nil, false
	}

	rawTimestamp, sep, title := line[0:len(dateFormat)], line[len(dateFormat):len(dateFormat)+2], line[len(dateFormat)+2:]

	time, err := time.Parse("2006-01-02 15:04", rawTimestamp)
	if err != nil {
		return nil, false
	}

	if sep != ": " {
		return nil, false
	}

	return &Entry{
		Timestamp: time,
		Title:     title,
	}, true
}
