package dog

import (
	"fmt"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/object"
)

var _ Dog = (*DefaultDogImpl)(nil)

const (
	defaultHappiness = 100
)

func NewDefaultDogImpl(
	name,
	breed,
	animalType string,
) *DefaultDogImpl {
	return &DefaultDogImpl{
		dogsName: name,
		DogBreed: breed,
		Animal: animal.NewDefaultAnimalImpl(
			"Dog",
			animalType,
		),
		Happiness: defaultHappiness,
		Hungry:    false,
	}
}

type DefaultDogImpl struct {
	animal.Animal

	dogsName string
	DogBreed string

	Happiness int

	Hungry bool
}

// Breed implements Dog.
func (di DefaultDogImpl) Breed() string {
	return di.DogBreed
}

// Happy implements Dog.
func (di DefaultDogImpl) Happy() int {
	panic("unimplemented")
}

// Bark implements Dog.
func (di *DefaultDogImpl) Bark() {
	panic("unimplemented")
}

// Fetch implements Dog.
func (di *DefaultDogImpl) Fetch(object object.Object) string {
	if object.IsFetchable() {
		return fmt.Sprintf("%s fetched %s", di.DogsName(), object.Name())
	}
	height, width, depth := object.Size()
	return fmt.Sprintf("%s can not fetch %s as it is too big %v,%v,%v", di.DogsName(), object.Name(), height, width, depth)
}

// LieDown implements Dog.
func (di *DefaultDogImpl) LieDown() string {
	panic("unimplemented")
}

// Play implements Dog.
func (di *DefaultDogImpl) Play() string {
	panic("unimplemented")
}

// Sit: The dog will sit.
func (di *DefaultDogImpl) Sit() string {
	panic("unimplemented")
}

// WagTail: The dog will wag its tail.
func (di *DefaultDogImpl) WagTail() string {
	panic("unimplemented")
}

// Work: Is called when the dog does a action and it consumes it energy level.
func (di *DefaultDogImpl) Work(energy float32) error {
	return di.UseEnergy(energy)
}

// DogsName: Will return the dogs name.
func (di *DefaultDogImpl) DogsName() string {
	return di.dogsName
}
