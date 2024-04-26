package wasm

import (
	"fmt"
	"os"
)

func (h *handler) webAssemblyCheck() {
	// chequear si existe wasm_main.go en la ruta de trabajo ej: cmd/frontend/wasm_main.go

	// fmt.Println("CARPETA DE TRABAJO: ", c.FrontBuildRootFolder, " ARCHIVO WASM: ", c.wasmFileName)

	_, err := os.Open(h.MainGoFilePathToBuildWasm)
	if err == nil {
		var wasm_compiler_name string

		h.wasm_build = true

		var remove_message string

		if h.tinyGoCompiler {

			// remove the message: syscall/js.finalizeRef not implemented
			//  https://github.com/tinygo-org/tinygo/issues/1140
			remove_message = `go.importObject.env["syscall/js.finalizeRef"] = () => {}`

			wasm_compiler_name = "TinyGo"

		} else {

			wasm_compiler_name = "Go"

		}

		h.wasmImportJS = `const go = new Go();
		` + remove_message + `
		WebAssembly.instantiateStreaming(fetch("static/` + h.wasmFileName + `"), go.importObject).then((result) => {
			go.run(result.instance);
		});`

		fmt.Printf("*** Proyecto WebAssembly: [%v] Compilador: [%v] ***\n", h.wasmFileName, wasm_compiler_name)
	}
}
