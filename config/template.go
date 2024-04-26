package config

import (
	"errors"
	"os"
)

func (c Config) WriteTemplateFile() error {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		tplFile, err := os.Create(FileName)
		if err != nil {
			return errors.New("Failed to create template file: " + err.Error())
		}
		defer tplFile.Close()

		content := c.Browser.ConfigTemplateContent() + `
		` + c.Compiler.ConfigTemplateContent()

		_, err = tplFile.WriteString(content)
		if err != nil {
			return errors.New("Failed to write template file: " + err.Error())

		}
	}

	return nil
}
