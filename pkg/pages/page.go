package pages

import (
	"regexp"
	"strings"
)

var (
	rarg = regexp.MustCompile(`{{.[^}}]+}}`)
)

type Page struct {
	Name string
	Desc string
	Tips []*Tip
}

type Tip struct {
	Desc string
	Cmd  *Command
}

type Command struct {
	Command string
	Args    []string
}

func ParsePage(s string) *Page {
	lines := strings.Split(s, "\n")

	n := lines[0][2:]

	tips := make([]*Tip, 0)

	p := &Page{
		Name: n,
		Desc: "test",
		Tips: tips,
	}

	return p
}

func (p *Page) String() string {
	s := p.Name + "\n" + p.Desc

	return s
}
