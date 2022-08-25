package pages

import (
	"io/ioutil"
	"os"

	"github.com/vladovidiu/tldr/pkg/config"
	utils "github.com/vladovidiu/tldr/pkg/utils"
)

var (
	sep = string(os.PathSeparator)
	ext = ".md"
)

func Read(page string) (p *Page, err error) {
	p, err = queryCommon(page)

	if err != nil {
		p, err = queryOS(page)
		if err != nil {
			return p, err
		}
	}

	return p, nil
}

func queryCommon(page string) (p *Page, err error) {
	path := config.SourceDir + sep + "pages" + sep + "common" + sep

	buffer, err := ioutil.ReadFile(path + page + ext)

	if err != nil {
		return p, err
	}

	p = ParsePage(string(buffer))

	return p, nil
}

func queryOS(page string) (p *Page, err error) {
	path := config.SourceDir + sep + "pages" + sep + utils.OSName() + sep

	buffer, err := ioutil.ReadFile(path + page + ext)

	if err != nil {
		return p, err
	}

	p = ParsePage(string(buffer))

	return p, nil
}
