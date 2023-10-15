package file

import "github.com/jeremycruzz/msds301-wk4.git/internal/common"

// Service provides methods for reading and writing census block data to a csv
type Service interface {

	// Read takes a filepath and attempts to read the data into a slice of Blocks.
	// It returns a slice of blocks and an error if any occur.
	Read(filepath string) ([]common.Block, error)

	// Write takes in a map of blocks and a filepath.
	// It returns an error if any occur.
	Write(results map[string]common.Block, filepath string) error
}

type service struct{}

// Create Service returns a file service.
func CreateService() service {
	return service{}
}
