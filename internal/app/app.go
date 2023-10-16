package app

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk4.git/internal/file"
	"github.com/jeremycruzz/msds301-wk4.git/internal/stats"
)

type app struct {
	file  file.Service
	stats stats.Service
}

// Create creates the app by creating its dependencies
func Create(readFrom, writeTo string) app {
	return CreateApp(file.CreateService(readFrom, writeTo), stats.CreateService())
}

// CreateApp creates the app with already created dependencies NOTE:feels like test damage
func CreateApp(fs file.Service, ss stats.Service) app {
	return app{
		file:  fs,
		stats: ss,
	}
}

func (a app) Run() {

	fmt.Println("Reading file...")
	housingData, err := a.file.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println("Analyzing data...")
	statistics, err := a.stats.Analyze(housingData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing file...")
	err = a.file.Write(statistics)
	if err != nil {
		panic(err)
	}

	fmt.Println("Complete!")
}
