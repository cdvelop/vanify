package svg

import (
	"github.com/cdvelop/vanify/pkg/minify"
	min "github.com/cdvelop/vanify/pkg/minify/svg"
)

// outFilePath ej "project/website/build/assets"
func NewHandler(outFilePath string) *handler {

	m := minify.New()
	m.AddFunc("image/svg+xml", min.Minify)

	return &handler{
		outFilePath: outFilePath + "/sprite.svg",
		sprite:      &sprite{},
		icons:       make(map[string]icon),
		min:         m,
	}
}
