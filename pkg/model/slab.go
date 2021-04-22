package model

type Slab struct {
	MagicHex    string   `json:"magic_hex"`
	Version     int      `json:"version"`
	AssetsCount int      `json:"assets_count"`
	Assets      []*Asset `json:"assets"`
	Bounds      *Bounds  `json:"bounds"`
}
