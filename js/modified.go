package js

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cdvelop/vanify/common"
)

func (h *handler) UpdateFileOnDisk(filePath string) error {
	const e = "UpdateFileOnDisk Js "
	if filePath == "" {
		return nil
	}

	fmt.Println("Compilando JS..." + filePath)

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	//1- read file content from filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(e + err.Error())
	}

	h.UpdateFileContentInMemory(filePath, content)

	h.buf.Reset()

	//  agregamos el js por defecto al inicio
	h.buf.WriteString("'use strict';\n")
	for _, f := range h.files {
		h.buf.Write(f.content)
	}

	// agregamos js wasm según su controlador
	h.buf.Write(h.wasmJs.WasmExecJsContent())

	if err := h.Minify(&h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	if err := common.FileWrite(h.FilePath, h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	return nil
}
