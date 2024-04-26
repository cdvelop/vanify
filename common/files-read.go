package common

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

type out interface {
	LoadMemoryContent(newPath string, content []byte)
}

func ReadAssets(Dir string, Html, Css, Js, Svg out) error {
	const this = "ReadAssets "
	err := filepath.Walk(Dir, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(newPath)

		if ext != ".html" && ext != ".css" && ext != ".js" && ext != ".svg" {
			return nil
		}

		file, er := os.Open(newPath)
		if er != nil {
			return er
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			switch filepath.Ext(newPath) {

			case ".html":
				if Html != nil {
					Html.LoadMemoryContent(newPath, scanner.Bytes())
				}
			case ".css":
				if Css != nil {
					Css.LoadMemoryContent(newPath, scanner.Bytes())
				}
			case ".js":
				if Js != nil {
					Js.LoadMemoryContent(newPath, scanner.Bytes())
				}
			case ".svg":
				if Svg != nil {
					Svg.LoadMemoryContent(newPath, scanner.Bytes())
				}
			}
		}

		if er := scanner.Err(); er != nil {
			return er
		}

		return nil
	})

	if err != nil {
		return errors.New(this + err.Error())
	}

	return nil
}
