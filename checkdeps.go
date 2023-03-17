package checkdeps

import (
	"fmt"
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
	r.deps = r.newDepsFromYmlLayers(d.Spec.Layers)
	r.obs = r.newObsFromYmlObserves(d.Spec.Observes)
}

func (r *Run) run(pass *analysis.Pass) (interface{}, error) {
	r.init()

	for _, f := range pass.Files {
		pkgName := r.pkgName(f.Name.Name)

		for _, i := range f.Imports {
			p, _ := strconv.Unquote(i.Path.Value)
			// TODO: error handling

			fmt.Println("#######", r.deps[pkgName], p, r.deps[pkgName].notIn(p), !r.skip(p))
			fmt.Println(r.obs)
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

func (r *Run) newDepsFromYmlLayers(layers map[string][]string) map[string]depsArr {
	ret := make(map[string]depsArr)
	for layer, deps := range layers {
		ret[r.pkgName(layer)] = r.newDepsFromYmlDeps(deps)
	}
	return ret
}

func (r *Run) newDepsFromYmlDeps(deps []string) []string {
	var ret []string
	for _, dep := range deps {
		ret = append(ret, r.pkgName(dep))
	}
	return ret
}

func (r *Run) newObsFromYmlObserves(observes []string) []string {
	var ret []string
	for _, obs := range observes {
		ret = append(ret, r.pkgName(obs))
	}
	return ret
}
