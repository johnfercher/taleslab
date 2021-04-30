package grid

type ElementType byte

const (
	NoneType       ElementType = 0
	GroundType     ElementType = 1
	MountainType   ElementType = 2
	TreeType       ElementType = 3
	StoneType      ElementType = 4
	RiverType      ElementType = 5
	BaseGroundType ElementType = 6
)

type Element struct {
	Height      uint16
	ElementType ElementType
}
