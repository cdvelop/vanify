package wasm

import (
	"path/filepath"
)

// rootDirectory ej: ./home/user/project
func New(rootDirectory string) *handler {

	h := &handler{
		rootDirectory:  rootDirectory,
		tinyGoCompiler: false,
		wasm_build:     false,
		wasmFileName:   "main.wasm",
	}

	h.MainGoFilePathToBuildWasm = filepath.Join(rootDirectory, "wasm", "main.go")

	h.buildDirectory = filepath.Join(rootDirectory, "cmd", "public", "static")

	h.WasmFileOutPath = filepath.Join(h.buildDirectory, h.wasmFileName)

	h.webAssemblyCheck()

	return h
}
