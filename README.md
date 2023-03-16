# checkdeps ![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square) ![Travis](https://img.shields.io/travis/mahiro72/checkdeps.svg?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/mahiro72/checkdeps)](https://goreportcard.com/report/github.com/mahiro72/checkdeps) [![codecov](https://codecov.io/gh/mahiro72/checkdeps/branch/master/graph/badge.svg)](https://codecov.io/gh/mahiro72/checkdeps)

checkdeps is check pkg dependencies

## examples

The error is caught because the usecase should not depend on the repository.
domain/repository ok.

```go
package usecase

import (
	"a4/domain/model"
	"a4/repository" // want "error: found bug in dependency import"
)

type A struct {
	repo repository.A
}

var _ *model.A = (*model.A)(nil)

```


# install

```sh
go install github.com/mahiro72/checkdeps/cmd/checkdeps@latest
```

## Useage

```sh
go vet -vettool=`which checkdeps` pkgname
```
