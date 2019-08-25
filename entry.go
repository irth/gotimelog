package gotimelog

import (
	"fmt"
	"time"
)

const dateFormat = "2006-01-02 15:04"

// Entry is a single timestamped entry in the timelog.txt file
type Entry struct {
	timestamp time.Time
	title     string

	lineMeta
}

func (e Entry) Text() string {
	if !e.changed {
		return e.original
	}
	return fmt.Sprintf("%s: %s", e.timestamp.Format(dateFormat), e.title)
}

func (e Entry) Timestamp() time.Time      { return e.timestamp }
func (e *Entry) SetTimestamp(t time.Time) { e.timestamp, e.changed = t, true }

func (e Entry) Title() string      { return e.title }
func (e *Entry) SetTitle(s string) { e.title, e.changed = s, true }
