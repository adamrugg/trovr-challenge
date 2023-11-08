package livestock

import (
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/food"
)

// Verify that sheep implements livestock.
var _ Livestock = (*DefaultSheepImpl)(nil)

// Verify that sheep implements Sheep.
var _ Sheep = (*DefaultSheepImpl)(nil)

// Verify that sheep implements food.
var _ food.Food = (*DefaultSheepImpl)(nil)

// The nutritional value of the sheep.
const FoodNutritionalValue = 60

type DefaultSheepImpl struct {
	animal.Animal
	id                   int
	FoodNutritionalValue int
	FoodIsTasty          bool
}

// NewSheep: Creates a new default sheep.
func NewSheep(id int) DefaultSheepImpl {
	return DefaultSheepImpl{
		Animal: animal.NewDefaultAnimalImpl(
			"Sheep",
			"Livestock",
		),
		FoodIsTasty:          true,
		FoodNutritionalValue: FoodNutritionalValue,
		id:                   id,
	}
}

// ID Will return the livestock ID.
func (dsi DefaultSheepImpl) ID() int {
	return dsi.id
}

// Returns the nutritional value of the sheep.
func (dsi DefaultSheepImpl) NutritionalValue() int {
	return dsi.FoodNutritionalValue
}

// IsTasty will return true if the food is clarified as tasty.
func (dsi DefaultSheepImpl) IsTasty() bool {
	return dsi.FoodIsTasty
}
