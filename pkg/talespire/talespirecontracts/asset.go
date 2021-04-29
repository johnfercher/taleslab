package talespirecontracts

type Asset struct {
	Id           []byte
	LayoutsCount int16
	Layouts      []*Bounds
}
