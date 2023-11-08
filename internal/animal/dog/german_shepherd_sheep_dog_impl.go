package dog

import (
	"errors"
	"fmt"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/random"
)

const (
	HerdingWorkEnergy = 15
	GuardWorkEnergy   = 5
	SheepDamage       = 10
	MaxRandomDamage   = 10
)

var _ SheepDog = (*GermanShepherdDogImpl)(nil)

type GermanShepherdDogImpl struct {
	Dog
	random random.Randomiser

	sort   string
	search string
}

// Attack implements Sheep.
func (gsdi *GermanShepherdDogImpl) Attack(
	animal animal.Animal,
) (bool, error) {
	if gsdi.Work(GuardWorkEnergy) == nil {
		status := gsdi.random.Bool()
		if status {
			animal.TakeDamage(
				gsdi.random.Float(MaxRandomDamage),
			)
		}
		return status, nil
	}

	return false, errors.New("could not attack " + animal.Name())
}

// Guard implements Sheep.
func (gsdi *GermanShepherdDogImpl) Guard(
	sheep ...livestock.Sheep,
) (bool, error) {
	if gsdi.Work(GuardWorkEnergy) != nil {
		status := gsdi.random.Bool()
		if !status {
			shp := sheep[gsdi.random.Int(len(sheep))]
			shp.TakeDamage(SheepDamage)

			return false, fmt.Errorf(
				"%s could not defend sheep %d it took %d damage and is %s",
				gsdi.DogsName(), shp.ID(), SheepDamage, shp.Alive(),
			)
		}
	}
	return true, nil
}

// Hurd implements Sheep.
func (gsdi *GermanShepherdDogImpl) Herd(
	sheep ...livestock.Sheep,
) ([]livestock.Sheep, error) {
	if workErr := gsdi.Work(HerdingWorkEnergy); workErr != nil {
		return nil, workErr
	}
	var sortedSheep []livestock.Sheep
	switch gsdi.sort {
	case "Bubble":
		sortedSheep = gsdi.bubbleSortSheep(sheep)
	default:
		sortedSheep = gsdi.bubbleSortSheep(sheep)
	}
	return sortedSheep, nil
}

func (gsdi *GermanShepherdDogImpl) bubbleSortSheep(
	sheeps []livestock.Sheep,
) []livestock.Sheep {
	n := len(sheeps)
	var swapped bool
	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if sheeps[j].ID() > sheeps[j+1].ID() {
				sheeps[j], sheeps[j+1] = sheeps[j+1], sheeps[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sheeps
}

// find implements SheepDog.
func (gsdi *GermanShepherdDogImpl) Find(
	id int,
	sortedSheep []livestock.Sheep,
) (livestock.Sheep, error) {
	var foundSheep livestock.Sheep
	switch gsdi.search {
	case "Linear":
		foundSheep = gsdi.linerSearch(id, sortedSheep)
	default:
		foundSheep = gsdi.linerSearch(id, sortedSheep)
	}
	if foundSheep == nil {
		return nil, errors.New("sheep not found")
	}
	return foundSheep, nil
}

func (gsdi *GermanShepherdDogImpl) linerSearch(
	id int,
	sortedSheep []livestock.Sheep,
) livestock.Sheep {
	for _, v := range sortedSheep {
		if id == v.ID() {
			return v
		}
	}
	return nil
}

func NewGermanShepherdDog(
	name,
	sortMethod,
	searchMethod string,
) *GermanShepherdDogImpl {
	return &GermanShepherdDogImpl{
		Dog: NewDefaultDogImpl(
			name,
			"German Shepherd Dog (GSD)",
			"Working Dog",
		),
		sort:   sortMethod,
		search: searchMethod,
		random: random.NewRandomiser(),
	}
}
