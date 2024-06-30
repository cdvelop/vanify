package server

import (
	"errors"
	"fmt"

	"github.com/cdvelop/vanify/common"
)

func (d *Server) StopProgram() error {
	e := errors.New("StopProgram: ")
	pid := d.cmd.Process.Pid

	common.PrintWarning(fmt.Sprintf("stop app PID %d\n", pid))

	err := d.cmd.Process.Kill()
	if err != nil {
		return errors.Join(e, err)
	}

	return nil
}
