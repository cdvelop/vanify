package common

import (
	"log"
	"testing"
)

func TestGetRootFiles(t *testing.T) {
	var testData = map[string]struct {
		file_path                  string
		extension                  string
		onlyNamesWithCapitalLetter bool
		expected                   []string
	}{
		"se espera solo 1 archivo .goo ": {"test", ".goo", true, []string{`test\Inputs1.goo`}},
		"se esperan 2 archivos .goo ":    {"test", ".goo", false, []string{`test\Inputs1.goo`, `test\inputs3.goo`}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := GetRootFiles(data.file_path, data.extension, data.onlyNamesWithCapitalLetter)

			if len(resp) != len(data.expected) {
				log.Fatalf("\n-se esperaba\n%v\n-pero se obtuvo\n%v", data.expected, resp)
			}
			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v\nresp:%v\nerr: %v", data.expected[i], valor, resp, err)
				}
			}

		})
	}
}
