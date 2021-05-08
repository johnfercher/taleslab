package taleslabentities

type Asset struct {
	Id          []byte
	Name        string
	Coordinates *Vector3d
	Rotation    int
	Dimensions  *Dimensions
	OffsetZ     int
}

type Assets []*Asset
