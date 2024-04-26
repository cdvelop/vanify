package config

import (
	"github.com/cdvelop/vanify/browser"
	"github.com/cdvelop/vanify/compiler"
)

const FileName = "vanify.yml"

type Config struct {
	browser.Browser   `yaml:",inline"` // yaml:",inline" is necessary
	compiler.Compiler `yaml:",inline"`
}
