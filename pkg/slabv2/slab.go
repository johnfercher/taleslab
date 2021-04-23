package slabv2

type Slab struct {
	MagicHex    []string `json:"magic_hex,omitempty"`
	Version     int16    `json:"version,omitempty"`
	AssetsCount int16    `json:"assets_count,omitempty"`
	Assets      []*Asset `json:"assets,omitempty"`
}