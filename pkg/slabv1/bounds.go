package slabv1

type Bounds struct {
	Center   *Vector3f `json:"center,omitempty"`
	Extents  *Vector3f `json:"extents,omitempty"`
	Rotation int8      `json:"rotation,omitempty"`
}
