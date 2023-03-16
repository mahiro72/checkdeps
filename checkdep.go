package checkdep

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = "checkdep is check pkg dependencies"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkdep",
	Doc:  doc,
	Run:  r.run,
}

type Run struct {
	gomod string              // go module name
	deps  map[string]string   // pkg dependencies
	obs   []string            // observed pkgs
}

var r Run

func init() {
	r.gomod = "a"

	r.deps = map[string]string{
		"controller":"a/usecase",
	}

	r.obs = []string{
		"a/controller","a/usecase",
	}
	
}

func (r *Run) run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		pkgName := r.pkgName(f.Name.Name)

		for _,i := range f.Imports {
			p, _ := strconv.Unquote(i.Path.Value)
			// TODO: errorハンドリング

			if !r.skip(p) {
				if r.deps[pkgName] != p {
					pass.Reportf(i.Pos(),"error: found bug in dependency import")
				}
			}
		}
	}

	return nil,nil
}

// returns pkgName with gomodule name added
func (r *Run) pkgName(pkg string) string {
	return r.gomod + pkg
}

// skip if not a observed pkg
func (r *Run) skip(pkg string) bool {
	for _, p := range r.obs {
		if p == pkg {
			return false
		}
	}
	return true
}
