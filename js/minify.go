package js

import (
	"bytes"
	"errors"
)

func (h *handler) Minify(in *bytes.Buffer) error {
	var out bytes.Buffer

	if err := h.min.Minify("text/javascript", &out, in); err != nil {
		return errors.New("minify JS " + err.Error())
	}

	in.Reset()
	in.Write(out.Bytes())

	return nil
}
