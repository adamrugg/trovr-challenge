package farm

import (
	"errors"
	"fmt"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/farmer"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
)

var _ Farm = (*DefaultFarmImpl)(nil)

type DefaultFarmImpl struct {
	farmer      farmer.Farmer
	farmName    string
	sheepFields []LiveStockField[livestock.Sheep]
}

func NewDefaultFarmImpl(
	owner farmer.Farmer,
	name string,
) *DefaultFarmImpl {
	return &DefaultFarmImpl{
		farmer:      owner,
		sheepFields: nil,
		farmName:    name,
	}
}

// Owner implements Farm.
func (dfi *DefaultFarmImpl) Owner() farmer.Farmer {
	return dfi.farmer
}

// implemented Visit() which uses the Name() func to print the current farm
// a potential  solution could be assigning an value of 'farmCurrentlyOn' to the farmer implementation
// Visit implements Farm.
func (dfi *DefaultFarmImpl) Visit() {
	fmt.Println("Visiting ", dfi.Name())
}

// implemented ListFields() which iterates through the fields and  prin
// ListFields implements Farm.
func (dfi *DefaultFarmImpl) ListFields() {
	fmt.Println("Listing fields:")
	for _, field := range dfi.sheepFields {
		fmt.Printf("Field: %s\n", field.Name())
	}
}

// WorkField implements Farm.
func (dfi *DefaultFarmImpl) WorkField() {
	fmt.Printf("%s is working on %s", dfi.farmer, dfi.farmName)
}

// Name implements Farm.
func (dfi *DefaultFarmImpl) Name() string {
	return dfi.farmName
}

func (dfi *DefaultFarmImpl) AddSheepField(sheepFields []LiveStockField[livestock.Sheep]) {
	dfi.sheepFields = append(dfi.sheepFields, sheepFields...)
}

func (dfi *DefaultFarmImpl) FindSheep(id int, field string) (livestock.Sheep, error) {
	for _, v := range dfi.sheepFields {
		if v.Name() == field {
			sheep, err := v.FindLivestock(id)
			switch {
			case err == nil:
				return sheep, nil
			case err.Error() == "Energy levels too low":
				return nil, err
			}
		}
	}
	return nil, errors.New("not found")
}
