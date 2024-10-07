package main

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Bipi"))
		err := RunUI(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
