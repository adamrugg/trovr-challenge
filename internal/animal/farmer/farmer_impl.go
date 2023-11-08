package farmer

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"

func NewFarmerImpl(
	firstName,
	lastName string,
) Farmer {
	return Farmer{
		Animal: animal.NewDefaultAnimalImpl(
			"Human",
			"Mammal",
		),
		FirstName: firstName,
		LastName:  lastName,
	}
}
