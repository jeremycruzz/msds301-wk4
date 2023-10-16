package file

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app1/common"
)

func (s service) Read() ([]common.Block, error) {

	// open file
	file, err := os.Open(s.readFrom)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read from csv
	reader := csv.NewReader(file)
	row, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var blocks []common.Block

	// skip header row
	if len(row) > 0 {
		row = row[1:]
	}

	// convert data into struct
	for _, data := range row {
		if len(data) != 7 {
			return nil, errors.New("invalid CSV format")
		}

		value, err := strconv.ParseFloat(data[0], 32)
		if err != nil {
			return nil, err
		}

		income, err := strconv.ParseFloat(data[1], 32)
		if err != nil {
			return nil, err
		}

		age, err := strconv.ParseFloat(data[2], 32)
		if err != nil {
			return nil, err
		}

		rooms, err := strconv.ParseFloat(data[3], 32)
		if err != nil {
			return nil, err
		}

		bedrooms, err := strconv.ParseFloat(data[4], 32)
		if err != nil {
			return nil, err
		}

		population, err := strconv.ParseFloat(data[5], 32)
		if err != nil {
			return nil, err
		}

		households, err := strconv.ParseFloat(data[6], 32)
		if err != nil {
			return nil, err
		}

		block := common.Block{
			Value:      value,
			Income:     income,
			Age:        age,
			Rooms:      rooms,
			Bedrooms:   bedrooms,
			Population: population,
			Households: households,
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}
