package index

import (
	"bytes"
	"path"
	"path/filepath"
	"strings"

	"github.com/cdvelop/vanify/pkg/minify"
	min "github.com/cdvelop/vanify/pkg/minify/html"
)

// rootDirectory for test ej "test/project1" default is current directory
func New(rootDirectory string, v appVersion) *handler {

	m := minify.New()
	m.AddFunc("text/html", min.Minify)

	h := handler{
		themeDir: path.Join(rootDirectory, "website", "theme"),
		FilePath: filepath.Join(rootDirectory, "build", "public", "index.html"),
		ver:      v,

		registeredFiles: map[string]*file{},
		buf:             bytes.Buffer{},

		min: m,
	}

	dirs := strings.Split(rootDirectory, "website")
	if len(dirs) == 2 && dirs[1] == "/theme" {
		h.themeDir = rootDirectory
	}

	return &h

}
