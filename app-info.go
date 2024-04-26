package vanify

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/cdvelop/vanify/git"
)

func (d Vanify) appVersion() string {

	tag, err := git.GitLatestTag()
	if err != "" {
		log.Println("appVersion error", err)
	}

	return tag
}

func (d *Vanify) AppName() string {
	const e = "AppName error "

	current_dir, err := os.Getwd()
	if err != nil {
		log.Println(e, err)
		return "app"
	}

	return filepath.Base(current_dir)

}

func (d Vanify) AppFileName() string {
	const e = "AppFileName error "

	if runtime.GOOS == "windows" {
		return d.AppName() + "-" + d.appVersion() + ".exe"
	}

	return d.AppName() + "-" + d.appVersion()

}
