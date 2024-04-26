package vanify

import (
	"fmt"
	"os"

	. "github.com/cdvelop/vanify/common"
)

// HomePath: ej: /index.html, / ,/home

func (v Vanify) Help() {

	for _, opt := range os.Args {

		switch {

		case opt == "help" || opt == "?" || opt == "ayuda":

			fmt.Println(`
			

			
			
			`)

			ShowErrorAndExit("")
		}

	}

}
