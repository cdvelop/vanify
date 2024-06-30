package svg_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/svg"
)

func TestCreateFileOnDisk(t *testing.T) {
	// Configuración inicial
	testPath := "./test"
	h := svg.NewHandler(testPath)

	// Asegurarse de que el directorio de prueba existe
	if err := os.MkdirAll(testPath, 0755); err != nil {
		t.Fatalf("No se pudo crear el directorio de prueba: %v", err)
	}

	// Limpiar después de las pruebas
	defer os.RemoveAll(testPath)

	spriteTest := bytes.NewBuffer([]byte(`<svg xmlns="http://www.w3.org/2000/svg" role="img" aria-hidden="true" focusable="false" class="sprite-icons"><defs></defs></svg>`))
	t.Run("Caso 1 Sprite vacío", func(t *testing.T) {

		err := h.CreateFileOnDisk(testPath, spriteTest.Bytes())
		if err != nil {
			t.Fatalf("Error al crear archivo: %v", err)
		}

		result, err := os.ReadFile(h.OutputFilePath())
		if err != nil {
			t.Fatalf("Error al leer archivo: %v", err)
		}

		err = h.Minify(spriteTest)
		if err != nil {
			t.Fatal(err)
		}

		if string(result) != spriteTest.String() {
			t.Errorf("\nEl contenido del archivo de salida:\n[%v]\n\nNo coincide con el esperado:\n[%v]\n", string(result), spriteTest.String())
		}
	})

	symbolHome := `<symbol id="icon-home" viewBox="0 0 576 512"><path fill="currentColor" d="M280.37 148.26L96" /></symbol>`

	t.Run("Caso 2 Agregar un símbolo", func(t *testing.T) {
		err := h.CreateFileOnDisk(testPath, []byte(symbolHome))
		if err != nil {
			t.Fatalf("Error al crear archivo: %v", err)
		}

		if count := common.TextCount(h.OutputFilePath(), `id="icon-home"`); count != 1 {
			t.Errorf("El contenido del archivo no contiene el símbolo esperado")
		}

	})

	// Caso 3: Repetir el mismo símbolo
	t.Run("Caso 3 Repetir símbolo Home", func(t *testing.T) {
		err := h.CreateFileOnDisk(testPath, []byte(symbolHome))
		if err != nil {
			t.Fatalf("Error al crear archivo: %v", err)
		}

		if count := common.TextCount(h.OutputFilePath(), `id="icon-home"`); count != 1 {
			t.Errorf("Se esperaba 1 símbolo 'icon-home', pero se encontraron %d", count)
		}
	})

	// Caso 4: Agregar un nuevo símbolo
	t.Run("Caso 4 Agregar un nuevo símbolo", func(t *testing.T) {
		symbolPrinter := `<symbol id="icon-printer" viewBox="0 0 16 16"><path fill="currentColor" d="M2" /></symbol>`
		err := h.CreateFileOnDisk(testPath, []byte(symbolPrinter))
		if err != nil {
			t.Fatalf("Error al crear archivo: %v", err)
		}

		result, err := os.ReadFile(h.OutputFilePath())
		if err != nil {
			t.Fatalf("Error al leer archivo: %v", err)
		}

		if !strings.Contains(string(result), `id="icon-home"`) || !strings.Contains(string(result), `id="icon-printer"`) {
			t.Errorf("El contenido del archivo no contiene ambos símbolos esperados")
		}

		homeCount := strings.Count(string(result), `id="icon-home"`)
		printerCount := strings.Count(string(result), `id="icon-printer"`)
		if homeCount != 1 || printerCount != 1 {
			t.Errorf("Se esperaba 1 símbolo de cada tipo, pero se encontraron %d 'icon-home' y %d 'icon-printer'", homeCount, printerCount)
		}
	})

	// Caso 5: Eliminar símbolo Home
	t.Run("Caso 5 Eliminar un símbolo Home", func(t *testing.T) {
		err := h.DeleteFileOnDisk(testPath, []byte(symbolHome))
		if err != nil {
			t.Fatalf("Error al eliminar archivo: %v", err)
		}

		if count := common.TextCount(h.OutputFilePath(), `id="icon-home"`); count != 0 {
			t.Errorf("Se esperaba que el símbolo 'icon-home' fuese eliminado del sprite, pero se encontraron %d", count)
		}
	})

	t.Run("Caso 6 Actualizar símbolo Printer", func(t *testing.T) {
		content := `<symbol id="icon-printer" viewBox="0 0 512 512"><path fill="newUpdateColor" d="M128" /></symbol>`
		err := h.UpdateFileOnDisk(testPath, []byte(content))
		if err != nil {
			t.Fatalf("Error al crear archivo: %v", err)
		}

		if count := common.TextCount(h.OutputFilePath(), `id="icon-printer"`); count != 1 {
			t.Errorf("Al actualizar se esperaba 1 símbolo 'icon-printer', pero se encontraron %d", count)
		}
		if count := common.TextCount(h.OutputFilePath(), `viewBox="0 0 512 512"`); count != 1 {
			t.Error("Al actualizar se esperaba encontrar viewBox=0 0 512 512")
		}
		if count := common.TextCount(h.OutputFilePath(), `newUpdateColor`); count != 1 {
			t.Error("Al actualizar se esperaba encontrar newUpdateColor")
		}
	})

	t.Run("Caso 7 actualizar class sprite-icons a update-sprite", func(t *testing.T) {
		spriteTest2 := bytes.NewBuffer([]byte(`<svg xmlns="http://www.w3.org/2000/svg" class="update-sprite" role="img" aria-hidden="true" focusable="false" ><defs></defs></svg>`))

		err := h.UpdateFileOnDisk(testPath, spriteTest2.Bytes())
		if err != nil {
			t.Fatalf("Error al eliminar sprite: %v", err)
		}

		if count := common.TextCount(h.OutputFilePath(), `class="update-sprite"`); count != 1 {
			t.Errorf("Al actualizar el sprite se esperaba la clase 'update-sprite', pero no se encontró %d", count)
		}
	})

	t.Run("Caso 8 Eliminar sprite", func(t *testing.T) {
		err := h.DeleteFileOnDisk(testPath, spriteTest.Bytes())
		if err != nil {
			t.Fatalf("Error al eliminar sprite: %v", err)
		}

		_, err = os.ReadFile(h.OutputFilePath())
		if err == nil {
			t.Fatal("se esperaba error al leer archivo ya que se elimino")
		}

	})
}
