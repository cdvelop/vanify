package server

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/cdvelop/vanify/common"
	"github.com/cdvelop/vanify/pkg/chromedp"
)

func (a *Server) StartProgram() {

	e := "StartProgram: "

	go func() {

		// para ejecutar moverse  al directorio "RunAppDir" si existe

		if _, err := os.Stat(a.RunAppDir); err == nil {
			er := os.Chdir(a.RunAppDir)
			if er != nil {
				a.errChan <- fmt.Errorf("%s error al cambiar al directorio '%s': %s", e, a.RunAppDir, er.Error())
				return
			}
		}

		// BUILD AND RUN
		err := a.goBuildAndExec()
		if err != nil {
			a.errChan <- fmt.Errorf("%s %s", e, err)
		}

	}()

	for {

		select {

		case message := <-a.programStartedMessages:

			if strings.Contains(strings.ToLower(message), "err") {
				common.PrintError(message)
			} else {
				common.PrintOK(message)

				if !a.appStarted {

					var wg sync.WaitGroup
					wg.Add(2)

					go a.bwr.BrowserStart(&wg)

					go a.fWatch.FileWatcherStart(&wg)

					a.appStarted = true
				}
			}

		case <-a.interrupt:
			// Detenga el navegador y cierre la aplicación cuando se recibe una señal de interrupción
			if er := chromedp.Cancel(a.bwr.BrowserContext()); er != nil {
				common.PrintError("al cerrar browser " + er.Error())
			}
			err := a.StopProgram()
			if err != nil {
				common.PrintError("al detener app: " + err.Error())
			}

			os.Exit(0)

		case err := <-a.errChan:
			common.PrintError(err.Error())
		}

	}

}
