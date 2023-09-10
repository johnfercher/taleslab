package elementtype

type ElementType string

const (
	None ElementType = "none_type"

	BaseGround ElementType = "base_ground"
	Ground     ElementType = "ground"
	Mountain   ElementType = "mountain"
	Water      ElementType = "water"

	Tree  ElementType = "tree"
	Stone ElementType = "stone"
	Misc  ElementType = "misc"
)
