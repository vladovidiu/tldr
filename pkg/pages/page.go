package pages

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	rarg    = regexp.MustCompile(`{{.[^}}]+}}`)
	bold    = color.New(color.Bold)
	blue    = color.New(color.FgBlue)
	red     = color.New(color.FgRed)
	magenta = color.New(color.FgMagenta)
	white   = color.New(color.FgWhite)
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

	var desc string
	var count int

	for ln := 2; ln < len(lines); ln++ {
		line := lines[ln]
		if len(line) > 0 && line[0] == '>' {
			desc = desc + line[2:] + "\n"
		} else {
			count = ln
			break
		}
	}

	tips := make([]*Tip, 0)
	for ln := count; ln < len(lines); {
		line := lines[ln]

		if len(line) > 0 && line[0] == '-' {
			d := line[2:]
			c := lines[ln+2]

			var cmd *Command
			if len(c) > 0 && c[0] == '`' {
				cmd = &Command{
					Command: c[:len(c)-1][1:],
					Args:    rarg.FindAllString(c, -1),
				}
				ln = ln + 2
			} else {
				break
			}

			tips = append(tips, &Tip{
				Desc: d,
				Cmd:  cmd,
			})
		}

		ln++
	}

	p := &Page{
		Name: magenta.Sprint(n),
		Desc: desc,
		Tips: tips,
	}

	return p
}

func (p *Page) String() string {
	s := bold.Sprint(p.Name) + "\n\n" + p.Desc + "\n"

	for _, t := range p.Tips {
		s = s + t.Display()
	}

	return s
}

func (t *Tip) String() string {
	return t.Desc
}

func (c *Command) String() string {
	return c.Command
}

func (p *Page) Display() string {
	return bold.Sprint(p.Name) + "\n\n" + p.Desc
}

func (t *Tip) Display() string {
	return "- " + blue.Sprint(t.Desc) + "\n" + t.Cmd.Display()
}

func (c *Command) Display() string {
	s := c.Command

	for _, arg := range c.Args {
		t := arg[2:]
		t = t[:len(t)-2]
		s = strings.Replace(s, arg, white.Sprint(bold.Sprint(t)), 1)
	}

	return "    " + s + "\n"
}
