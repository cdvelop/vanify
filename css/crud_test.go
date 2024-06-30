package css_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/css"
)

func TestCrud(t *testing.T) {

	publicFolder := "test/public/static"
	h := css.New(publicFolder)

	content := bytes.Buffer{}
	if _, err := content.Write([]byte(`body {
		color: red;
	}`)); err != nil {
		t.Fatal(err)
	}

	// escribimos el archivo fuente en disco
	filePath := "test/css/new.css"
	if err := common.FileWrite(filePath, content); err != nil {
		t.Fatal(err)
	}

	h.UpdateFileContentInMemory(filePath, content.Bytes())

	// archivo modificado en disco
	if err := h.UpdateFileOnDisk(filePath); err != nil {
		t.Fatal(err)
	}

	// minificar antes de comparar
	if err := h.Minify(&content); err != nil {
		t.Fatal(err)
	}

	// fmt.Println("CONTENIDO", content.String())
	// verificar si el contenido se agrego
	if resp := strings.Count(h.FilePath, content.String()); resp == 0 {
		t.Fatal("EN: "+h.FilePath+" NO EXISTE EL CONTENIDO: ", content.String())
	}

	// // actualizamos contenido
	content.Reset()
	content.WriteString(`body { color: green; }`)

	if err := common.FileWrite(filePath, content); err != nil {
		t.Fatal(err)
	}

	if err := h.UpdateFileOnDisk(filePath); err != nil {
		t.Fatal(err)
	}

	// // minificar antes de comparar
	if err := h.Minify(&content); err != nil {
		t.Fatal(err)
	}
	fmt.Println("contenido ok:", content.String())

	// // verificar si el nuevo contenido se agrego
	resp := strings.Count(h.FilePath, content.String())
	if resp == 0 {
		t.Fatal("EN: "+h.FilePath+" NO EXISTE EL CONTENIDO: ", content.String())
	}

	if resp > 1 {
		t.Fatal("EN: "+h.FilePath+" ESTA REPETIDO: ", content.String())
	}

}
