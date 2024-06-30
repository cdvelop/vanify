package server

import (
	"errors"
	"fmt"
)

func (d *Server) Restart(event_name string) error {
	e := errors.New("Restart Server: ")
	fmt.Println("Reiniciando APP..." + event_name)

	// STOP
	err := d.StopProgram()
	if err != nil {
		return errors.Join(e, err)
	}

	// BUILD AND RUN
	err = d.goBuildAndExec()
	if err != nil {
		return errors.Join(e, err)
	}

	return nil
}
