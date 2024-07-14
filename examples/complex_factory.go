// Code generated by go-factory 1.0.2; DO NOT EDIT.

package examples

import (
	"github.com/jotadrilo/go-factory/examples/inner"
	alias1 "github.com/jotadrilo/go-factory/examples/inner"
	alias2 "github.com/jotadrilo/go-factory/examples/inner"
	alias3 "github.com/jotadrilo/go-factory/examples/inner"
)

// FactoryComplex is a helper factory to ease creating data of type Complex
type FactoryComplex struct {
	Factory Complex
}

func NewFactoryComplex() *FactoryComplex {
	return &FactoryComplex{}
}

func (f *FactoryComplex) WithName(value string) *FactoryComplex {
	f.Factory.Name = value
	return f
}

func (f *FactoryComplex) WithTypes(value *Types) *FactoryComplex {
	f.Factory.Types = value
	return f
}

func (f *FactoryComplex) Build() *Complex {
	return deepCopyComplex(&f.Factory)
}

func deepCopyComplex(src *Complex) *Complex {
	if src == nil {
		return nil
	}
	copyInstance := *src
	if src.Types != nil {
		copyInstance.Types = new(Types)
		*copyInstance.Types = *src.Types
	}
	return &copyInstance
}

// FactoryComplex2 is a helper factory to ease creating data of type Complex2
type FactoryComplex2 struct {
	Factory Complex2
}

func NewFactoryComplex2() *FactoryComplex2 {
	return &FactoryComplex2{}
}

func (f *FactoryComplex2) WithName(value string) *FactoryComplex2 {
	f.Factory.Name = value
	return f
}

func (f *FactoryComplex2) WithFooPtr(value *inner.Foo) *FactoryComplex2 {
	f.Factory.FooPtr = value
	return f
}

func (f *FactoryComplex2) WithFoo(value alias1.Foo) *FactoryComplex2 {
	f.Factory.Foo = value
	return f
}

func (f *FactoryComplex2) WithFooSlice(values ...alias2.Foo) *FactoryComplex2 {
	f.Factory.FooSlice = values
	return f
}

func (f *FactoryComplex2) AddFooSlice(values ...alias2.Foo) *FactoryComplex2 {
	f.Factory.FooSlice = append(f.Factory.FooSlice, values...)
	return f
}

func (f *FactoryComplex2) WithFooPtrSlice(values ...*alias3.Foo) *FactoryComplex2 {
	f.Factory.FooPtrSlice = values
	return f
}

func (f *FactoryComplex2) AddFooPtrSlice(values ...*alias3.Foo) *FactoryComplex2 {
	f.Factory.FooPtrSlice = append(f.Factory.FooPtrSlice, values...)
	return f
}

func (f *FactoryComplex2) Build() *Complex2 {
	return deepCopyComplex2(&f.Factory)
}

func deepCopyComplex2(src *Complex2) *Complex2 {
	if src == nil {
		return nil
	}
	copyInstance := *src
	if src.FooPtr != nil {
		copyInstance.FooPtr = new(inner.Foo)
		*copyInstance.FooPtr = *src.FooPtr
	}
	if src.FooSlice != nil {
		copyInstance.FooSlice = make([]alias2.Foo, len(src.FooSlice))
		copy(copyInstance.FooSlice, src.FooSlice)
	}
	if src.FooPtrSlice != nil {
		copyInstance.FooPtrSlice = make([]*alias3.Foo, len(src.FooPtrSlice))
		copy(copyInstance.FooPtrSlice, src.FooPtrSlice)
	}
	return &copyInstance
}

