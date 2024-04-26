package common

import (
	"fmt"
	"os"
)

var format_color string

func init() {
	for _, arg := range os.Args {
		if arg == "dev" {
			format_color = "\033[%sm%s\033[0m"
		}
	}
}

func ShowErrorAndExit(message any) {
	PrintError(fmt.Sprintf("Error: %v\n", message))
	fmt.Println("")
	PrintInfo("Presiona enter para salir...\n")
	fmt.Scanln()
	os.Exit(1)
}

// options: ok,err,warn,info
func print(message string, options ...string) {
	if format_color == "" {
		fmt.Print(message)
	} else {

		var color string

		for _, opt := range options {
			switch opt {
			case "ok":
				color = "32" //green
			case "err":
				color = "31" //red
			case "warn":
				color = "33" //yellow
			case "info":
				color = "36" //magenta blue=34
			default:
				color = "0"
			}
		}

		fmt.Printf(format_color, color, message)
	}
}

func PrintOK(messages ...string) {
	print(joinMessages(messages...), "ok")
}

func PrintWarning(messages ...string) {

	print(joinMessages(messages...), "warn")
}
func PrintError(messages ...string) {
	print(joinMessages(messages...), "err")
}

func PrintInfo(messages ...string) {
	print(joinMessages(messages...), "info")
}

func joinMessages(messages ...string) (message string) {
	var space string
	for _, m := range messages {
		message += space + m
		space = " "
	}
	return
}
