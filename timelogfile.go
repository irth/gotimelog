package gotimelog

import (
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type TimelogFile struct {
	Timelog
	sync.RWMutex
	Path string
}

func (f *TimelogFile) Load() error {
	file, err := os.OpenFile(f.Path, os.O_RDONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "opening timelog.txt for reading")
	}
	return f.Timelog.Load(file)
}

func (f *TimelogFile) Save() error {
	file, err := os.OpenFile(f.Path, os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "opening timelog.txt for writing")
	}
	return f.Timelog.Save(file)
}

func (f *TimelogFile) Append(l Line) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND, 0644)
	if err != nil {
		return errors.Wrap(err, "opening timelog.txt for appending")
	}

	_, err = fmt.Fprintln(file, l.Text())
	return err
}
