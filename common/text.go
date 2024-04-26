package common

import (
	"os"
	"strings"
)

// ej: index.html", "info_id"
// static\main.js", "js\code.js"
func TextExists(filePath string, textSearchInContent string) int {
	file_content, err := os.ReadFile(filePath)
	if err != nil {
		return 0
	}
	string_content := string(file_content)
	TextMinify(&string_content)

	// Variable para almacenar el contenido de búsqueda
	var text_search string

	for _, value := range []string{".", "\\", "/"} {
		// Puede que sea una ruta
		if strings.Contains(textSearchInContent, value) {
			search_content, err := os.ReadFile(textSearchInContent)
			if err == nil {
				text_search = string(search_content)
			}
			break
		}
	}

	if text_search == "" {
		text_search = textSearchInContent
	}

	TextMinify(&text_search)

	// fmt.Printf("CONTENIDO: [%v]\n", string_content)
	// fmt.Printf("TEXTO A BUSCAR: [%v]\n", text_search)

	return strings.Count(string_content, text_search)
}

func TextMinify(text *string) {

	if text != nil {

		// Eliminar espacios en blanco al inicio y al final
		*text = strings.TrimSpace(*text)

		// Eliminar espacios en blanco entre caracteres y saltos de línea
		*text = strings.Join(strings.Fields(*text), "")

		// Convertir a minúsculas
		*text = strings.ToLower(*text)

	}

}
