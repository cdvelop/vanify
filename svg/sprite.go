package svg

import (
	"encoding/xml"
	"os"
)

func (h *handler) processSprite(content, action string) error {

	if action == "delete" {
		return os.Remove(h.outFilePath)
	}

	// var sp sprite
	if err := xml.Unmarshal([]byte(content), h.sprite); err != nil {
		return err
	}
	// for _, symbol := range h.sprite.Defs.Symbols {
	// 	h.symbols[symbol.ID] = symbolToString(symbol)
	// }
	return nil
}
