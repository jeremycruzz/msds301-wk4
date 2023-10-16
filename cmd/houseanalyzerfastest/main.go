package main

import (
	"github.com/jeremycruzz/msds301-wk4.git/internal/app4"
)

func main() {
	// default values
	n := 100

	app := app4.Create()

	for i := 0; i < n; i++ {
		app.Run()
	}
}
