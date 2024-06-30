package wasm

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
)

func (h *handler) setWasmExecJsPathTinyGo() error {

	path, er := exec.LookPath("tinygo")
	if er != nil {
		return errors.New("TinyGo no encontrado en el PATH. " + er.Error())
	}
	// Obtener el directorio de instalación
	tinyGoDir := filepath.Dir(path)

	// Limpiar la ruta y quitar "\bin"
	tinyGoDir = strings.TrimSuffix(tinyGoDir, "\\bin")

	// Construir la ruta completa al archivo wasm_exec.js
	h.wasmExecJsPathFile = filepath.Join(tinyGoDir, "targets", "wasm_exec.js")

	return nil
}

func (h *handler) setWasmExecJsPathGo() error {

	// Obtener la ruta del directorio de instalación de Go desde la variable de entorno GOROOT
	path, er := exec.LookPath("go")
	if er != nil {
		return errors.New("Go no encontrado en el PATH. " + er.Error())
	}

	// Obtener el directorio de instalación
	GoDir := filepath.Dir(path)

	// Limpiar la ruta y quitar "\bin"
	GoDir = strings.TrimSuffix(GoDir, "\\bin")

	// Construir la ruta completa al archivo wasm_exec.js
	h.wasmExecJsPathFile = filepath.Join(GoDir, "misc", "wasm", "wasm_exec.js")

	return nil
}
