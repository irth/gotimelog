package gotimelog_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/irth/gotimelog"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	line := gotimelog.ParseLine("2021-04-14 21:37: title: kek")
	entry, ok := line.(gotimelog.Entry)
	assert.True(t, ok)
	assert.Equal(t, time.Date(2021, 04, 14, 21, 37, 0, 0, time.UTC), entry.Timestamp())
	assert.Equal(t, "title: kek", entry.Title())

	for _, line := range []string{
		"2021-04-42 21:37: title: kek",
		"2021-04-14 21:37 title: kek",
		"2021-04-42 21:37",
		"foo: bar",
		"",
	} {
		comment, ok := gotimelog.ParseLine(line).(gotimelog.OldStyleComment)
		assert.True(t, ok)
		assert.Equal(t, line, comment.Contents())
	}

	comment, ok := gotimelog.ParseLine("# kek").(gotimelog.Comment)
	assert.True(t, ok)
	assert.Equal(t, " kek", comment.Contents())
}

func TestLoad(t *testing.T) {
	f := gotimelog.Timelog{}
	err := f.Load(bytes.NewReader(
		[]byte(`2009-10-11 12:13: kek

# test
2010-11-14 21:38: nightdrive`),
	))

	assert.NoError(t, err)
	assert.Len(t, f.Lines, 4)

	e1, ok := f.Lines[0].(*gotimelog.Entry)
	assert.True(t, ok)
	e2, ok := f.Lines[1].(*gotimelog.OldStyleComment)
	assert.True(t, ok)
	e3, ok := f.Lines[2].(*gotimelog.Comment)
	assert.True(t, ok)
	e4, ok := f.Lines[3].(*gotimelog.Entry)
	assert.True(t, ok)

	assert.Equal(t, time.Date(2009, 10, 11, 12, 13, 0, 0, time.UTC), e1.Timestamp())
	assert.Equal(t, "kek", e1.Title())

	assert.Equal(t, "", e2.Contents())
	assert.Equal(t, " test", e3.Contents())

	assert.Equal(t, time.Date(2010, 11, 14, 21, 38, 0, 0, time.UTC), e4.Timestamp())
	assert.Equal(t, "nightdrive", e4.Title())
}
