package css

import (
	"bytes"

	"github.com/cdvelop/vanify/pkg/minify"
)

type handler struct {
	buildDirectory string
	FilePath       string // ej: ./project/cmd/public/static/style.css

	registeredFiles map[string]*file

	files []*file
	buf   bytes.Buffer
	min   *minify.M
}

type file struct {
	index   int
	path    string
	content []byte
}
