package server

import (
	"strconv"
	"strings"
)

func (a *Server) SetDefault() error {

	a.Config.RunAppDir = "website"
	a.Config.RunArguments = []string{"dev"}
	a.Config.Cache = false

	return nil
}

func (a Server) ConfigTemplateContent() string {

	return `# Backend Server config:
runAppDir: ` + a.Config.RunAppDir + `
runArguments: ` + strings.Join(a.Config.RunArguments, ",") + `
cache: ` + strconv.FormatBool(a.Config.Cache)
}
