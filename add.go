package vanify

import (
	"os"
	"os/exec"

	"strings"

	"github.com/cdvelop/vanify/browser"
	"github.com/cdvelop/vanify/compiler"
	"github.com/cdvelop/vanify/token"
	"github.com/cdvelop/vanify/watch"
)

func Add() *Vanify {
	d := &Vanify{
		app_path:               "app.exe",
		Browser:                browser.Add(),
		WatchFiles:             &watch.WatchFiles{},
		Cmd:                    &exec.Cmd{},
		Interrupt:              make(chan os.Signal, 1),
		ProgramStartedMessages: make(chan string),
		TwoKeys:                &token.TwoKeys{},
	}

	d.app_path = d.AppFileName()

	var e2eWasmTestSuffix string

	for _, v := range os.Args {

		if strings.Contains(v, "test:") {
			d.run_arguments = append(d.run_arguments, v)
			e2eWasmTestSuffix = "test:"
		}
		if v == "dev" {
			d.run_arguments = append(d.run_arguments, v)
		}

	}

	d.Compiler = compiler.Add(&compiler.Config{
		AppInfo:             d,
		TwoPublicKeyAdapter: d.TwoKeys,
	}, e2eWasmTestSuffix, "compile_dir:cmd")

	d.run_arguments = append(d.run_arguments, Cache)

	d.WatchFiles = watch.Add(d, d, d, d.DirectoriesRegistered, d.Compiler.ThemeDir())

	return d
}
