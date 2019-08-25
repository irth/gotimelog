package gotimelog

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

type TimelogFile struct {
	Timelog
	Path string
}

func (f *TimelogFile) Load() error {
	rawContents, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return errors.Wrap(err, "loading timelog.txt")
	}
	return f.Timelog.Parse(string(rawContents))
}

func (f *TimelogFile) Save() error {
	return ioutil.WriteFile(f.Path, []byte(f.String()), 0644)
}
