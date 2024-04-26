package vanify

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	. "github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/ldflags"
)

func (d *Vanify) buildAndRun() (err string) {
	const this = "buildAndRun error "
	PrintWarning(fmt.Sprintf("Building and Running %s...\n", d.app_path))

	os.Remove(d.app_path)

	flags, err := ldflags.Add(
		d.TwoKeys.GetTwoPublicKeysWasmClientAndGoServer(),
	// sessionbackend.AddPrivateSecretKeySigning(),
	)
	if err != "" {
		return this + err
	}

	// var ldflags = `-X 'main.version=` + tag + `'`

	d.Cmd = exec.Command("go", "build", "-o", d.app_path, "-ldflags", flags, "main.go")
	// d.Cmd = exec.Command("go", "build", "-o", d.app_path, "main.go" )

	stderr, er := d.Cmd.StderrPipe()
	if er != nil {
		return this + er.Error()
	}

	stdout, er := d.Cmd.StdoutPipe()
	if er != nil {
		return this + er.Error()
	}

	er = d.Cmd.Start()
	if er != nil {
		return this + er.Error()
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := io.ReadAll(stderr)

	// Esperar
	er = d.Cmd.Wait()
	if er != nil {
		return this + string(errBuf) + " " + er.Error()
	}

	return d.run()
}

// Construir el comando con argumentos dinámicos
// cmdArgs := append([]string{"go", "build", "-o", d.app_path, "main.go"}, os.Args...)
// d.Cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)

func (d *Vanify) run() (err string) {
	const this = "run error "

	d.Cmd = exec.Command("./"+d.app_path, d.run_arguments...)

	stderr, er := d.Cmd.StderrPipe()
	if er != nil {
		return this + er.Error()
	}

	stdout, er := d.Cmd.StdoutPipe()
	if er != nil {
		return this + er.Error()
	}

	er = d.Cmd.Start()
	if er != nil {
		return this + er.Error()
	}

	go io.Copy(d, stderr)
	go io.Copy(d, stdout)

	return ""
}

func (d Vanify) Write(p []byte) (n int, err error) {
	d.ProgramStartedMessages <- string(p)
	// fmt.Println(string(p))
	return len(p), nil
}

func (d *Vanify) StopProgram() (err string) {

	pid := d.Cmd.Process.Pid

	PrintWarning(fmt.Sprintf("stop app PID %d\n", pid))

	er := d.Cmd.Process.Kill()
	if er != nil {
		err = er.Error()
	}

	return
}
