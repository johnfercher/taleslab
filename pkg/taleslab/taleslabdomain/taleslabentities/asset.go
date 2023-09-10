package taleslabentities

type Asset struct {
	ID          []byte
	Name        string
	Coordinates *Vector3d
	Rotation    int
	Dimensions  *Dimensions
	OffsetZ     int
}

type Assets []*Asset
