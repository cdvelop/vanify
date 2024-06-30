package svg

import (
	"bytes"
	"encoding/xml"
	"errors"
	"os"
)

func (h *handler) writeToFile() error {

	const e = "writeToFile SVG "

	// reset def symbols
	h.sprite.Defs = defs{Symbols: []symbol{}}

	for _, icon := range h.icons {
		h.sprite.Defs.Symbols = append(h.sprite.Defs.Symbols, icon.symbol)
	}

	output, err := xml.MarshalIndent(h.sprite, "", "  ")
	if err != nil {
		return errors.New(e + err.Error())
	}

	buf := bytes.NewBuffer(output)

	if err := h.Minify(buf); err != nil {
		return errors.New(e + err.Error())
	}

	// fmt.Println("contenido:", string(output))

	return os.WriteFile(h.outFilePath, buf.Bytes(), 0644)
}
