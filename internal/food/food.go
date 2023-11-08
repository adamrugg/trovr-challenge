package food

type Food interface {
	Name() string
	Type() string
	NutritionalValue() int
	IsTasty() bool
}
