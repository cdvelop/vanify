package common

import (
	"errors"
	"fmt"
	"os"
)

func CheckDirectories(directories ...string) error {

	for _, dir := range directories {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return errors.New("Error creando directorio " + dir)
			}
			fmt.Printf("Directorio %s creado correctamente.\n", dir)
		} else if err != nil {
			return errors.New("Error al verificar directorio " + dir + " " + err.Error())
		}
	}
	return nil
}
