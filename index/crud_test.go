package index_test

import (
	"path"
	"testing"

	"github.com/cdvelop/vanify/index"
)

type version struct{}

func (v *version) AppVersion() string {
	return "1.0.0"
}

// se espera que el index.html se compile en el directorio public
// y los links de los activos sean relativos a ese directorio
func TestCase1WorkingOnlyInThemeDirectory(t *testing.T) {
	rootDirectory := "test/project1/website/theme"
	filePath := path.Join(rootDirectory, "index.html")
	// buildDirectory := path.Join(rootDirectory, "build", "public")

	h := index.New(rootDirectory, &version{})

	// simulamos que modificamos el index html del tema
	if err := h.UpdateFileOnDisk(filePath); err != nil {
		t.Fatal(err)
	}

	// eliminamos el directorio construido
	// err := os.RemoveAll(path.Join(rootDirectory, "build"))
	// if err != nil {
	// 	t.Fatal(err)
	// }

}
