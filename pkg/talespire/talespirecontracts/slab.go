package talespirecontracts

type Slab struct {
	MagicBytes  []byte
	Version     int16
	AssetsCount int16
	Assets      []*Asset
}
