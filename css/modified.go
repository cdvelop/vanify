package css

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cdvelop/vanify/common"
)

func (h *handler) FileWasModifiedOnDisk(filePath string) error {
	const this = "FileWasModifiedOnDisk Css "
	if filePath == "" {
		return nil
	}
	fmt.Println("Compilando CSS..." + filePath)

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	//1- read file content from filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New(this + err.Error())
	}

	h.LoadMemoryContent(filePath, content)

	h.buf.Reset()
	for _, f := range h.files {
		h.buf.Write(f.content)
	}

	// fmt.Println("4- >>> escribiendo archivos app.css y style.css")
	if err := h.CssMinify(&h.buf); err != nil {
		return errors.New(this + err.Error())
	}

	if err := common.FileWrite(h.CssFilePath, h.buf); err != nil {
		return errors.New(this + err.Error())
	}

	return nil
}
