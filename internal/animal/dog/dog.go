package dog

import (
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/object"
)

type Dog interface {
	animal.Animal

	DogsName() string
	Breed() string

	// Dogs happiness.
	Play() string
	Happy() int

	// Work Commands.
	Work(energy float32) error

	Bark()
	WagTail() string

	obedienceCommands
}

type obedienceCommands interface {
	Sit() string
	LieDown() string
	Fetch(object object.Object) string
}
