package index

import (
	"bytes"
	"errors"
)

func (h *handler) Minify(in *bytes.Buffer) error {

	var out bytes.Buffer

	if err := h.min.Minify("text/html", &out, in); err != nil {
		return errors.New("Minify HTML " + err.Error())
	}

	in.Reset()
	in.Write(out.Bytes())

	return nil
}
