# SemVer

[![GoDoc](https://godoc.org/github.com/axllent/semver?status.svg)](https://godoc.org/github.com/axllent/semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/axllent/semver)](https://goreportcard.com/report/github.com/axllent/semver)

This semver package provides the ability to sort & compare [Semantic Versions](https://semver.org/) in Go. 

This package follows Semantic Versioning 2.0.0 with one exception: It recognizes MAJOR and MAJOR.MINOR (with no prereleases or build suffixes) as shorthands for MAJOR.0.0 and MAJOR.MINOR.0.

It is based on Golang's [internal](https://github.com/golang/mod/blob/master/semver/semver.go) semver functionality, and has the following features:

- Compare semantic versions, including support for prereleases
- Sort semantic versions (low to high, high to low)
- Ignores any `v` prefix
- See [docs](https://godoc.org/github.com/axllent/semver) for more features

If you require a full-blown solution then try [masterminds/semver](https://github.com/masterminds/semver).


## Usage

```go
if semver.Compare("1.0.1", "1.0.0") == 1 { // higher
  //...
}

if semver.Compare("1.0.1", "1.0.1-beta") == 1 { // higher
  //...
}

fmt.Println(semver.Max("v1.0.1", "v1.0.1-beta"))
// v1.0.1

versions := []string{"5.0.0", "v5.1.0", "5.1.0-beta", "1.2.3"}
fmt.Prinlnln(semver.SortMax(versions))
// [v5.1.0 5.1.0-beta 5.0.0 1.2.3]
```
