package vanify

import (
	"errors"
	"path/filepath"
)

func Run(currentDir string) error {

	if filepath.Base(currentDir) == "vanify" {
		return errors.New("cambia al directorio de tu aplicación antes de ejecutar vanify")
	}

	// b := browser.

	// canal para recibir señales de interrupción
	// signal.Notify(a.SignalNotify(), os.Interrupt, syscall.SIGTERM)

	// fmt.Println("app conf:", a)

	return nil
}
