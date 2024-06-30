package svg

import (
	"errors"
	"strings"
)

// action ej: "create", "modify", "delete"
func (h *handler) processContent(content []byte, action string) error {
	contentStr := string(content)
	if strings.Contains(contentStr, "<svg") {
		return h.processSprite(contentStr, action)
	} else if strings.Contains(contentStr, "<symbol") {
		return h.processSymbol(contentStr, action)
	}
	return errors.New("contenido no reconocido")
}
