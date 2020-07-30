package generate

import (
	"bufio"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

var structName []string

func (g *generate) ReadFile() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error when read file\n")
		}
	}()
	content, err := ioutil.ReadFile(g.path)
	if err != nil {
		return errors.New("file invalid\n")
	}
	strContent := string(content)

	scanner := bufio.NewScanner(strings.NewReader(strContent))
	for scanner.Scan() {
		checkStructName(scanner.Text())
	}
	g.SetStructName(structName)
	return nil
}

func checkStructName(text string) {
	var rgx = regexp.MustCompile(`\Atype\s+[A-Za-z0-9]+\sstruct`)

	rs := rgx.FindStringSubmatch(text)
	if len(rs) == 0 {
		return
	}
	s := strings.Fields(rs[0])
	structName = append(structName, s[1])
}
