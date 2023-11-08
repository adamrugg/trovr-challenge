package predator

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"

var _ Predator = (*FoxPredatorImpl)(nil)

type FoxPredatorImpl struct {
	animal.Animal
}
