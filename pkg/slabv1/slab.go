package slabv1

type Slab struct {
	MagicHex    []byte   `json:"magic_hex"`
	Version     int16    `json:"version"`
	AssetsCount int16    `json:"assets_count"`
	Assets      []*Asset `json:"assets"`
	Bounds      *Bounds  `json:"bounds"`
}
