package grid

type ElementType string

const (
	NoneType       ElementType = "none_type"
	GroundType     ElementType = "ground_type"
	MountainType   ElementType = "mountain_type"
	TreeType       ElementType = "tree_type"
	StoneType      ElementType = "stone_type"
	RiverType      ElementType = "river_type"
	BaseGroundType ElementType = "base_ground_type"
	StoneWallType  ElementType = "stone_wall_type"
	MiscType       ElementType = "misc_type"
)

type Element struct {
	Height      int
	ElementType ElementType
}
