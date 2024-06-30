package wasm_test

import (
	"testing"

	"github.com/cdvelop/vanify/wasm"
)

func TestFileWasModifiedOnDisk(t *testing.T) {

	// Create handler with test config
	h := wasm.New("test", nil)

	err := h.UpdateFileOnDisk(h.MainGoFilePathToBuildWasm)
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(h)

}
