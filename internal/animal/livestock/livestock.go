package livestock

import "gitlab.com/Trovr/recruitment/sheep-dog/internal/animal"

// Livestock: provides the implementation abstraction for all livestock animals.
type Livestock interface {
	animal.Animal // requires the implementation of the animal interface.
	ID() int      // Returns the ID of the livestock.
}
