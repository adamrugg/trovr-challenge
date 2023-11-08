package farm_test

import (
	fmt2 "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/dog"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/farmer"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/farm"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/random"
)

const (
	NumberOfSheep = 20
)

func (fts *FarmTestSuite) TestRandomSheepFind() {
	for i := 0; i < NumberOfSheep; i++ {
		theChosenSheep, err := fts.farm.Old.FindSheep(fts.randomiser.Int(NumberOfSheep), "Seed Quadrant 1")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
	for i := 0; i < NumberOfSheep; i++ {
		theChosenSheep, err := fts.farm.New.FindSheep(fts.randomiser.Int(NumberOfSheep), "Seed Quadrant 2")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
}

func (fts *FarmTestSuite) TestAscendingFind() {
	for i := 0; i < NumberOfSheep; i++ {
		theChosenSheep, err := fts.farm.Old.FindSheep(i, "Seed Quadrant 1")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
	for i := 0; i < NumberOfSheep; i++ {
		theChosenSheep, err := fts.farm.New.FindSheep(i, "Seed Quadrant 2")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
}

func (fts *FarmTestSuite) TestDescendingFind() {
	for i := NumberOfSheep - 1; i >= 0; i-- {
		theChosenSheep, err := fts.farm.Old.FindSheep(i, "Seed Quadrant 1")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
	for i := NumberOfSheep - 1; i >= 0; i-- {
		theChosenSheep, err := fts.farm.New.FindSheep(i, "Seed Quadrant 2")
		fmt2.Printf("Sheep: %+v Error: %+v\n", theChosenSheep, err)
	}
}

func CreateSheep(numberOfSheep int) []livestock.Sheep {
	sheep := make([]livestock.Sheep, numberOfSheep)
	for i := 0; i < numberOfSheep; i++ {
		sheep[i] = livestock.NewSheep(i)
		fmt2.Printf("Sheep Added: %+v \n", sheep[i])
	}
	return sheep
}

type FarmTestSuite struct {
	suite.Suite
	monty      SheepDogs
	farm       Farms
	randomiser random.Randomiser
	farmer     farmer.Farmer
}

type Farms struct {
	Old farm.Farm
	New farm.Farm
}

type SheepDogs struct {
	MontyOld dog.SheepDog
	MontyNew dog.SheepDog
}

func TestSuite(t *testing.T) {
	farmerTed := farmer.NewFarmerImpl(
		"Ted",
		"Sneed",
	)
	oldMonty := dog.NewGermanShepherdDog("Monty", "Bubble", "Linear")
	newMonty := dog.NewGermanShepherdDog("Monty", "Default", "Default")
	oldSneedFarm := farm.NewDefaultFarmImpl(farmerTed, "Sneed Farm")
	newSneedFarm := farm.NewDefaultFarmImpl(farmerTed, "Sneed Farm")
	oldSneedFarm.AddSheepField([]farm.LiveStockField[livestock.Sheep]{
		farm.NewSheepField(farmerTed, "Seed Quadrant 1", CreateSheep(NumberOfSheep), oldMonty),
	})
	newSneedFarm.AddSheepField([]farm.LiveStockField[livestock.Sheep]{
		farm.NewSheepField(farmerTed, "Seed Quadrant 2", CreateSheep(NumberOfSheep), newMonty),
	})

	suite.Run(t, &FarmTestSuite{
		Suite: suite.Suite{},
		monty: SheepDogs{
			MontyOld: oldMonty,
			MontyNew: newMonty,
		},
		farmer: farmerTed,
		farm: Farms{
			Old: oldSneedFarm,
			New: newSneedFarm,
		},
		randomiser: random.NewRandomiser(),
	})
}
