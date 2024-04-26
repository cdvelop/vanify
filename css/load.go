package css

func (h *handler) LoadMemoryContent(filePath string, content []byte) {

	if fileFound, exist := h.registeredFiles[filePath]; exist {

		fileFound.content = content

	} else {

		new := &file{
			index:   len(h.files),
			path:    filePath,
			content: content,
		}

		h.registeredFiles[filePath] = new

		h.files = append(h.files, new)
	}

}
