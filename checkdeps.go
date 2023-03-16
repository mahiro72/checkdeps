package checkdeps

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/mahiro72/checkdeps/pkg/yml"
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

func (r *Run) init() {
	checkDepsPath := os.Getenv("CHECKDEPS_YML")
	if checkDepsPath == "" {
		panic("error: not found checkdeps.yml")
	}

	//FIXME: 非推奨かも?
	b, err := ioutil.ReadFile(checkDepsPath)
	if err != nil {
		panic(err)
	}

	d, err := yml.Parse(b)
	if err != nil {
		panic(err)
	}

	r.gomod = d.Spec.Module.Name
	r.deps = newDepsFromYmlLayers(d.Spec.Layers)
	r.obs = newObsFromYmlObserves(d.Spec.Observes)
}

func (r *Run) run(pass *analysis.Pass) (any, error) {
	r.init()

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

func newDepsFromYmlLayers(layers map[string][]string) map[string]depsArr {
	r := make(map[string]depsArr)
	for layer, deps := range layers {
		r[layer] = deps
	}
	return r
}

func newObsFromYmlObserves(observes []string) []string {
	var r []string
	r = append(r, observes...)
	return r
}
