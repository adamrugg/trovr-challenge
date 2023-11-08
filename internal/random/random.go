package random

type Randomiser interface {
	Bool() bool
	Int(int) int
	Float(int) float32
}
