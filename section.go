package configurit

import (
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	Lines []line
	Name  string
}

func createSection(n string) section {
	s := section{}
	s.Name = n
	s.Lines = make([]line, 0)
	return s
}

func (s *section) add(n int, l string) {
	s.Lines = append(s.Lines, createLine(n, l))
}

func (s *section) Print() {
	fmt.Println("#####======#####")
	fmt.Printf("%s\n", s.Name)
	fmt.Println("#####======#####")
	for i := range s.Lines {
		s.Lines[i].Print()
	}
}

func (s section) GetValue(n string) (int, error) {
	var ok bool
	var value string
	for i := range s.Lines {
		ok, value = s.Lines[i].isValue(n)
		if ok {
			v, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				return 0, fmt.Errorf("Configurit: %s\n", err)
			}
			return v, nil
		}
	}
	return 0, fmt.Errorf("Configurit: No value found for %s in %s\n", n, s.Name)
}

func (s section) GetString(n string) (string, error) {
	var ok bool
	var value string
	for i := range s.Lines {
		ok, value = s.Lines[i].isValue(n)
		if ok {
			return value, nil
		}
	}
	return "", fmt.Errorf("Configurit: No value found for %s in %s\n", n, s.Name)
}

func (s section) GetBool(n string) (bool, error) {
	var ok bool
	var value string
	for i := range s.Lines {
		ok, value = s.Lines[i].isValue(n)
		if ok {
			if strings.ToLower(value) == "true" {
				return true, nil
			} else {
				return false, nil
			}
		}
	}
	return false, fmt.Errorf("Configurit: No value found for %s in %s\n", n, s.Name)
}
