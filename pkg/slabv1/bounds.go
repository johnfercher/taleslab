package slabv1

type Bounds struct {
	Center   *Vector3f `json:"center"`
	Extents  *Vector3f `json:"extents"`
	Rotation int8      `json:"rotation"`
}
