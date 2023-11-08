package dog

import (
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
)

type SheepDog interface {
	Dog
	Herd(sheeps ...livestock.Sheep) ([]livestock.Sheep, error)

	Guard(sheeps ...livestock.Sheep) (bool, error)
	Attack(Animal animal.Animal) (bool, error)

	Find(id int, sortedSheep []livestock.Sheep) (livestock.Sheep, error)
}
