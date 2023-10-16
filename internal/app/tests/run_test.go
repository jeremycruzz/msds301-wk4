package app_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	"github.com/jeremycruzz/msds301-wk4.git/internal/app"
	"github.com/jeremycruzz/msds301-wk4.git/internal/common"
	"github.com/jeremycruzz/msds301-wk4.git/internal/mocks"
)

var _ = Describe("app", func() {
	var (
		ctrl      *gomock.Controller
		mockFile  *mocks.File
		mockStats *mocks.Stats

		readReturnData    []common.Block
		analyzeReturnData map[string]common.Block
	)
	Context("when there are no errors", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFile = mocks.NewFile(ctrl)
			mockStats = mocks.NewStats(ctrl)

			readReturnData = []common.Block{{1, 2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 8}}
			analyzeReturnData = map[string]common.Block{"any": {3, 4, 5, 6, 7, 8, 9}, "other": {4, 5, 6, 7, 8, 9, 10}}

		})

		It("should read, analyze, and write", func() {
			mockFile.EXPECT().Read().Return(readReturnData, nil)
			mockStats.EXPECT().Analyze(readReturnData).Return(analyzeReturnData, nil)
			mockFile.EXPECT().Write(analyzeReturnData).Return(nil)

			testApp := app.CreateApp(mockFile, mockStats)
			testApp.Run()

		})
	})

	Context("when there is an error in read", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFile = mocks.NewFile(ctrl)
			mockStats = mocks.NewStats(ctrl)

			readReturnData = []common.Block{{1, 2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 8}}
			analyzeReturnData = map[string]common.Block{"any": {3, 4, 5, 6, 7, 8, 9}, "other": {4, 5, 6, 7, 8, 9, 10}}

		})

		It("should panic", func() {
			mockFile.EXPECT().Read().Return(nil, errors.New("read error"))

			Expect(func() {
				testApp := app.CreateApp(mockFile, mockStats)
				testApp.Run()
			}).To(Panic())
		})
	})
	Context("when there is an error in analyze", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFile = mocks.NewFile(ctrl)
			mockStats = mocks.NewStats(ctrl)

			readReturnData = []common.Block{{1, 2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 8}}
			analyzeReturnData = map[string]common.Block{"any": {3, 4, 5, 6, 7, 8, 9}, "other": {4, 5, 6, 7, 8, 9, 10}}

		})

		It("should panic", func() {
			mockFile.EXPECT().Read().Return(readReturnData, nil)
			mockStats.EXPECT().Analyze(readReturnData).Return(nil, errors.New("analyze error"))

			Expect(func() {
				testApp := app.CreateApp(mockFile, mockStats)
				testApp.Run()
			}).To(Panic())
		})
	})
	Context("when there is an error in write", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFile = mocks.NewFile(ctrl)
			mockStats = mocks.NewStats(ctrl)

			readReturnData = []common.Block{{1, 2, 3, 4, 5, 6, 7}, {2, 3, 4, 5, 6, 7, 8}}
			analyzeReturnData = map[string]common.Block{"any": {3, 4, 5, 6, 7, 8, 9}, "other": {4, 5, 6, 7, 8, 9, 10}}

		})

		It("should panic", func() {
			mockFile.EXPECT().Read().Return(readReturnData, nil)
			mockStats.EXPECT().Analyze(readReturnData).Return(analyzeReturnData, nil)
			mockFile.EXPECT().Write(analyzeReturnData).Return(errors.New("write error"))

			Expect(func() {
				testApp := app.CreateApp(mockFile, mockStats)
				testApp.Run()
			}).To(Panic())
		})
	})
})
