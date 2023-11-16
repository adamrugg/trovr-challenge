package farm_test

import (
	"testing"

	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/dog"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/farmer"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/animal/livestock"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/farm"
	"gitlab.com/Trovr/recruitment/sheep-dog/internal/random"
)

func BenchmarkOldRandomSheepFind(b *testing.B) {
	farm := CreateOldFarm()
	randomiser := random.NewRandomiser()
	for a := 0; a < b.N; a++ {
		for i := 0; i < NumberOfSheep; i++ {
			_, _ = farm.FindSheep(randomiser.Int(NumberOfSheep), "Seed Quadrant 1")
		}
	}
}
func BenchmarkNewRandomSheepFind(b *testing.B) {
	farm := CreateNewFarm()
	randomiser := random.NewRandomiser()
	for a := 0; a < b.N; a++ {
		for i := 0; i < NumberOfSheep; i++ {
			_, _ = farm.FindSheep(randomiser.Int(NumberOfSheep), "Seed Quadrant 2")
		}
	}
}
func BenchmarkOldAscendingFind(b *testing.B) {
	farm := CreateOldFarm()
	for a := 0; a < b.N; a++ {
		for i := 0; i < NumberOfSheep; i++ {
			_, _ = farm.FindSheep(i, "Seed Quadrant 1")
		}
	}
}
func BenchmarkNewAscendingFind(b *testing.B) {
	farm := CreateNewFarm()
	for a := 0; a < b.N; a++ {
		for i := 0; i < NumberOfSheep; i++ {
			_, _ = farm.FindSheep(i, "Seed Quadrant 2")
		}
	}
}

func BenchmarkOldDescendingFind(b *testing.B) {
	farm := CreateOldFarm()
	for a := 0; a < b.N; a++ {
		for i := NumberOfSheep - 1; i >= 0; i-- {
			_, _ = farm.FindSheep(i, "Seed Quadrant 1")
		}
	}
}
func BenchmarkNewDescendingFind(b *testing.B) {
	farm := CreateNewFarm()
	for a := 0; a < b.N; a++ {
		for i := NumberOfSheep - 1; i >= 0; i-- {
			_, _ = farm.FindSheep(i, "Seed Quadrant 2")
		}
	}
}

func BenchmarkOldNotFoundFind(b *testing.B) {
	farm := CreateOldFarm()
	for a := 0; a < b.N; a++ {
		_, _ = farm.FindSheep(2400, "Seed Quadrant 1")
	}
}
func BenchmarkNewNotFoundFind(b *testing.B) {
	farm := CreateNewFarm()
	for a := 0; a < b.N; a++ {
		_, _ = farm.FindSheep(2400, "Seed Quadrant 2")
	}
}

func CreateOldFarm() farm.Farm {
	farmerTed := farmer.NewFarmerImpl(
		"Ted",
		"Sneed",
	)
	oldMonty := dog.NewGermanShepherdDog("Monty", "Bubble", "Linear")
	oldSneedFarm := farm.NewDefaultFarmImpl(farmerTed, "Sneed Farm")
	oldSneedFarm.AddSheepField([]farm.LiveStockField[livestock.Sheep]{
		farm.NewSheepField(farmerTed, "Seed Quadrant 1", CreateSheep(NumberOfSheep), oldMonty),
	})
	return oldSneedFarm
}

func CreateNewFarm() farm.Farm {
	newMonty := dog.NewGermanShepherdDog("Monty", "Quick", "Binary") // replaced default selections with newly implemented quickSort and binarySearch functions
	newSneedFarm := farm.NewDefaultFarmImpl(CreateFarmerTed(), "Sneed Farm")
	newSneedFarm.AddSheepField([]farm.LiveStockField[livestock.Sheep]{
		farm.NewSheepField(CreateFarmerTed(), "Seed Quadrant 2", CreateSheep(NumberOfSheep), newMonty),
	})
	return newSneedFarm
}

func CreateFarmerTed() farmer.Farmer {
	return farmer.NewFarmerImpl(
		"Ted",
		"Sneed",
	)
}
