package talespirecontracts

type Asset struct {
	Id           []byte
	IdString     string
	LayoutsCount int16
	Layouts      []*Bounds
}
