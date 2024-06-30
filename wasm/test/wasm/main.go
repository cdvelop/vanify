//go:build wasm
// +build wasm

package main

import "syscall/js"

var key = ""

func main() {

	js.Global().Get("console").Call("log", "¡Hi Go y WebAssembly!")
	select {}

}
