package svg

import (
	"encoding/xml"
)

func (h *handler) processSymbol(content, action string) error {
	var sym symbol
	if err := xml.Unmarshal([]byte(content), &sym); err != nil {
		return err
	}

	if action == "delete" {
		delete(h.icons, sym.ID)
		h.writeToFile()
		return nil
	}

	h.icons[sym.ID] = icon{
		content: content,
		symbol:  sym,
	}

	// fmt.Println("simbolo", sym.ID, sym.ViewBox)

	return nil
}

// func symbolToString(symbol symbol) string {
// 	return strings.TrimSpace(strings.Replace(strings.Replace(symbol.Path.D, "\n", "", -1), "\t", "", -1))
// }
