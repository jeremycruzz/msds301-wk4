package file_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jeremycruzz/msds301-wk4.git/internal/common"
	"github.com/jeremycruzz/msds301-wk4.git/internal/file"
)

func TestFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Suite")
}

var _ = Describe("Read", func() {
	var (
		fileService    file.Service
		testData       string
		tempFilePath   string
		expectedBlocks []common.Block
	)

	Context("with a properly formatted csv file", func() {
		BeforeEach(func() {
			fileService = file.CreateService()
			testData = `value,income,age,rooms,bedrooms,pop,hh
452600,8.3252,41,880,129,322,126
358500,8.3014,21,7099,1106,2401,1138
352100,7.2574,52,1467,190,496,177
`
			tempFile, err := os.CreateTemp("", "test_data_*.csv")
			Expect(err).ToNot(HaveOccurred())

			tempFilePath = tempFile.Name()
			defer tempFile.Close()

			_, err = tempFile.WriteString(testData)
			Expect(err).ToNot(HaveOccurred())

			expectedBlocks = []common.Block{
				{
					Value:      452600,
					Income:     8.325200080871582,
					Age:        41,
					Rooms:      880,
					Bedrooms:   129,
					Population: 322,
					Households: 126,
				},
				{
					Value:      358500,
					Income:     8.301400184631348,
					Age:        21,
					Rooms:      7099,
					Bedrooms:   1106,
					Population: 2401,
					Households: 1138,
				},
				{
					Value:      352100,
					Income:     7.257400035858154,
					Age:        52,
					Rooms:      1467,
					Bedrooms:   190,
					Population: 496,
					Households: 177,
				},
			}
		})

		It("should return a slice of Blocks with the information", func() {

			// execute
			blocks, err := fileService.Read(tempFilePath)

			// assert
			Expect(err).ToNot(HaveOccurred())
			Expect(blocks).To(Equal(expectedBlocks))
		})

	})

	Context("with a csv file with the wrong number of fields", func() {
		BeforeEach(func() {
			fileService = file.CreateService()
			testData = `value,income,age,rooms,bedrooms,pop,hh
452600,41,880,129,322,126
`
			tempFile, err := os.CreateTemp("", "test_data_*.csv")
			Expect(err).ToNot(HaveOccurred())

			tempFilePath = tempFile.Name()
			defer tempFile.Close()

			_, err = tempFile.WriteString(testData)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error", func() {

			// execute
			blocks, err := fileService.Read(tempFilePath)

			// assert
			Expect(blocks).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("wrong number of fields"))
		})
	})

	Context("with a csv file with improperly formatted fields", func() {
		BeforeEach(func() {
			fileService = file.CreateService()
			testData = `value,income,age,rooms,bedrooms,pop,hh
452600,8.3252,41,880,one-hundred,322,126
`
			tempFile, err := os.CreateTemp("", "test_data_*.csv")
			Expect(err).ToNot(HaveOccurred())

			tempFilePath = tempFile.Name()
			defer tempFile.Close()

			_, err = tempFile.WriteString(testData)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error", func() {

			// execute
			blocks, err := fileService.Read(tempFilePath)

			// assert
			Expect(blocks).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid syntax"))
		})
	})

})
