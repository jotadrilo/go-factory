package examples

import (
	"github.com/jotadrilo/go-factory/examples/inner"
	alias1 "github.com/jotadrilo/go-factory/examples/inner"
	alias2 "github.com/jotadrilo/go-factory/examples/inner"
	alias3 "github.com/jotadrilo/go-factory/examples/inner"
)

//go:generate go-factory -n Complex
type Complex struct {
	Name  string
	Types *Types
}

//go:generate go-factory -n Complex2
type Complex2 struct {
	Name        string
	FooPtr      *inner.Foo
	Foo         alias1.Foo
	FooSlice    []alias2.Foo
	FooPtrSlice []*alias3.Foo
}
