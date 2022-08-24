package pages

import (
	"io/ioutil"
	"os"

	"github.com/vladovidiu/tldr/pkg/config"
)

var (
	sep = string(os.PathSeparator)
	ext = ".md"
)

func Read(page string) (p *Page, err error) {
	p, err = queryCommon(page)

	if err != nil {
		p, err = queryOs(page)

		if err != nil {
			return p, err
		}
	}

	return p, nil
}

func queryCommon(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + "common" + sep
	b, err := ioutil.ReadFile(d + page + ext)

	if err != nil {
		return p, err
	}

	p = ParsePage(string(b))

	return p, nil
}

func queryOs(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + "osx" + sep
	b, err := ioutil.ReadFile(d + page + ext)

	if err != nil {
		return p, err
	}

	p = ParsePage(string(b))

	return p, nil
}
