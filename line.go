package configurit

import (
	"fmt"
	"strings"
)

type line struct {
	Number int
	Line   string
}

func (l line) Print() {
	fmt.Printf("%03d|  %s\n", l.Number, l.Line)
}

func createLine(n int, l string) line {
	return line{n, l}
}

func (l line) isValue(s string) (bool, string) {
	t := strings.Split(l.Line, "=")
	if len(t) <= 0 {
		return false, ""
	}
	if strings.ToLower(strings.TrimSpace(t[0])) == strings.ToLower(s) {
		return true, strings.TrimSpace(t[1])
	}
	return false, ""
}
