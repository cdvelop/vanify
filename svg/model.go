package svg

import (
	"encoding/xml"

	"github.com/cdvelop/vanify/pkg/minify"
)

type handler struct {
	outFilePath string //ej: project/website/build/assets/sprite.svg
	sprite      *sprite
	icons       map[string]icon

	min *minify.M
}

type icon struct {
	content string
	symbol  symbol
}

type sprite struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Role    string   `xml:"role,attr"`
	Hidden  string   `xml:"aria-hidden,attr"`
	Focus   string   `xml:"focusable,attr"`
	Class   string   `xml:"class,attr"`
	Defs    defs     `xml:"defs"`
}

type defs struct {
	Symbols []symbol `xml:"symbol"`
}

type symbol struct {
	XMLName xml.Name `xml:"symbol"`
	ID      string   `xml:"id,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Path    path     `xml:"path"`
}

type path struct {
	Fill string `xml:"fill,attr"`
	D    string `xml:"d,attr"`
}
