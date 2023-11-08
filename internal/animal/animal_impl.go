package animal

import (
	"errors"
	"fmt"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/food"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/random"
)

var _ Animal = (*DefaultAnimalImpl)(nil)

const (
	HealthLevel = 100
	EnergyLevel = 100
	FoodLevel   = 100
)

func NewDefaultAnimalImpl(
	animalName string,
	animalType string,
) *DefaultAnimalImpl {
	return &DefaultAnimalImpl{
		animalName:  animalName,
		animalType:  animalType,
		HealthLevel: HealthLevel,
		EnergyLevel: EnergyLevel,
		FoodLevel:   FoodLevel,
		alive:       true,
		random:      random.NewRandomiser(),
	}
}

type DefaultAnimalImpl struct {
	random      random.Randomiser
	animalName  string
	animalType  string
	HealthLevel float32
	EnergyLevel float32
	FoodLevel   float32
	alive       bool
}

func (ai *DefaultAnimalImpl) Name() string {
	return ai.animalName
}

func (ai *DefaultAnimalImpl) Type() string {
	return ai.animalType
}

// Alive implements Animal.
func (ai *DefaultAnimalImpl) Alive() string {
	if ai.alive {
		return "Alive"
	}
	return "Dead"
}

// TakeDamage implements Animal.
func (ai *DefaultAnimalImpl) TakeDamage(damage float32) {
	if ai.HealthLevel-damage <= 0 {
		ai.alive = false
		return
	}
	ai.HealthLevel -= damage
}

// UseEnergy implements Animal.
func (ai *DefaultAnimalImpl) UseEnergy(energy float32) error {
	value, err := ai.doAction(func() string {
		if ai.EnergyLevel-energy <= 0 {
			return "Energy levels too low"
		}
		ai.EnergyLevel -= energy
		return "Energy Level Decreased"
	})
	switch {
	case err != nil:
		return err
	case value == "Energy levels too low":
		return errors.New(value)
	default:
		return nil
	}
}

// Energy implements Animal.
func (ai *DefaultAnimalImpl) Energy() float32 {
	return ai.EnergyLevel
}

// Health implements Animal.
func (ai *DefaultAnimalImpl) Health() float32 {
	return ai.HealthLevel
}

// Hunger implements Animal.
func (ai *DefaultAnimalImpl) Hunger() float32 {
	return ai.FoodLevel
}

// Eat implements Animal.
func (ai *DefaultAnimalImpl) Eat(food food.Food) (string, error) {
	return ai.doAction(func() string {
		if !food.IsTasty() {
			if !ai.random.Bool() {
				return fmt.Sprintf("food %s was not eaten because its not tasty", food.Name())
			}
		}
		return fmt.Sprintf("eat food %s", food.Name())
	})
}

// Move implements Animal.
func (ai *DefaultAnimalImpl) Move() (string, error) {
	return ai.doAction(func() string { return "moved" })
}

// Sleep implements Animal.
func (ai *DefaultAnimalImpl) Sleep(hours int) (string, error) {
	return ai.doAction(func() string {
		ai.EnergyLevel += float32(hours)
		return fmt.Sprintf("Slept for %d", hours)
	})
}

func (ai *DefaultAnimalImpl) doAction(doAction func() string) (string, error) {
	if !ai.alive {
		return "", errors.New("animal is not alive")
	}
	return doAction(), nil
}
