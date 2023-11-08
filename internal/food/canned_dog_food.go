package food

var _ Food = (*CannedDogFoodImpl)(nil)

const (
	CannedFoodNutritionalValue = 20
)

type CannedDogFoodImpl struct {
}

// IsTasty implements Food.
func (cdfi *CannedDogFoodImpl) IsTasty() bool {
	return false
}

// Name implements Food.
func (cdfi *CannedDogFoodImpl) Name() string {
	return "Canned Dog Food"
}

// NutritionalValue implements Food.
func (cdfi *CannedDogFoodImpl) NutritionalValue() int {
	return CannedFoodNutritionalValue
}

// Type implements Food.
func (cdfi *CannedDogFoodImpl) Type() string {
	return "FCG"
}
