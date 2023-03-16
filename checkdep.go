package checkdeps

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = "checkdeps is check pkg dependencies"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkdeps",
	Doc:  doc,
	Run:  r.run,
}

type Run struct {
	gomod string             // go module name
	deps  map[string]depsArr // pkg dependencies
	obs   []string           // observed pkgs
}

type depsArr []string

func (d depsArr) notIn(pkg string) bool {
	for _, dep := range d {
		if dep == pkg {
			return false
		}
	}
	return true
}

var r Run

func init() {
	r.gomod = "a"
	r.deps = map[string]depsArr{
		"a/controller": []string{"a/usecase"},
	}
	r.obs = []string{
		"a/controller", "a/usecase",
	}
}

func (r *Run) run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		pkgName := r.pkgName(f.Name.Name)

		for _, i := range f.Imports {
			p, _ := strconv.Unquote(i.Path.Value)
			// TODO: error handling

			if !r.skip(p) && r.deps[pkgName].notIn(p) {
				pass.Reportf(i.Pos(), "error: found bug in dependency import")
			}
		}
	}
	return nil, nil
}

// returns pkgName with gomodule name added
func (r *Run) pkgName(pkg string) string {
	return r.gomod + "/" + pkg
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
