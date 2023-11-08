package animal

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/food"

type Animal interface {
	Name() string
	Type() string

	Eat(food food.Food) (string, error)
	Sleep(hours int) (string, error)
	Move() (string, error)

	TakeDamage(float32)
	UseEnergy(float32) error

	status
}

type status interface {
	Health() float32
	Hunger() float32
	Energy() float32

	Alive() string
}
