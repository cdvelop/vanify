package svg

import (
	"bytes"
	"errors"
)

func (h *handler) Minify(in *bytes.Buffer) error {
	var out bytes.Buffer

	if err := h.min.Minify("image/svg+xml", &out, in); err != nil {
		return errors.New("minify SVG " + err.Error())
	}

	in.Reset()
	in.Write(out.Bytes())

	return nil
}
