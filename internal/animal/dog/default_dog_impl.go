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
	return di.Happiness
}

// implemented Bark (work function) which can only be performed if the dog has 5 energy
// Bark implements Dog.
func (di *DefaultDogImpl) Bark() {
	if err := di.Work(5.0); err != nil {
		fmt.Println("Failed to bark:", err.Error())
	}
	fmt.Println(di.DogsName() + ": Woof!")
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
	return fmt.Sprintf("%s lied down", di.DogsName())
}

// Play implements Dog.
func (di *DefaultDogImpl) Play() string {
	di.Happiness += 10
	return fmt.Sprintf("%s is playing", di.DogsName())

}

// Sit: The dog will sit.
func (di *DefaultDogImpl) Sit() string {
	return fmt.Sprintf("%s sat down", di.DogsName())
}

// implemented WagTail (work function) which can only be performed if the dog has 5 energy
// WagTail: The dog will wag its tail.
func (di *DefaultDogImpl) WagTail() string {
	if err := di.Work(3.0); err != nil {
		fmt.Printf("Failed to wag tail: %s\n", err.Error())
	}
	return fmt.Sprintf("%s wagged its tail", di.DogsName())
}

// Work: Is called when the dog does a action and it consumes it energy level.
func (di *DefaultDogImpl) Work(energy float32) error {
	return di.UseEnergy(energy)
}

// DogsName: Will return the dogs name.
func (di *DefaultDogImpl) DogsName() string {
	return di.dogsName
}
