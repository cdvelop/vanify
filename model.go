package vanify

import (
	"os"
	"os/exec"

	"github.com/cdvelop/vanify/browser"
	"github.com/cdvelop/vanify/compiler"
	"github.com/cdvelop/vanify/token"
	"github.com/cdvelop/vanify/watch"
)

type Vanify struct {
	app_path string //ej: app.exe

	*browser.Browser
	*watch.WatchFiles
	*compiler.Compiler

	*exec.Cmd

	// Scanner   *bufio.Scanner
	Interrupt chan os.Signal

	ProgramStartedMessages chan string

	run_arguments []string

	*token.TwoKeys
}
