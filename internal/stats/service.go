package stats

import "github.com/jeremycruzz/msds301-wk4.git/internal/common"

// Service provides methods for analyzing block data
type Service interface {

	// Analyze takes in a Block slice and calculates count, mean, std, min, 25%, median, 75%, and max for each block field.
	// It returns a map[string]Block where the stat is the key. For example map["mean"] will return a Block whos
	// values are the mean for each field)
	Analyze(blocks []common.Block) (map[string]common.Block, error)
}

type service struct{}

// Create Service creates a stats service.
func CreateService() service {
	return service{}
}
