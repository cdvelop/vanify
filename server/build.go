package server

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/cdvelop/vanify/common"
)

func (a *Server) goBuildAndExec() error {
	e := errors.New("goBuildAndExec: ")

	common.PrintWarning(fmt.Sprintf("Building and Running %s...\n", a.appPath))

	os.Remove(a.appPath)

	flags, err := a.flags.GetLdFlags()
	if err != nil {
		return errors.Join(err, e)
	}

	// var ldflags = `-X 'main.version=` + tag + `'`

	a.cmd = exec.Command("go", "build", "-o", a.appPath, "-ldflags", flags, "main.go")
	// a.cmd = exec.Command("go", "build", "-o", d.appPath, "main.go" )

	stderr, err := a.cmd.StderrPipe()
	if err != nil {
		return errors.Join(e, err)
	}

	stdout, er := a.cmd.StdoutPipe()
	if er != nil {
		return errors.Join(e, err)
	}

	er = a.cmd.Start()
	if er != nil {
		return errors.Join(e, err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := io.ReadAll(stderr)

	// Esperar
	er = a.cmd.Wait()
	if er != nil {
		return errors.Join(e, errors.New(string(errBuf)), err)
	}

	return a.cmdExec()
}
