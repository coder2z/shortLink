package main

import (
	"fmt"
	"os"
	"shortLink/cmd/master/app"
	"shortLink/pkg/signals"
)

func main() {
	err := app.Run(signals.SetupSignalHandler())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
