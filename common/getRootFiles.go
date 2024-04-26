package common

import (
	"path/filepath"
	"unicode"
)

func GetRootFiles(dir, ext string, onlyNamesWithCapitalLetter bool) (out []string, err error) {

	files, err := FileCheck(dir, ext)
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		// omitir si es directorio
		if file.IsDir() {
			continue
		}

		// omitir si comienza con minúscula
		if onlyNamesWithCapitalLetter && !unicode.IsUpper([]rune(file.Name())[0]) {
			continue
		}

		// omitir si no tiene extensión
		if ext != filepath.Ext(file.Name()) {
			continue
		}

		out = append(out, filepath.Join(dir, file.Name()))

	}

	return
}
