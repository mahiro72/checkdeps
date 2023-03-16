# checkdep ![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square) ![Travis](https://img.shields.io/travis/mahiro72/checkdep.svg?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/mahiro72/checkdep)](https://goreportcard.com/report/github.com/mahiro72/checkdep) [![codecov](https://codecov.io/gh/mahiro72/checkdep/branch/master/graph/badge.svg)](https://codecov.io/gh/mahiro72/checkdep)

checkdep is check pkg dependencies


# install

```sh
go install github.com/mahiro72/checkdep/cmd/checkdep@latest
```

## Useage

```sh
go vet -vettool=`which checkdep` pkgname
```
