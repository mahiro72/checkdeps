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

	// 通常のテスト
	t.Run("依存関係のルールが守られている場合、特にエラーは発生しない (2層アーキテクチャ)", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a/...")
	})

	t.Run("usecaseがcontrollerに依存している場合、依存関係のルールに反するのでエラーが発生する", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a2/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a2/...")
	})

	t.Run("依存関係のルールが守られている場合、特にエラーは発生しない (3層アーキテクチャ)", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a3/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a3/...")
	})

	t.Run("usecaseから触るrepositoryが抽象ではなく実体を指している場合、エラーが発生する", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/a4/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "a4/...")
	})


	// イレギュラーなテスト
	t.Run("repositoryをusecaseという名前でimportしたとき、依存関係のルールに反するのでエラーが発生する", func(t *testing.T) {
		t.Setenv("CHECKDEPS_YML", "./testdata/src/b/checkdeps.yml")
		analysistest.Run(t, testdata, checkdeps.Analyzer, "b/...")
	})
}
