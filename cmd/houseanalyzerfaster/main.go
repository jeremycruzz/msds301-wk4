package main

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app3"
)

func main() {
	// default values
	n := 100

	fmt.Println("Running faster")
	fmt.Println(n)

	app := app3.Create()

	for i := 0; i < n; i++ {
		app.Run()
	}
	fmt.Println("complete!")
}
