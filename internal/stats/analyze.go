package stats

import (
	"github.com/jeremycruzz/msds301-wk4.git/internal/common"
	"github.com/montanaflynn/stats"
)

// these vars are used for the Describe Function
var (
	allowNan    = true
	percentiles = []float64{25, 50, 75}
	metrics     = []string{"count", "mean", "std", "min", "25%", "50%", "75%", "max"}
)

func (service) Analyze(blocks []common.Block) (map[string]common.Block, error) {
	var err error

	// stores the stats for each block field
	blockStats := map[string]common.Block{}

	// holds a *stats.Description struct for each header
	statData := map[string]*stats.Description{
		"value":    {},
		"income":   {},
		"age":      {},
		"rooms":    {},
		"bedrooms": {},
		"pop":      {},
		"hh":       {},
	}

	pivotBlockData := pivotBlockData(blocks)

	// do stats on all of the columns
	for key, floats := range pivotBlockData {
		statData[key], err = stats.Describe(floats, allowNan, &percentiles)
		if err != nil {
			return nil, err
		}
	}

	for _, metric := range metrics {
		blockStats[metric] = common.Block{
			Value:      getStatValue(statData["value"], metric),
			Income:     getStatValue(statData["income"], metric),
			Age:        getStatValue(statData["age"], metric),
			Rooms:      getStatValue(statData["rooms"], metric),
			Bedrooms:   getStatValue(statData["bedrooms"], metric),
			Population: getStatValue(statData["pop"], metric),
			Households: getStatValue(statData["hh"], metric),
		}
	}

	return blockStats, nil
}

// pivotBlockData pivots the data for easy analysis
func pivotBlockData(blocks []common.Block) map[string][]float64 {
	data := map[string][]float64{
		"value":    {},
		"income":   {},
		"age":      {},
		"rooms":    {},
		"bedrooms": {},
		"pop":      {},
		"hh":       {},
	}

	for _, block := range blocks {
		data["value"] = append(data["value"], block.Value)
		data["income"] = append(data["income"], block.Income)
		data["age"] = append(data["age"], block.Age)
		data["rooms"] = append(data["rooms"], block.Rooms)
		data["bedrooms"] = append(data["bedrooms"], block.Bedrooms)
		data["pop"] = append(data["pop"], block.Population)
		data["hh"] = append(data["hh"], block.Households)

	}
	return data
}

// getStatValue gets the corresponding statistical value for a metric
func getStatValue(description *stats.Description, metric string) float64 {
	switch metric {
	case "count":
		return float64(description.Count)
	case "mean":
		return description.Mean
	case "std":
		return description.Std
	case "min":
		return description.Min
	case "25%":
		return getPercentileValue(description, 25)
	case "50%":
		return getPercentileValue(description, 50)
	case "75%":
		return getPercentileValue(description, 75)
	case "max":
		return description.Max
	default:
		panic("metric not found")
	}
}

// getPercentileValue returns the percentile value for a given percentile
func getPercentileValue(description *stats.Description, percentile float64) float64 {
	for _, dp := range description.DescriptionPercentiles {
		if dp.Percentile == percentile {
			return dp.Value
		}
	}
	panic("percentile not found")
}
