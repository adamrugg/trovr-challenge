package farm

import (
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/farmer"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
)

type Farm interface {
	Visit()
	Owner() farmer.Farmer
	ListFields()
	WorkField()
	AddSheepField(sheepFields []LiveStockField[livestock.Sheep])
	FindSheep(id int, field string) (livestock.Sheep, error)
	Name() string
}
