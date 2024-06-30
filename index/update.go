package index

import (
	"bytes"
	"fmt"
	"strings"
)

func (h *handler) UpdateFileContentInMemory(filePath string, content []byte) {
	h.buf.Reset()

	fmt.Println("filePath in:", filePath)
	fmt.Println("h.themeDir:", h.themeDir)
	// filePath es modulo o es el archivo index.html?
	if strings.Contains(filePath, h.themeDir) {
		fmt.Println("filePath is index.html rebuilding index.html")

		stringContent := string(content)
		for _, line := range strings.Split(stringContent, "\n") {

			switch {
			// cambiar la ruta del icono de la app
			case strings.Contains(line, `link rel="icon" `):
				line = ReplaceRoot(line, "assets")

				// eliminar todos los link stylesheets
			case strings.Contains(line, `link rel="stylesheet"`):
				line = ""

			case strings.Contains(line, `</head>`):
				fmt.Println("</head> end line:", line)

			case strings.Contains(line, `<body>`):
				fmt.Println("body line:", line)

			}

			// fmt.Println("line:", line)
			if line != "" {
				h.buf.WriteString(line + "\n")
			}
		}

	} else {
		fmt.Println("filePath is module")
	}

}

func (h *handler) makeHtmlTemplate() (html bytes.Buffer, err error) {
	// const e = "makeHtmlTemplate "
	// data, err := os.ReadFile(filepath.Join(h.themeDir, "index.html"))
	// if err != nil {
	// 	err = errors.New(e + "THEME FOLDER: " + h.themeDir + " " + err.Error())
	// 	return
	// }
	// t, err := template.New("").Parse(string(data))
	// if err != nil {
	// 	err = errors.New(e + err.Error())
	// 	return
	// }

	// // template.HTMLEscapeString()
	// if h.JsonBootActions == "" {
	// 	h.JsonBootActions = `{{.JsonBootActions}}`
	// }

	// h.AppName = h.Config.AppName()
	// h.appVersion = h.Config.appVersion()
	// h.StyleSheet = "static/style.css" + h.versionStatics()
	// h.Script = "static/main.js" + h.versionStatics()

	// er = t.Execute(&html, h.Page)
	// if er != nil {
	// 	err = errors.New(e + er.Error())
	// }

	return
}
