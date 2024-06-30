package index

import (
	"bytes"

	"github.com/cdvelop/vanify/pkg/minify"
)

type handler struct {
	themeDir string //ej: project/website/theme
	FilePath string // ej: ./project/website/build/public/index.html

	ver appVersion

	registeredFiles map[string]*file

	// files []*file
	buf bytes.Buffer
	min *minify.M
}

type appVersion interface {
	AppVersion() string
}

type file struct {
	index   int
	path    string
	content []byte
}

type module struct {
	filePath string
	content  []byte
}
