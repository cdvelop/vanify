package css

import (
	"bytes"
	"path/filepath"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/compiler/minify"
	min "github.com/cdvelop/vanify/compiler/minify/css"
)

// publicCssDirectory ej: ./project/cmd/public/static
func New(publicCssDirectory string) *handler {

	err := common.CheckDirectories(publicCssDirectory)
	if err != nil {
		panic("css New " + err.Error())
	}

	m := minify.New()
	m.AddFunc("text/css", min.Minify)

	return &handler{
		buildDirectory:  publicCssDirectory,
		CssFilePath:     filepath.Join(publicCssDirectory, "style.css"),
		registeredFiles: map[string]*file{},
		files:           []*file{},
		buf:             bytes.Buffer{},
		min:             m,
	}

}
