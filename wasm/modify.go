package wasm

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func (h *handler) UpdateFileOnDisk(event_name string) error {
	const this = "UpdateFileOnDisk error "
	if !h.wasmProjectType {
		return nil
	}

	if event_name != "" {
		fmt.Println("Compilando WASM..." + event_name)
	}

	// Ejecutar go mod tidy en el directorio del proyecto
	if h.gomodExist {
		tidyCmd := exec.Command("go", "mod", "tidy")
		tidyCmd.Dir = h.rootDirectory
		tidyOutput, tidyErr := tidyCmd.CombinedOutput()
		if tidyErr != nil {
			return errors.New(this + "al ejecutar 'go mod tidy': " + tidyErr.Error() + " " + string(tidyOutput))
		}
	}

	var cmd *exec.Cmd

	// log.Println("*** c.e2eWasmTestFolder: ", c.e2eWasmTestFolder)

	// delete last file
	os.Remove(h.WasmFileOutPath)

	var flags string
	if h.flags != nil {
		flags = h.Flags()
	}

	// log.Println("*** MainGoFilePathToBuildWasm: ", MainGoFilePathToBuildWasm)
	// Ajustamos los parámetros de compilación según la configuración
	if h.tinyGoCompiler {
		// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
		cmd = exec.Command("tinygo", "build", "-o", h.WasmFileOutPath, "-target", "wasm", "--no-debug", "-ldflags", flags, h.MainGoFilePathToBuildWasm)

	} else {
		// compilación normal...

		cmd = exec.Command("go", "build", "-o", h.WasmFileOutPath, "-tags", "dev", "-ldflags", flags, "-v", h.MainGoFilePathToBuildWasm)
		// cmd = exec.Command("go", "build", "-o", WasmFileOutPath, "-tags", "dev", "-ldflags", "-s -w", "-v", MainGoFilePathToBuildWasm)
		cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	}

	output, er := cmd.CombinedOutput()
	if er != nil {
		return errors.New("error al compilar a WebAssembly error: " + er.Error() + " string(output):" + string(output))
	}

	// Verificamos si el archivo wasm se creó correctamente
	if _, er := os.Stat(h.WasmFileOutPath); er != nil {
		return errors.New("error el archivo WebAssembly no se creó correctamente: " + er.Error())
	}

	// fmt.Printf("WebAssembly compilado correctamente y guardado en %s\n", WasmFileOutPath)

	return nil
}
