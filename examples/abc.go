package examples

//go:generate go-factory -n A
type A struct {
	Bool   bool
	String string
}

//go:generate go-factory -n B
type B struct {
	Name string
}

//go:generate go-factory -n C
type C struct {
	A  *A
	B  *B
	As []*A
	Bs []*B
}
