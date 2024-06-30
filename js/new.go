package js

import (
	"bytes"
	"path/filepath"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/pkg/minify"
	min "github.com/cdvelop/vanify/pkg/minify/js"
)

// publicJsDirectory ej: ./project/cmd/public/static
func New(publicJsDirectory string, w wasmExecJsContent) *handler {

	err := common.CheckDirectories(publicJsDirectory)
	if err != nil {
		panic("js New " + err.Error())
	}

	m := minify.New()
	m.AddFunc("text/javascript", min.Minify)

	return &handler{
		buildDirectory:  publicJsDirectory,
		FilePath:        filepath.Join(publicJsDirectory, "main.js"),
		registeredFiles: map[string]*file{},
		wasmJs:          w,
		files:           []*file{},
		buf:             bytes.Buffer{},
		min:             m,
	}

}
