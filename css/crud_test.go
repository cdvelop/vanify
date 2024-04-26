package css_test

import (
	"bytes"
	"testing"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/css"
)

func TestCrud(t *testing.T) {

	publicFolder := "test/public/static"
	h := css.New(publicFolder)

	var content bytes.Buffer
	if _, err := content.Write([]byte(`body {
		color: red;
	}`)); err != nil {
		t.Fatal(err)
	}

	filePath := "test/css/new.css"
	if err := common.FileWrite(filePath, content); err != nil {
		t.Fatal(err)
	}

	h.LoadMemoryContent(filePath, content.Bytes())

	// archivo modificado en disco
	if err := h.FileWasModifiedOnDisk(filePath); err != nil {
		t.Fatal(err)
	}

	// minificar antes de comparar
	if err := h.CssMinify(&content); err != nil {
		t.Fatal(err)
	}

	// fmt.Println("CONTENIDO", content.String())
	// verificar si el contenido se agrego
	if resp := common.TextExists(h.CssFilePath, content.String()); resp == 0 {
		t.Fatal("EN: "+h.CssFilePath+" NO EXISTE EL CONTENIDO: ", content.String())
	}

	// actualizamos contenido
	content.Reset()
	content.WriteString(`body { color: green; }`)

	if err := common.FileWrite(filePath, content); err != nil {
		t.Fatal(err)
	}

	if err := h.FileWasModifiedOnDisk(filePath); err != nil {
		t.Fatal(err)
	}

	// minificar antes de comparar
	if err := h.CssMinify(&content); err != nil {
		t.Fatal(err)
	}

	// verificar si el nuevo contenido se agrego
	resp := common.TextExists(h.CssFilePath, content.String())
	if resp == 0 {
		t.Fatal("EN: "+h.CssFilePath+" NO EXISTE EL CONTENIDO: ", content.String())
	}

	if resp > 1 {
		t.Fatal("EN: "+h.CssFilePath+" ESTA REPETIDO: ", content.String())
	}

	// fmt.Println("respuesta", resp)

}
