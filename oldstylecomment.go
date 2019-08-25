package gotimelog

type OldStyleComment struct {
	contents string
	lineMeta
}

func (o OldStyleComment) Text() string {
	if !o.changed {
		return o.original
	}

	return o.contents
}

func (o OldStyleComment) Contents() string      { return o.contents }
func (o *OldStyleComment) SetContents(s string) { o.contents, o.changed = s, true }
