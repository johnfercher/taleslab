package taleslabentities

import (
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
)

func NewSlab() *Slab {
	return &Slab{
		MagicBytes: taleslabconsts.MagicBytes,
		Version:    2,
		assets:     make(map[string]*Asset),
	}
}

type Slab struct {
	MagicBytes []byte
	Version    int16
	assets     map[string]*Asset
}

type Asset struct {
	Id         []byte
	Layouts    []*Bounds
	Name       string
	Dimensions *assetloader.Dimensions
	OffsetZ    int
}

type Bounds struct {
	Coordinates *Vector3d
	Rotation    int
}

type Vector3d struct {
	X int
	Y int
	Z int
}

func (self *Slab) AddAsset(asset *Asset) {
	idString := string(asset.Id)
	if self.assets[idString] == nil {
		self.assets[idString] = asset
		return
	}

	oldAsset := self.assets[idString]

	oldAsset.Layouts = append(oldAsset.Layouts, asset.Layouts...)

	self.assets[idString] = oldAsset
}

func (self *Slab) AddLayoutToAsset(assetId []byte, layout *Bounds) {
	if self.assets[string(assetId)] == nil {
		self.assets[string(assetId)] = &Asset{
			Id:      assetId,
			Layouts: []*Bounds{layout},
		}

		return
	}

	asset := self.assets[string(assetId)]

	asset.Layouts = append(asset.Layouts, layout)

	self.assets[string(assetId)] = asset
}

func (self *Slab) GetAsset(id string) *Asset {
	return self.assets[id]
}

func (self *Slab) GetAssets() map[string]*Asset {
	return self.assets
}
