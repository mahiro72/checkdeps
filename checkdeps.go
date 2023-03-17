package checkdeps

import (
	"fmt"
	"os"
	// "strconv"

	// "golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"

	"github.com/mahiro72/checkdeps/pkg/config"
	"github.com/mahiro72/checkdeps/pkg/yml"
	"github.com/mahiro72/checkdeps/internal"
)

const doc = "checkdeps is check pkg dependencies"

// // Analyzer is ...
// var Analyzer = &analysis.Analyzer{
// 	Name: "checkdeps",
// 	Doc:  doc,
// 	Run:  r.run,
// }

// Analyzer is ...
var Analyzer = &internal.Analyzer{
	Name: "checkdeps",
	Doc:  doc,
	Config: &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo |
			packages.NeedModule,
	},
	SSABuilderMode: 0,
	Run:            r.run,
}

type Run struct {
	gomod string             // go module name
	deps  map[string]depsArr // pkg dependencies
	obs   []string           // observed pkgs
}

func init() {
	println("START!!")
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
	b, err := os.ReadFile(config.GetCheckDepsYmlPath("CHECKDEPS_YML"))
	fmt.Println(os.Getwd())
	if err != nil {
		// panic(err)
	}

	d, err := yml.Parse(b)
	if err != nil {
		panic(err)
	}

	r.gomod = d.Spec.Module.Name
	r.deps = r.newDepsFromYmlLayers(d.Spec.Layers)
	r.obs = r.newObsFromYmlObserves(d.Spec.Observes)
}


// func (r *Run) run(pass *analysis.Pass) (interface{}, error) {
func (r *Run) run(pass *internal.Pass) error {
	// r.init()

	fmt.Println(pass.PkgPath)

	fmt.Println("####",pass.Package.Imports)
	fmt.Println("@@@@@",pass.Package.GoFiles)

	// for _, f := range pass.GoFiles {
	// 	fmt.Println("$$$",f)
	// 	pkgName := r.pkgName(f)

	// 	for _, i := range pass.Imports {
	// 		p, err := strconv.Unquote(i.Path.Value)
	// 		if err != nil {
	// 			return nil,err
	// 		}

	// 		if !r.skip(p) && r.deps[pkgName].notIn(p) {
	// 			pass.Reportf(i.Pos(), "error: found bug in dependency import")
	// 		}
	// 	}
	// }
	return nil
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
