package slabv2

type Bounds struct {
	Coordinates *Vector3d `json:"coordinates,omitempty"`
	Rotation    int16     `json:"rotation,omitempty"`
}
