package vanify_test

import (
	"testing"

	"github.com/cdvelop/vanify"
)

func TestStartProgram(t *testing.T) {

	err := vanify.Run("test")
	if err != nil {
		t.Fatal(err)
	}
}
