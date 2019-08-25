package main

import (
	"fmt"

	"github.com/irth/gotimelog"
)

func main() {
	f := gotimelog.Timelog{}
	err := f.LoadFile("/home/me/.local/share/gtimelog/timelog.txt")
	fmt.Printf("%v %+v", err, f.Entries[0])
}
