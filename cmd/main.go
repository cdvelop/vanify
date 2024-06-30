package main

import (
	"os"

	"github.com/cdvelop/vanify"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	vanify.Run(currentDir)

}
