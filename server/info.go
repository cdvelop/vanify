package server

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/cdvelop/vanify/common"
)

func (d Server) appVersion() (string, error) {

	tag, err := common.ExecCommand("git", "describe", "--abbrev=0", "--tags")
	if err != nil {
		return "", errors.New("appVersion error: " + err.Error())
	}

	return tag, nil
}

func (d *Server) AppName() (string, error) {

	current_dir, err := os.Getwd()
	if err != nil {
		return "", errors.New("AppName error: " + err.Error())
	}

	return filepath.Base(current_dir), nil

}

func (a Server) AppFileName() (string, error) {
	const e = "AppFileName error: "

	name, err := a.AppName()
	if err != nil {
		return "", errors.New(e + err.Error())
	}

	ver, err := a.appVersion()
	if err != nil {
		return "", errors.New(e + err.Error())
	}

	var ext string

	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	return name + "-" + ver + ext, nil

}
