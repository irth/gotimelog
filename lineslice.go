package gotimelog

import (
	"time"
)

type LineSlice []Line

func (l LineSlice) EntriesByRange(start time.Time, end time.Time) []*Entry {
	var matching []*Entry
	for _, line := range l {
		entry, ok := line.(*Entry)
		if !ok {
			continue
		}

		t := entry.Timestamp()
		if t.Equal(start) || t.Equal(end) || (t.After(start) && t.Before(end)) {
			matching = append(matching, entry)
		}
	}
	return matching
}
