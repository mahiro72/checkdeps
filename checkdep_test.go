package checkdep_test

import (
	"testing"

	"github.com/mahiro72/checkdep"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	// t.Run("層の依存関係に問題がない",func(t *testing.T) {
	// 	analysistest.Run(t, testdata, checkdep.Analyzer, "a/...")
	// })

	t.Run("usecase層がcontroller層に依存している場合、エラーが発生する",func(t *testing.T) {
		analysistest.Run(t, testdata, checkdep.Analyzer, "a2/...")
	})
}
