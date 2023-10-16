//go:generate mockgen -source=service.go -destination=../mocks/file_mock.go -package=mocks -mock_names=Service=File
package file

import "github.com/jeremycruzz/msds301-wk4.git/internal/app1/common"

// Service provides methods for reading and writing census block data to a csv
type Service interface {

	// Read takes a filepath and attempts to read the data into a slice of Blocks.
	// It returns a slice of blocks and an error if any occur.
	Read() ([]common.Block, error)

	// Write takes in a map of blocks and a filepath.
	// It returns an error if any occur.
	Write(results map[string]common.Block) error
}

type service struct {
	readFrom string
	writeTo  string
}

// Create Service returns a file service.
func CreateService(readFrom, writeTo string) service {
	return service{
		readFrom: readFrom,
		writeTo:  writeTo,
	}
}
