package server

import (
	"context"
	"os"
	"os/exec"
	"sync"
)

type Server struct {
	*Config

	flags ldFlags

	bwr browser

	fWatch fileWatch

	appPath string //ej: cmd/app.exe

	cmd *exec.Cmd

	programStartedMessages chan string

	appStarted bool

	interrupt chan os.Signal

	errChan chan error
}

type Config struct {
	RunAppDir    string   // ej: project/cmd
	RunArguments []string //ej: "dev","prod"
	Cache        bool     // default false = "no-cache" true = "cache"
}

type ldFlags interface {
	GetLdFlags() (string, error)
}

type browser interface {
	BrowserStart(wg *sync.WaitGroup)
	BrowserContext() context.Context
}

type fileWatch interface {
	FileWatcherStart(wg *sync.WaitGroup)
}
