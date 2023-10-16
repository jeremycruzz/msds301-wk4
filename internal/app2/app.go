package app2

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/montanaflynn/stats"
)

var (
	allowNan    = true
	percentiles = []float64{25, 50, 75}
)

type app2 struct{}

func Create() app2 {
	return app2{}
}

func (app2) Run() {
	data := readCsv()
	stats := analyze(data)
	write(stats)
}

func readCsv() blocks {
	// open file
	file, err := os.Open("./data/housesInput.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read from csv
	reader := csv.NewReader(file)

	//throw away header
	reader.Read()

	blocksData := blocks{
		[]float64{},
		[]float64{},
		[]float64{},
		[]float64{},
		[]float64{},
		[]float64{},
		[]float64{},
	}

	for {
		record, err := reader.Read()

		// break on EOF
		if err != nil {
			break
		}

		value, err := strconv.ParseFloat(record[0], 32)
		if err != nil {
			panic(err)
		}

		income, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			panic(err)
		}

		age, err := strconv.ParseFloat(record[2], 32)
		if err != nil {
			panic(err)
		}

		rooms, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			panic(err)
		}

		bedrooms, err := strconv.ParseFloat(record[4], 32)
		if err != nil {
			panic(err)
		}

		population, err := strconv.ParseFloat(record[5], 32)
		if err != nil {
			panic(err)
		}

		household, err := strconv.ParseFloat(record[6], 32)
		if err != nil {
			panic(err)
		}

		blocksData.Values = append(blocksData.Values, value)
		blocksData.Incomes = append(blocksData.Incomes, income)
		blocksData.Ages = append(blocksData.Ages, age)
		blocksData.Rooms = append(blocksData.Rooms, rooms)
		blocksData.Bedrooms = append(blocksData.Bedrooms, bedrooms)
		blocksData.Populations = append(blocksData.Populations, population)
		blocksData.Households = append(blocksData.Households, household)
	}

	return blocksData
}

func analyze(blocks blocks) map[string][]float64 {
	valueDesc, _ := stats.Describe(blocks.Values, allowNan, &percentiles)
	incomeDesc, _ := stats.Describe(blocks.Incomes, allowNan, &percentiles)
	ageDesc, _ := stats.Describe(blocks.Ages, allowNan, &percentiles)
	roomsDesc, _ := stats.Describe(blocks.Rooms, allowNan, &percentiles)
	bedroomsDesc, _ := stats.Describe(blocks.Bedrooms, allowNan, &percentiles)
	populationDesc, _ := stats.Describe(blocks.Populations, allowNan, &percentiles)
	householdDesc, _ := stats.Describe(blocks.Households, allowNan, &percentiles)

	data := map[string][]float64{
		"count": {},
		"mean":  {},
		"std":   {},
		"min":   {},
		"25%":   {},
		"50%":   {},
		"75%":   {},
		"max":   {},
	}
	descriptions := []*stats.Description{valueDesc, incomeDesc, ageDesc, roomsDesc, bedroomsDesc, populationDesc, householdDesc}

	//should append values in order
	for _, desc := range descriptions {
		quarter := 0.0
		median := 0.0
		threeQuarter := 0.0

		for _, dp := range desc.DescriptionPercentiles {
			if dp.Percentile == 25 {
				quarter = dp.Value
			} else if dp.Percentile == 50 {
				median = dp.Value
			} else if dp.Percentile == 75 {
				threeQuarter = dp.Value
			}
		}
		data["count"] = append(data["count"], float64(desc.Count))
		data["mean"] = append(data["mean"], float64(desc.Mean))
		data["std"] = append(data["std"], float64(desc.Std))
		data["min"] = append(data["min"], float64(desc.Min))
		data["25%"] = append(data["25%"], float64(quarter))
		data["50%"] = append(data["50%"], float64(median))
		data["75%"] = append(data["75%"], float64(threeQuarter))
		data["max"] = append(data["max"], float64(desc.Max))
	}

	return data
}

func write(data map[string][]float64) {
	// give permissions (6) to everyone
	file, err := os.OpenFile("./data/benchmark/housesOutputGoFast.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// use tabwriter to format nicely
	writer := tabwriter.NewWriter(file, 1, 0, 0, ' ', tabwriter.AlignRight)

	// header row
	header := "x \tvalue \tincome \tage \trooms \tbedrooms \tpopulation \thh \t\n"
	fmt.Fprint(writer, header)

	// write data to file
	for key, value := range data {
		line := fmt.Sprintf("%s \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t\n",
			key, value[0], value[1], value[2], value[3], value[4], value[5], value[6])
		fmt.Fprint(writer, line)
	}

	// flush at end of writing
	writer.Flush()
}
