package taleslabmappers

import (
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func TaleSpireSlabFromEntity(entity *taleslabentities.Slab) *talespirecontracts.Slab {
	assets := entity.GetAssets()

	taleSpire := &talespirecontracts.Slab{
		MagicBytes:  entity.MagicBytes,
		Version:     entity.Version,
		AssetsCount: int16(len(assets)),
	}

	for _, asset := range assets {
		taleSpire.Assets = append(taleSpire.Assets, taleSpireAssetFromEntity(asset))
	}

	return taleSpire
}

func taleSpireAssetFromEntity(entity *taleslabentities.Asset) *talespirecontracts.Asset {
	taleSpire := &talespirecontracts.Asset{
		Id:           entity.Id,
		LayoutsCount: int16(len(entity.Layouts)),
	}

	for _, layout := range entity.Layouts {
		taleSpire.Layouts = append(taleSpire.Layouts, taleSpireBoundsFromEntity(layout))
	}

	return taleSpire
}

func taleSpireBoundsFromEntity(entity *taleslabentities.Bounds) *talespirecontracts.Bounds {
	taleSpire := &talespirecontracts.Bounds{
		Coordinates: &talespirecontracts.Vector3d{
			X: uint16(entity.Coordinates.X),
			Y: uint16(entity.Coordinates.Y),
			Z: uint16(entity.Coordinates.Z),
		},
		Rotation: uint16(entity.Rotation),
	}

	return taleSpire
}

func EntitySlabFromTaleSpire(taleSpire *talespirecontracts.Slab) *taleslabentities.Slab {
	entity := taleslabentities.NewSlab()

	for _, asset := range taleSpire.Assets {
		entity.AddAsset(entityAssetFromTaleSpire(asset))
	}

	return entity
}

func entityAssetFromTaleSpire(taleSpire *talespirecontracts.Asset) *taleslabentities.Asset {
	entity := &taleslabentities.Asset{
		Id: taleSpire.Id,
	}

	for _, layout := range taleSpire.Layouts {
		entity.Layouts = append(entity.Layouts, entityBoundsFromTaleSpire(layout))
	}

	return entity
}

func entityBoundsFromTaleSpire(taleSpire *talespirecontracts.Bounds) *taleslabentities.Bounds {
	entity := &taleslabentities.Bounds{
		Coordinates: &taleslabentities.Vector3d{
			X: int(taleSpire.Coordinates.X),
			Y: int(taleSpire.Coordinates.Y),
			Z: int(taleSpire.Coordinates.Z),
		},
		Rotation: int(taleSpire.Rotation),
	}

	return entity
}
