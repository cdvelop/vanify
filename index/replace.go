package index

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func ReplaceRoot(html string, newRoot string) string {
	// Compilar la expresión regular
	re := regexp.MustCompile(`<[^>]+>`)

	// Encontrar todas las coincidencias en la cadena de entrada
	items := re.FindAllString(html, -1)
	var out string
	for _, item := range items {
		out += replaceRoot(item, newRoot)
	}

	return out

}

func replaceRoot(html string, newRoot string) string {
	var rex *regexp.Regexp
	var linkType string
	var newPath string

	// fmt.Println("HTML IN:", html)
	// Expresión regular para seleccionar el contenido de href
	if strings.Contains(html, "href=") {
		rex = regexp.MustCompile(`href="([^"]+)"`)
		linkType = "href"

	} else if strings.Contains(html, "src=") {

		rex = regexp.MustCompile(`src="([^"]+)"`)
		linkType = "src"

	} else {
		return html
	}

	// Seleccionar el contenido de href
	match := rex.FindStringSubmatch(html)
	if len(match) > 1 {

		oldHref := match[1]
		// fmt.Println("Antes: ", oldHref)

		// extract the file name
		file := filepath.Base(oldHref)

		// fmt.Println("file: ", file)

		// Nuevo path with / separator
		newPath = newRoot + "/" + file
		// newPath = filepath.Join(newRoot, file)

	} else {
		return html
	}

	// Join the parts back together
	return rex.ReplaceAllString(html, fmt.Sprintf(`%s="%s"`, linkType, newPath))
}
