package main

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app2"
)

func main() {
	// default values
	n := 100

	fmt.Println("Running fast")
	fmt.Println(n)

	app := app2.Create()

	for i := 0; i < n; i++ {
		app.Run()
	}
	fmt.Println("complete!")
}
