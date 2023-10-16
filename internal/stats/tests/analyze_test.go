package stats_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jeremycruzz/msds301-wk4.git/internal/common"
	"github.com/jeremycruzz/msds301-wk4.git/internal/stats"
)

var _ = Describe("Analyze", func() {
	var (
		statService        stats.Service
		blocks             []common.Block
		blockStats         map[string]common.Block
		expectedBlockStats map[string]common.Block
	)
	Context("when analyzing blocks", func() {
		BeforeEach(func() {
			statService = stats.CreateService()
			blocks = []common.Block{
				{
					Value:      120.0,
					Income:     9.0,
					Age:        45.0,
					Rooms:      4.5,
					Bedrooms:   2.5,
					Population: 190.0,
					Households: 75.0,
				},
				{
					Value:      130.0,
					Income:     11.0,
					Age:        52.0,
					Rooms:      4.2,
					Bedrooms:   2.2,
					Population: 210.0,
					Households: 85.0,
				},
				{
					Value:      150.0,
					Income:     12.0,
					Age:        55.0,
					Rooms:      5.0,
					Bedrooms:   3.0,
					Population: 220.0,
					Households: 90.0,
				},
				{
					Value:      100.0,
					Income:     10.0,
					Age:        50.0,
					Rooms:      4.0,
					Bedrooms:   2.0,
					Population: 200.0,
					Households: 80.0,
				},
			}

			expectedBlockStats = map[string]common.Block{
				"25%": {
					Value:      100,
					Income:     9,
					Age:        45,
					Rooms:      4,
					Bedrooms:   2,
					Population: 190,
					Households: 75,
				},
				"50%": {
					Value:      120,
					Income:     10,
					Age:        50,
					Rooms:      4.2,
					Bedrooms:   2.2,
					Population: 200,
					Households: 80,
				},
				"75%": {
					Value:      130,
					Income:     11,
					Age:        52,
					Rooms:      4.5,
					Bedrooms:   2.5,
					Population: 210,
					Households: 85,
				},
				"count": {
					Value:      4,
					Income:     4,
					Age:        4,
					Rooms:      4,
					Bedrooms:   4,
					Population: 4,
					Households: 4,
				},
				"max": {
					Value:      150,
					Income:     12,
					Age:        55,
					Rooms:      5,
					Bedrooms:   3,
					Population: 220,
					Households: 90,
				},
				"mean": {
					Value:      125,
					Income:     10.5,
					Age:        50.5,
					Rooms:      4.425,
					Bedrooms:   2.425,
					Population: 205,
					Households: 82.5,
				},
				"min": {
					Value:      100,
					Income:     9,
					Age:        45,
					Rooms:      4,
					Bedrooms:   2,
					Population: 190,
					Households: 75,
				},
				"std": {
					Value:      18.027756377319946,
					Income:     1.118033988749895,
					Age:        3.640054944640259,
					Rooms:      0.37666297933298404,
					Bedrooms:   0.37666297933298404,
					Population: 11.180339887498949,
					Households: 5.5901699437494745,
				},
			}
		})

		It("should analyze the blocks and populate blockStats", func() {
			// execute
			var err error
			blockStats, err = statService.Analyze(blocks)

			// assert
			Expect(err).ToNot(HaveOccurred())
			Expect(blockStats["count"]).To(Equal(expectedBlockStats["count"]))
			Expect(blockStats["25%"]).To(Equal(expectedBlockStats["25%"]))
			Expect(blockStats["50%"]).To(Equal(expectedBlockStats["50%"]))
			Expect(blockStats["75%"]).To(Equal(expectedBlockStats["75%"]))
			Expect(blockStats["min"]).To(Equal(expectedBlockStats["min"]))
			Expect(blockStats["max"]).To(Equal(expectedBlockStats["max"]))
			Expect(blockStats["std"]).To(Equal(expectedBlockStats["std"]))
			Expect(blockStats["mean"]).To(Equal(expectedBlockStats["mean"]))
		})
	})

	Context("when analyzing blocks with bad data", func() {
		//TODO: Add tests
	})

})
