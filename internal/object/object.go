package object

type Object interface {
	Name() string
	Size() (float32, float32, float32)
	IsFetchable() bool
}
