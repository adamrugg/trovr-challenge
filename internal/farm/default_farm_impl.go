package farm

import (
	"errors"

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

// Visit implements Farm.
func (*DefaultFarmImpl) Visit() {
	panic("unimplemented")
}

// ListFields implements Farm.
func (*DefaultFarmImpl) ListFields() {
	panic("unimplemented")
}

// WorkField implements Farm.
func (dfi *DefaultFarmImpl) WorkField() {

}

// WorkField implements Farm.
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
