package model

type Bounds struct {
	Center   *Vector3 `json:"center"`
	Extents  *Vector3 `json:"extents"`
	Rotation int8     `json:"rotation"`
}
