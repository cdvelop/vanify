package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
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
// data ej: *bytes.Buffer
// NOTA: la data del buf sera eliminada después de escribir el archivo
func FileWrite(pathFile string, data bytes.Buffer) error {
	const e = "FileWrite "
	dst, err := os.Create(pathFile)
	if err != nil {
		if strings.Contains(err.Error(), "system cannot find the path") {
			dir := filepath.Dir(pathFile)
			os.MkdirAll(dir, 0777)
			dst, err = os.Create(pathFile)
			if err != nil {
				return errors.New(e + "al crear archivo " + err.Error())
			}
		} else {
			return errors.New(e + "al crear archivo " + err.Error())
		}

	}
	defer dst.Close()

	// fmt.Println("data antes de escribir:", data.String())
	// Copy the uploaded File to the filesystem at the specified destination
	// _, e = io.Copy(dst, bytes.NewReader(data.Bytes()))
	_, err = io.Copy(dst, &data)
	if err != nil {
		return errors.New(e + "no se logro escribir el archivo " + pathFile + " en el destino " + err.Error())
	}
	// fmt.Println("data después de copy:", data.String(), "bytes:", data.Bytes())

	return nil
}

// ej: gotools.DeleteFilesByExtension(main_folder\files, []string{".js", ".css", ".wasm"})
func DeleteFilesByExtension(dir string, exts []string) error {
	const e = "DeleteFilesByExtension "

	files, err := os.ReadDir(dir)
	if err != nil {

		if strings.Contains(err.Error(), "cannot find the path specified") {
			return nil
		}

		return errors.New(e + err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, ext := range exts {
			if ext == filepath.Ext(file.Name()) {
				path := filepath.Join(dir, file.Name())
				err := os.Remove(path)
				if err != nil {
					return errors.New(e + err.Error())
				}
				break
			}
		}
	}

	return nil
}

func CopyFile(src string, dest string) error {
	const e = "CopyFile "
	// Abrir el archivo origen
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.New(e + err.Error())
	}
	defer srcFile.Close()

	// Crear el archivo destino
	destFile, err := os.Create(dest)
	if err != nil {
		return errors.New(e + err.Error())
	}
	defer destFile.Close()

	// Copiar el contenido del archivo origen al archivo destino
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.New(e + err.Error())
	}

	return nil
}

func FindFilesWithNonZeroSize(dir string, filenames []string) error {
	const e = "FindFilesWithNonZeroSize "
	// Esperar
	time.Sleep(50 * time.Millisecond)

	// Crea un mapa para hacer un seguimiento de los archivos encontrados
	found := make(map[string]bool)
	for _, filename := range filenames {
		found[filename] = false
	}

	// Recorre el directorio en busca de archivos
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Comprueba si el archivo actual es uno de los que estamos buscando
		filename := filepath.Base(path)
		if _, ok := found[filename]; ok && info.Size() > 0 {
			found[filename] = true
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Verifica que se encontraron todos los archivos y que tienen tamaño mayor que cero
	for filename, ok := range found {
		if !ok {
			return errors.New(e + "no se encontró el archivo " + filename)
		}

	}

	return nil
}
