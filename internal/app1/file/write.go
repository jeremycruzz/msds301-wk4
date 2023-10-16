package file

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jeremycruzz/msds301-wk4.git/internal/app1/common"
)

func (s service) Write(results map[string]common.Block) error {
	// give permissions (6) to everyone
	file, err := os.OpenFile(s.writeTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// use tabwriter to format nicely
	writer := tabwriter.NewWriter(file, 1, 0, 0, ' ', tabwriter.AlignRight)

	// header row
	header := "x \tvalue \tincome \tage \trooms \tbedrooms \tpopulation \thh \t\n"
	fmt.Fprint(writer, header)

	// write data to file
	for stat, block := range results {
		line := fmt.Sprintf("%s \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t%.6f \t\n",
			stat, block.Value, block.Income, block.Age, block.Rooms, block.Bedrooms, block.Population, block.Households)
		fmt.Fprint(writer, line)
	}

	// flush at end of writing
	writer.Flush()

	return nil
}
