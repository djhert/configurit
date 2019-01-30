package configurit

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type Conf struct {
	config map[string]map[string]string
	Name   string
}

func Open(path string) (*Conf, error) {
	c := new(Conf)
	c.config = make(map[string]map[string]string)
	c.makeSection("")
	c.Name = path
	e := c.readConfig(path)
	return c, e
}

func (c *Conf) makeSection(key string) {
	c.config[key] = make(map[string]string)
}

func (c *Conf) readConfig(path string) error {
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
				c.config[curSection][key] = value
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

func (c *Conf) Print() {
	fmt.Println()
	for i := range c.config {
		fmt.Println("Section: ", i)
		for j := range c.config[i] {
			fmt.Printf("  Key: %s | Value: %s\n", j, c.config[i][j])
		}
	}
}

func (c Conf) get(section string, key string) (string, error) {
	if a, ok := c.config[strings.ToLower(section)][strings.ToLower(key)]; !ok {
		return a, fmt.Errorf("Configurit: Unable to find value for key %s in section %s", key, section)
	} else {
		return a, nil
	}
}

func (c Conf) set(section string, key string, value string) error {
	if _, ok := c.config[strings.ToLower(section)][strings.ToLower(key)]; !ok {
		return fmt.Errorf("Configurit: Unable to find value for key %s in section %s", key, section)
	}
	c.config[strings.ToLower(section)][strings.ToLower(key)] = value
	return nil
}

func (c Conf) GetInt(section string, key string) (int, error) {
	val, err := c.get(section, key)
	if err != nil {
		return 0, err
	}
	o, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return o, nil
}

func (c Conf) GetFloat32(section string, key string) (float32, error) {
	val, err := c.get(section, key)
	if err != nil {
		return 0, err
	}
	o, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0, err
	}
	return float32(o), nil
}

func (c Conf) GetFloat64(section string, key string) (float64, error) {
	val, err := c.get(section, key)
	if err != nil {
		return 0, err
	}
	o, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return o, nil
}

func (c Conf) GetString(section string, key string) (string, error) {
	val, err := c.get(section, key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c Conf) GetBool(section string, key string) (bool, error) {
	val, err := c.get(section, key)
	if err != nil {
		return false, err
	}
	o, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return o, nil
}

/*

func (c conf) SetInt(section string, key string, value int) error {
	return nil
}

func (c conf) SetFloat32(section string, key string, value float32) error {
	return nil
}

func (c conf) SetFloat64(section string, key string, value float64) error {
	return nil
}

func (c conf) SetString(section string, key string, value string) error {
	return nil
}

func (c conf) SetBool(section string, key string, value bool) error {
	return nil
}*/
