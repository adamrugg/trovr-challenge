package object

var _ Object = (*BallObjectImpl)(nil)

func NewBall(
	height, // The sticks height in millimetres.
	width, // The sticks width in millimetres.
	depth float32, // The sticks depth in millimetres.
) BallObjectImpl {
	var fetchable bool
	switch {
	case height > MaxFetchSize:
		fetchable = false
	case width > MaxFetchSize:
		fetchable = false
	case depth > MaxFetchSize:
		fetchable = false
	default:
		fetchable = true
	}

	return BallObjectImpl{
		name:      "Ball",
		fetchable: fetchable,
		height:    height,
		width:     width,
		depth:     depth,
	}
}

type BallObjectImpl struct {
	name string

	fetchable bool

	height float32
	width  float32
	depth  float32
}

// IsFetchable implements Object.
func (soi *BallObjectImpl) IsFetchable() bool {
	return soi.fetchable
}

// Name implements Object.
func (soi *BallObjectImpl) Name() string {
	return soi.name
}

// Size implements Object.
func (soi *BallObjectImpl) Size() (float32, float32, float32) {
	return soi.height, soi.width, soi.depth
}
