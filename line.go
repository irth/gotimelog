package gotimelog

import "time"

// Line is a single entry (a line) in the timelog.txt file
type Line interface {
	Text() string
}

type lineMeta struct {
	original string
	changed  bool
}

func ParseLine(line string) Line {
	meta := lineMeta{original: line}

	if len(line) > 0 && line[0] == '#' {
		return &Comment{
			contents: line[1:],
			lineMeta: meta,
		}
	}

	l := len(EntryDateFormat)
	if len(line) < l+2 {
		// valid entries contain a timestamp and a title separated by ": "
		return &OldStyleComment{line, meta}
	}

	rawTimestamp, sep, title := line[0:l], line[l:l+2], line[l+2:]

	time, err := time.Parse(EntryDateFormat, rawTimestamp)
	if err != nil {
		return &OldStyleComment{line, meta}
	}

	if sep != ": " {
		return &OldStyleComment{line, meta}
	}

	return &Entry{
		timestamp: time,
		title:     title,
		lineMeta:  meta,
	}
}
