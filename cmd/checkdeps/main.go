// package main

// import (
// 	"github.com/mahiro72/checkdeps"

// 	"golang.org/x/tools/go/analysis/unitchecker"
// )

// func main() { unitchecker.Main(checkdeps.Analyzer) }

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/mahiro72/checkdeps"
	"github.com/mahiro72/checkdeps/internal"

	"golang.org/x/tools/go/packages"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	checkdeps.Analyzer.Flags = flag.NewFlagSet(checkdeps.Analyzer.Name, flag.ExitOnError)
	checkdeps.Analyzer.Flags.Parse(os.Args[1:])

	if checkdeps.Analyzer.Flags.NArg() < 1 {
		return errors.New("patterns of packages must be specified")
	}

	pkgs, err := packages.Load(checkdeps.Analyzer.Config, checkdeps.Analyzer.Flags.Args()...)
	if err != nil {
		return err
	}

	fmt.Println(pkgs)

	for _, pkg := range pkgs {
		prog, srcFuncs, err := internal.BuildSSA(pkg, checkdeps.Analyzer.SSABuilderMode)
		if err != nil {
			return err
		}

		fmt.Println("?",pkg,pkg.Imports)

		pass := &internal.Pass{
			Package:  pkg,
			SSA:      prog,
			SrcFuncs: srcFuncs,
			Stdin:    os.Stdin,
			Stdout:   os.Stdout,
			Stderr:   os.Stderr,
		}

		if err := checkdeps.Analyzer.Run(pass); err != nil {
			return err
		}
	}

	return nil
}
