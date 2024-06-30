package server

import (
	"errors"
	"io"
	"os/exec"
)

// Construir el comando con argumentos dinámicos
// cmdArgs := append([]string{"go", "build", "-o", d.appPath, "main.go"}, os.Args...)
// d.cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)

func (d *Server) cmdExec() error {
	e := errors.New("app run: ")

	d.cmd = exec.Command("./"+d.appPath, d.RunArguments...)

	stderr, err := d.cmd.StderrPipe()
	if err != nil {
		return errors.Join(e, err)
	}

	stdout, err := d.cmd.StdoutPipe()
	if err != nil {
		return errors.Join(e, err)
	}

	err = d.cmd.Start()
	if err != nil {
		return errors.Join(e, err)
	}

	go io.Copy(d, stderr)
	go io.Copy(d, stdout)

	return nil
}
