package css

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cdvelop/vanify/common"
)

func (h *handler) UpdateFileOnDisk(filePath string) error {
	const e = "UpdateFileOnDisk Css "
	if filePath == "" {
		return nil
	}
	fmt.Println("Compilando CSS..." + filePath)

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	//1- read file content from filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(e + err.Error())
	}

	h.UpdateFileContentInMemory(filePath, content)

	h.buf.Reset() // update and join content files
	for _, f := range h.files {
		h.buf.Write(f.content)
	}

	if err := h.Minify(&h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	// fmt.Println("escribiendo", h.buf.String())

	if err := common.FileWrite(h.FilePath, h.buf); err != nil {
		return errors.New(e + err.Error())
	}

	return nil
}
