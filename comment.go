package gotimelog

import "fmt"

type Comment struct {
	contents string
	lineMeta
}

func (c Comment) Text() string {
	if !c.changed {
		return c.original
	}
	return fmt.Sprintf("#%s", c.contents)
}

func (c Comment) Contents() string      { return c.contents }
func (c *Comment) SetContents(s string) { c.contents, c.changed = s, true }
