package common_test

import (
	"testing"

	"github.com/cdvelop/vanify/common"
)

func TestExist(t *testing.T) {
	var testData = map[string]struct {
		path   string
		search string
		expect int
	}{
		"archivo file.txt en carpeta test":                         {"test\\file.txt", "contenido", 1},
		"archivo file.txt contra archivo":                          {"test\\file.txt", "test\\file.txt", 1},
		"archivo style.css minificado":                             {"test\\style.css", "search-test-style{background:#ff0}", 1},
		"agregado al inicio:punto, espacio adicional y Mayúscula ": {"test\\style.css", " .Search-Test-style{background:#ff0}", 1},
		"clase a buscar sin minificacion": {"test\\style.css", `.miClaseABuscar{
			background-color: #ff0;
		}`, 1},
		"sin data se espera 0": {"", "", 0},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp := common.TextCount(data.path, data.search)

			if resp != data.expect {
				t.Fatalf("Para la entrada '%s', '%s' se esperaba '%v' pero se obtuvo '%v'", data.path, data.search, data.expect, resp)
			}
		})
	}
}
