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
	return fmt.Sprintf("%s %s%s", NAME, VERSION, TAG)
}

type conf struct {
	Config map[string]map[string]string
	Name   string
}

func Open(path string) (*conf, error) {
	c := new(conf)
	c.Config = make(map[string]map[string]string)
	c.makeSection("")
	c.Name = path
	e := c.readConfig(path)
	return c, e
}

func (c *conf) makeSection(key string) {
	c.Config[key] = make(map[string]string)
}

func (c *conf) readConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Configurit: %s", err)
	}
	scanner := bufio.NewScanner(file)
	curSection := ""
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			if text[:1] == "#" {
				continue
			} else if text[:1] == ";" {
				curSection = strings.ToLower(strings.TrimSpace(text[1:]))
				if curSection != "" {
					c.makeSection(curSection)
				}
			} else {
				key, value, err := keyandValue(text)
				if err != nil {
					return err
				}
				c.Config[curSection][key] = value
			}
		}
	}
	if scanner.Err() != nil {
		return fmt.Errorf("Configurit: %s", scanner.Err())
	}
	return nil
}

func keyandValue(line string) (key string, value string, err error) {
	t := strings.Split(strings.TrimSpace(line), "=")
	if len(t) <= 1 {
		return "", "", fmt.Errorf("Configurit: Got empty key/value from line: %s", line)
	}

	key = strings.ToLower(strings.TrimSpace(t[0]))

	if len(t) > 2 {
		value = strings.Join(t[1:], "=")
	} else {
		value = strings.TrimSpace(t[1])
	}

	return key, value, nil
}

func (c *conf) Print() {
	fmt.Println()
	for i := range c.Config {
		fmt.Println("Section: ", i)
		for j := range c.Config[i] {
			fmt.Printf("  Key: %s | Value: %s\n", j, c.Config[i][j])
		}
	}
}

func (c conf) get(section string, key string) (string, error) {
	if a, ok := c.Config[strings.ToLower(section)][strings.ToLower(key)]; ok != false {
		return a, fmt.Errorf("Configurit: Unable to find value for key %s in section %s", key, section)
	} else {
		return a, nil
	}
}

func (c conf) GetInt(section string, key string) (int, error) {
	return 0, nil
}

func (c conf) GetFloat32(section string, key string) (float32, error) {
	return 0, nil
}

func (c conf) GetFloat64(section string, key string) (float64, error) {
	return 0, nil
}

func (c conf) GetString(section string, key string) (string, error) {
	return "", nil
}

func (c conf) GetBool(section string, key string) (bool, error) {
	return false, nil
}
