package farm

import (
	fmt2 "fmt"
	"time"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/dog"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/farmer"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
)

var _ LiveStockField[livestock.Sheep] = (*SheepFieldImpl)(nil)

const attackTime = 20 * time.Minute

type SheepFieldImpl struct {
	sheepDog      dog.SheepDog
	owner         farmer.Farmer
	name          string
	fieldType     string
	sheep         []livestock.Sheep
	sortedSheep   []livestock.Sheep
	isSheepSorted bool
}

func NewSheepField(
	owner farmer.Farmer,
	fieldName string,
	sheep []livestock.Sheep,
	sheepDog dog.SheepDog,
) *SheepFieldImpl {
	sheepField := &SheepFieldImpl{
		owner:         owner,
		name:          fieldName,
		fieldType:     "Grazing",
		sheep:         sheep,
		sortedSheep:   nil,
		sheepDog:      sheepDog,
		isSheepSorted: false,
	}
	sheepField.unherd()
	sheepField.predatorAttack()
	return sheepField
}

// Visit implements Field.
func (*SheepFieldImpl) Visit() {
	panic("unimplemented")
}

// Work implements Field.
func (sfi *SheepFieldImpl) FindLivestock(id int) (livestock.Sheep, error) {
	if !sfi.isSheepSorted {
		sortedSheep, err := sfi.sheepDog.Herd(sfi.sheep...)
		if err != nil {
			return nil, err
		}
		sfi.sortedSheep = sortedSheep
		sfi.isSheepSorted = true
	}
	sheep, err := sfi.sheepDog.Find(id, sfi.sortedSheep)
	if err != nil {
		return nil, err
	}
	return sheep, nil
}

// Visit implements Field.
func (sfi *SheepFieldImpl) Name() string {
	return sfi.name
}

func (sfi *SheepFieldImpl) unherd() {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			if _, ok := <-ticker.C; ok {
				sfi.isSheepSorted = false
			}
		}
	}()
}

func (sfi *SheepFieldImpl) predatorAttack() {
	ticker := time.NewTicker(attackTime)
	go func() {
		for {
			if _, ok := <-ticker.C; ok {
				sfi.isSheepSorted = false
				_, err := sfi.sheepDog.Guard(sfi.sheep...)
				if err != nil {
					fmt2.Printf("%+v", err)
				}
			}
		}
	}()
}
