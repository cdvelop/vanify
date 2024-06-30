package server

import (
	"os"
	"os/exec"
)

func New(c *Config, f ldFlags, b browser, w fileWatch) *Server {

	a := &Server{
		Config:                 c,
		flags:                  f,
		bwr:                    b,
		fWatch:                 w,
		appPath:                "",
		cmd:                    &exec.Cmd{},
		programStartedMessages: make(chan string),
		appStarted:             false,
		interrupt:              make(chan os.Signal, 1),
		errChan:                make(chan error),
	}

	path, err := a.AppFileName()
	if err != nil {
		panic(err)
	}
	a.appPath = path

	a.SetDefault()

	return a

}

func (a Server) SignalNotify() chan os.Signal {
	return a.interrupt
}
