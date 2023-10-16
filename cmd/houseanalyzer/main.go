package main

import (
	"fmt"
	"os"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app"
)

func main() {
	// default values
	readFrom := "./data/housesInput.csv"
	writeTo := "./data/housesOutputGo.txt"
	n := 100

	// get args
	args := os.Args[1:]

	if len(args) >= 1 {
		readFrom = args[0]
	}

	if len(args) >= 2 {
		writeTo = args[1]
	}

	fmt.Println(n)
	fmt.Println("Reading data from:\t\t" + readFrom)
	fmt.Println("Writing data to:\t\t" + writeTo)

	app := app.Create(readFrom, writeTo)

	for i := 0; i < n; i++ {
		app.Run()
	}
	fmt.Println("complete!")
}
