package gotimelog

import (
	"encoding/json"
	"fmt"
	"time"
)

const EntryDateFormat = "2006-01-02 15:04"

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
	return fmt.Sprintf("%s: %s", e.timestamp.Format(EntryDateFormat), e.title)
}

func (e Entry) Timestamp() time.Time      { return e.timestamp }
func (e *Entry) SetTimestamp(t time.Time) { e.timestamp, e.changed = t, true }

func (e Entry) Title() string      { return e.title }
func (e *Entry) SetTitle(s string) { e.title, e.changed = s, true }

func (e Entry) MarshalJSON() ([]byte, error) {
	var j struct {
		Timestamp string `json:"timestamp"`
		Title     string `json:"title"`
	}

	j.Timestamp = e.timestamp.Format(EntryDateFormat)
	j.Title = e.title

	return json.Marshal(j)
}
