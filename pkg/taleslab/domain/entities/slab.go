package entities

import (
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/consts"
)

func NewSlab() *Slab {
	return &Slab{
		MagicBytes: consts.MagicBytes,
		Version:    2,
		assets:     make(map[string]*Asset),
	}
}

type Slab struct {
	MagicBytes []byte            `json:"magic_bytes"`
	Version    int16             `json:"version"`
	assets     map[string]*Asset `json:"assets"`
}

type Asset struct {
	Id         []byte                  `json:"id"`
	Layouts    []*Bounds               `json:"layouts"`
	Name       string                  `json:"name"`
	Dimensions *assetloader.Dimensions `json:"dimensions"`
	OffsetZ    uint16                  `json:"offset_z"`
}

type Bounds struct {
	Coordinates *Vector3d `json:"coordinates"`
	Rotation    uint16    `json:"rotation"`
}

type Vector3d struct {
	X uint16 `json:"x"`
	Y uint16 `json:"y"`
	Z uint16 `json:"z"`
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
