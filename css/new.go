package css

import (
	"bytes"
	"path/filepath"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/pkg/minify"
	min "github.com/cdvelop/vanify/pkg/minify/css"
)

// publicCssDirectory ej: ./user/project/build/public/assets
func New(publicCssDirectory string) *handler {

	err := common.CheckDirectories(publicCssDirectory)
	if err != nil {
		panic("css New " + err.Error())
	}

	m := minify.New()
	m.AddFunc("text/css", min.Minify)

	return &handler{
		buildDirectory:  publicCssDirectory,
		FilePath:        filepath.Join(publicCssDirectory, "style.css"),
		registeredFiles: map[string]*file{},
		files:           []*file{},
		buf:             bytes.Buffer{},
		min:             m,
	}

}
