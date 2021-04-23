package model

type Bounds struct {
	Coordinates *Vector3d `json:"coordinates"`
	Center      *Vector3f `json:"center"`
	Extents     *Vector3f `json:"extents"`
	Rotation    int8      `json:"rotation"`
	RotationNew int16     `json:"rotation_new"`
}
