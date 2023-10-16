package main

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app4"
)

func main() {
	// default values
	n := 100

	fmt.Println("Running fastest")
	fmt.Println(n)

	app := app4.Create()

	for i := 0; i < n; i++ {
		app.Run()
	}
	fmt.Println("complete!")
}
