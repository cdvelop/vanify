package wasm

import (
	"fmt"
	"os"
)

func (h *handler) webAssemblyCheck() error {
	// chequear si existe wasm_main.go en la ruta de trabajo ej: cmd/frontend/wasm_main.go

	// fmt.Println("CARPETA DE TRABAJO: ", c.FrontBuildRootFolder, " ARCHIVO WASM: ", c.wasmFileName)

	if _, err := os.Open(h.MainGoFilePathToBuildWasm); err == nil {
		h.wasmProjectType = true

		var wasm_compiler_name string

		var remove_message string

		if h.tinyGoCompiler {

			if err = h.setWasmExecJsPathTinyGo(); err != nil {
				return err
			}

			// remove the message: syscall/js.finalizeRef not implemented
			//  https://github.com/tinygo-org/tinygo/issues/1140
			remove_message = `go.importObject.env["syscall/js.finalizeRef"] = () => {}`

			wasm_compiler_name = "TinyGo"

		} else {
			if err = h.setWasmExecJsPathGo(); err != nil {
				return err
			}

			wasm_compiler_name = "Go"

		}

		h.wasmImportJS = []byte(`const go = new Go();
		` + remove_message + `
		WebAssembly.instantiateStreaming(fetch("static/` + h.wasmFileName + `"), go.importObject).then((result) => {
			go.run(result.instance);
		});`)

		// Leemos el contenido del archivo
		wasm_js_content, err := os.ReadFile(h.wasmExecJsPathFile)
		if err != nil {
			return err
		}

		h.wasmImportJS = append(h.wasmImportJS, wasm_js_content...)

		fmt.Printf("*** Proyecto WebAssembly: [%v] Compilador: [%v] ***\n", h.wasmFileName, wasm_compiler_name)
	} else {
		fmt.Println("No se encontró el archivo: ", h.MainGoFilePathToBuildWasm)
	}

	return nil
}
