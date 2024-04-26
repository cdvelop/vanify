package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func FindFirstFileWithExtension(dir, extension string) (string, error) {

	files, err := FileCheck(dir, extension)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if extension == filepath.Ext(file.Name()) {
			file_path := filepath.Join(dir, file.Name())

			content, err := os.ReadFile(file_path)
			if err != nil {
				return "", err
			}
			return string(content), nil

		}

	}

	return "", errors.New("extension " + extension + " no encontrada")
}

func FileCheck(dir string, files_names ...string) ([]fs.DirEntry, error) {

	if dir == "" {
		return nil, errors.New("el parámetro dir no pueden estar vacío")
	}

	if len(files_names) == 0 {
		return nil, errors.New("no hay nombre de archivos para revisar")
	} else {
		for _, pathFile := range files_names {
			if pathFile == "" {
				return nil, errors.New("el parámetro pathFile no pueden estar vacío")
			}
		}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// dir ej: modules/mymodule
// ext ej: .js, .html, .css
func ReadFiles(dir, ext string, out *bytes.Buffer) error {
	const this = "ReadFiles error "
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			file, er := os.Open(path)
			if er != nil {
				return er
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				out.Write(scanner.Bytes())
				out.WriteString("\n")
			}
			if er := scanner.Err(); er != nil {
				return er
			}
		}
		return nil
	})

	if err != nil {
		return errors.New(this + err.Error())
	}

	return nil
}

// pathFile ej: "theme/index.html"
func FileWrite(pathFile string, data bytes.Buffer) error {
	const this = "FileWrite error "
	dst, e := os.Create(pathFile)
	if e != nil {
		return errors.New(this + "al crear archivo " + e.Error())
	}
	defer dst.Close()

	// Copy the uploaded File to the filesystem at the specified destination
	_, e = io.Copy(dst, &data)
	if e != nil {
		return errors.New(this + "no se logro escribir el archivo " + pathFile + " en el destino " + e.Error())
	}

	return nil
}

// ej: gotools.DeleteFilesByExtension(main_folder\files, []string{".js", ".css", ".wasm"})
func DeleteFilesByExtension(dir string, exts []string) (err string) {
	files, e := os.ReadDir(dir)
	if e != nil {
		return e.Error()
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, ext := range exts {
			if ext == filepath.Ext(file.Name()) {
				path := filepath.Join(dir, file.Name())
				e := os.Remove(path)
				if e != nil {
					return "Error deleting file: " + e.Error()
				}
				break
			}
		}
	}

	return ""
}
