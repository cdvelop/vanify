package js

import (
	"bytes"

	"github.com/cdvelop/vanify/pkg/minify"
)

type handler struct {
	buildDirectory string
	FilePath       string // ej: ./project/cmd/public/static/main.js

	registeredFiles map[string]*file

	wasmJs wasmExecJsContent

	files []*file

	buf bytes.Buffer

	min *minify.M
}

type file struct {
	index   int
	path    string
	content []byte
}

type wasmExecJsContent interface {
	WasmExecJsContent() []byte
}
