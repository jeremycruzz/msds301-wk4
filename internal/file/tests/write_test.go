package file_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jeremycruzz/msds301-wk4.git/internal/common"
	"github.com/jeremycruzz/msds301-wk4.git/internal/file"
)

var _ = Describe("Write", func() {
	var (
		testData     map[string]common.Block
		tempFilePath string
		fileService  file.Service

		expectedStat1 string
		expectedStat2 string
	)

	Context("with a valid filepath", func() {

		BeforeEach(func() {
			testData = map[string]common.Block{
				"stat1": {
					Value:      452600,
					Income:     8.3252,
					Age:        41,
					Rooms:      880,
					Bedrooms:   129,
					Population: 322,
					Households: 126,
				},
				"stat2": {
					Value:      358500,
					Income:     8.3014,
					Age:        21,
					Rooms:      7099,
					Bedrooms:   1106,
					Population: 2401,
					Households: 1138,
				},
			}

			tempFile, err := os.CreateTemp("", "test_data_*.txt")
			Expect(err).ToNot(HaveOccurred())

			tempFilePath = tempFile.Name()

			fileService = file.CreateService("", tempFilePath)
			defer tempFile.Close()

			// TODO: need to find a better way to test a formatted doc
			expectedStat1 = "452600.000000"
			expectedStat2 = "358500.000000"
		})

		It("should write data to the file with the expected format", func() {
			// execute
			err := fileService.Write(testData)

			// assert
			Expect(err).ToNot(HaveOccurred())

			// Check if the written content in the file matches the expected output
			fileContent, err := os.ReadFile(tempFilePath)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(fileContent)).To(ContainSubstring(expectedStat1))
			Expect(string(fileContent)).To(ContainSubstring(expectedStat2))
		})
	})

	Context("with an invalid file path", func() {
		BeforeEach(func() {
			fileService = file.CreateService("", "some/invalid/path")

		})

		It("should return an error", func() {
			// execute
			err := fileService.Write(testData)

			// assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cannot find the path"))
		})
	})
})
