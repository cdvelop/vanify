package wasm

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
)

func (h *handler) wasmExecJsPathTinyGo() (string, error) {

	path, er := exec.LookPath("tinygo")
	if er != nil {
		return "", errors.New("TinyGo no encontrado en el PATH. " + er.Error())
	}
	// fmt.Println("RUTA OBTENIDA:", path)
	// Obtener el directorio de instalación
	tinyGoDir := filepath.Dir(path)

	// Limpiar la ruta y quitar "\bin"
	tinyGoDir = strings.TrimSuffix(tinyGoDir, "\\bin")

	// Construir la ruta completa al archivo wasm_exec.js
	wasmExecPath := filepath.Join(tinyGoDir, "targets", "wasm_exec.js")

	// fmt.Println("Ruta de instalación de TinyGo:", tinyGoDir)
	// fmt.Println("Ruta completa de wasm_exec.js:", wasmExecPath)

	return wasmExecPath, nil
}

func (h *handler) wasmExecJsPathGo() (string, error) {

	// Obtener la ruta del directorio de instalación de Go desde la variable de entorno GOROOT
	path, er := exec.LookPath("go")
	if er != nil {
		return "", errors.New("Go no encontrado en el PATH. " + er.Error())
	}

	// fmt.Println("RUTA OBTENIDA:", path)
	// Obtener el directorio de instalación
	GoDir := filepath.Dir(path)

	// Limpiar la ruta y quitar "\bin"
	GoDir = strings.TrimSuffix(GoDir, "\\bin")

	// Construir la ruta completa al archivo wasm_exec.js
	wasmExecPath := filepath.Join(GoDir, "misc", "wasm", "wasm_exec.js")

	// fmt.Println("Ruta completa de wasm_exec.js:", wasmExecPath)

	return wasmExecPath, nil
}
