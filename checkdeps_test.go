package checkdeps_test

import (
	"testing"

	"github.com/mahiro72/checkdeps"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	t.Run("層の依存関係に問題がない", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a/...")
	})

	t.Run("usecase層がcontroller層に依存している場合、エラーが発生する", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a2/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a2/...")
	})

	// t.Run("3層のアーキテクチャでも依存関係に問題なければ、エラーは発生しない", func(t *testing.T) {
	// 	analysistest.Run(t, testdata, checkdeps.Analyzer, "a3/...")
	// })
}
