package css

import (
	"errors"
)

func (h *handler) Remove(path string) error {
	const this = "Remove Css "

	if file, exist := h.registeredFiles[path]; exist {
		h.files[file.index].content = nil

		delete(h.registeredFiles, path)

		h.files = append(h.files[:file.index], h.files[file.index+1:]...)

	} else {
		return errors.New(this + "file not found")
	}

	return nil
}
