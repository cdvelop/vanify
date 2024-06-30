package wasm

import (
	"path/filepath"

	"github.com/cdvelop/vanify/gomod"
)

// rootDirectory ej: ./home/user/project
func New(rootDirectory string, f flags) *handler {

	h := &handler{
		rootDirectory:   rootDirectory,
		tinyGoCompiler:  false,
		wasmProjectType: false,
		wasmFileName:    "main.wasm",
		flags:           f,
	}

	h.MainGoFilePathToBuildWasm = filepath.Join(rootDirectory, "wasm", "main.go")

	h.buildDirectory = filepath.Join(rootDirectory, "cmd", "public", "static")

	h.WasmFileOutPath = filepath.Join(h.buildDirectory, h.wasmFileName)

	if err := h.webAssemblyCheck(); err != nil {
		panic(err)
	}

	if _, err := gomod.Exist(rootDirectory); err == nil {
		h.gomodExist = true
	}

	return h
}
