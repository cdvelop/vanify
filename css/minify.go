package css

import (
	"bytes"
	"errors"
)

func (h *handler) Minify(in *bytes.Buffer) error {
	var out bytes.Buffer

	if err := h.min.Minify("text/css", &out, in); err != nil {
		return errors.New("Minify CSS " + err.Error())
	}

	in.Reset()
	in.Write(out.Bytes())

	return nil
}
