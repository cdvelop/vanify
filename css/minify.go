package css

import (
	"bytes"
	"errors"
)

// func (h *handler) CssMinify() error {

// 	var out bytes.Buffer

// 	if err := h.minify(&h.buf, &out); err != nil {
// 		return err
// 	}

// 	h.buf.Reset()
// 	h.buf = out

// 	return nil
// }

// func (h *handler) MinifyForTest(in *bytes.Buffer) error {

// 	var out bytes.Buffer

// 	if err := h.minify(in, &out); err != nil {
// 		return err
// 	}

// 	*content = out.Bytes()

// 	return nil
// }

func (h *handler) CssMinify(in *bytes.Buffer) error {
	var out bytes.Buffer

	if err := h.min.Minify("text/css", &out, in); err != nil {
		return errors.New("minify CSS " + err.Error())
	}

	in.Reset()
	in.Write(out.Bytes())

	return nil
}
