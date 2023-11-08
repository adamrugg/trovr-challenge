package livestock

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"

// sleep: provides the implementation abstraction for sheep.
type Sheep interface {
	animal.Animal // requires the implementation of the animal interface.
	ID() int      // Returns the ID of the sheep.
}
