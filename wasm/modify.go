package wasm

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cdvelop/vanify/ldflags"
)

func (h *handler) FileWasModifiedOnDisk(event_name string) error {
	const this = "FileWasModifiedOnDisk error "
	if !h.wasm_build {
		return nil
	}

	if event_name != "" {
		fmt.Println("Compilando WASM..." + event_name)
	}

	// fmt.Println("c.FrontBuildRootFolder:", c.FrontBuildRootFolder)

	// Ejecutar go mod tidy en el directorio del proyecto
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = h.rootDirectory
	tidyOutput, tidyErr := tidyCmd.CombinedOutput()
	if tidyErr != nil {
		return errors.New(this + "al ejecutar 'go mod tidy': " + tidyErr.Error() + " " + string(tidyOutput))
	}

	var cmd *exec.Cmd

	// log.Println("*** c.e2eWasmTestFolder: ", c.e2eWasmTestFolder)

	// delete last file
	os.Remove(h.WasmFileOutPath)

	flags, err := ldflags.Add(h.GetTwoPublicKeysWasmClientAndGoServer())
	if err != "" {
		return errors.New(this + err)
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
