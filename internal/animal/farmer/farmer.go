package farmer

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"

type Farmer struct {
	animal.Animal
	FirstName string
	LastName  string
}
