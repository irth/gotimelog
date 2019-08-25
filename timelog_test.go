package gotimelog_test

import (
	"testing"
	"time"

	"github.com/irth/gotimelog"
	"github.com/stretchr/testify/assert"
)

func TestParseEntry(t *testing.T) {
	entry, ok := gotimelog.ParseEntry("2021-04-14 21:37: title: kek")
	assert.True(t, ok)
	assert.Equal(t, time.Date(2021, 04, 14, 21, 37, 0, 0, time.UTC), entry.Timestamp)
	assert.Equal(t, "title: kek", entry.Title)

	_, ok = gotimelog.ParseEntry("2021-04-42 21:37: title: kek")
	assert.False(t, ok)
	_, ok = gotimelog.ParseEntry("2021-04-14 21:37 title: kek")
	assert.False(t, ok)
	_, ok = gotimelog.ParseEntry("2021-04-42 21:37")
	assert.False(t, ok)
	_, ok = gotimelog.ParseEntry("foo: bar")
	assert.False(t, ok)
	_, ok = gotimelog.ParseEntry("# kek")
	assert.False(t, ok)
	_, ok = gotimelog.ParseEntry("")
	assert.False(t, ok)
}

func TestLoad(t *testing.T) {
	f := gotimelog.File{}
	err := f.Load(`
2009-10-11 12:13: kek

# test
2010-11-14 21:38: nightdrive
	`)

	assert.NoError(t, err)
	assert.Len(t, f.Entries, 2)

	assert.Equal(t, time.Date(2009, 10, 11, 12, 13, 0, 0, time.UTC), f.Entries[0].Timestamp)
	assert.Equal(t, "kek", f.Entries[0].Title)

	assert.Equal(t, time.Date(2010, 11, 14, 21, 38, 0, 0, time.UTC), f.Entries[1].Timestamp)
	assert.Equal(t, "nightdrive", f.Entries[1].Title)
}
