package wasm

type handler struct {
	rootDirectory             string
	tinyGoCompiler            bool `yaml:"tinyGoCompiler"`
	buildDirectory            string
	MainGoFilePathToBuildWasm string // ej: ./project/wasm/main.go
	WasmFileOutPath           string // ej: ./project/cmd/public/static/main.wasm
	wasmProjectType           bool

	wasmFileName       string
	wasmExecJsPathFile string
	wasmImportJS       []byte

	gomodExist bool

	flags
}

type flags interface {
	Flags() string
}
