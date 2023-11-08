package random

import "math/rand"

var _ Randomiser = (*RandomiserStdImpl)(nil)

type RandomiserStdImpl struct {
}

// Bool implements Randomiser.
func (RandomiserStdImpl) Bool() bool {
	return rand.Intn(2) == 1
}

// Int implements Randomiser.
func (rsi RandomiserStdImpl) Int(max int) int {
	return rand.Intn(max)
}

// Int implements Randomiser.
func (rsi RandomiserStdImpl) Float(max int) float32 {
	return rand.Float32() + float32(rand.Intn(max))
}

func NewRandomiser() RandomiserStdImpl {
	return RandomiserStdImpl{}
}
