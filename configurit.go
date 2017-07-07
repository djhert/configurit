package configurit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NAME    = "configurit"
	VERSION = "0.8"
	TAG     = "-alpha"
)

func Version() string {
	return fmt.Sprintf("%s %s%s\n", NAME, VERSION, TAG)
}

type Conf struct {
	Sections []section
	Name     string
}

func CreateFrom(path string) (*Conf, error) {
	c := new(Conf)
	c.Sections = []section{createSection(""), createSection("comment")}
	e := c.ReadConfig(path)
	return c, e
}

func (c *Conf) ReadConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Configurit: %s", err)
	}
	fullText := make([]string, 0)
	scanner := bufio.NewScanner(file)
	i := 0
	curSection := 0
	for scanner.Scan() {
		i++
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			if text[:1] == "#" {
				c.Sections[1].add(i, text)
			} else if text[:1] == ";" {
				if curSection == 0 {
					curSection = 2
				} else {
					curSection++
				}
				c.Sections = append(c.Sections, createSection(strings.TrimSpace(text[1:])))
			} else {
				c.Sections[curSection].add(i, text)
			}
			fullText = append(fullText, text)
		}
	}
	c.Name = path
	if scanner.Err() == nil {
		return nil
	}
	return fmt.Errorf("Configurit: %s", scanner.Err())
}

func (c Conf) GetSection(name string) (*section, error) {
	for i := range c.Sections {
		if c.Sections[i].Name == name {
			return &c.Sections[i], nil
		}
	}
	return nil, fmt.Errorf("Configurit: No Section called %s\n", name)
}

func (c Conf) GetSectionNames() []string {
	s := make([]string, 0)
	for i := range c.Sections {
		s = append(s, c.Sections[i].Name)
	}
	return s
}

func (c *Conf) Print() {
	fmt.Println()
	for i := range c.Sections {
		fmt.Println()
		c.Sections[i].Print()
	}
}
