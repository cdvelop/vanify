package svg

func (h *handler) CreateFileOnDisk(path string, content []byte) error {
	if err := h.processContent(content, "create"); err != nil {
		return err
	}
	return h.writeToFile()
}

func (h *handler) UpdateFileOnDisk(path string, content []byte) error {
	if err := h.processContent(content, "update"); err != nil {
		return err
	}
	return h.writeToFile()
}

func (h *handler) DeleteFileOnDisk(path string, content []byte) error {
	return h.processContent(content, "delete")
}
