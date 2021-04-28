package mappers

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecontracts"
)

func TaleSpireSlabFromEntity(entity *entities.Slab) *talespirecontracts.Slab {
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

func taleSpireAssetFromEntity(entity *entities.Asset) *talespirecontracts.Asset {
	taleSpire := &talespirecontracts.Asset{
		Id:           entity.Id,
		LayoutsCount: int16(len(entity.Layouts)),
	}

	for _, layout := range entity.Layouts {
		taleSpire.Layouts = append(taleSpire.Layouts, taleSpireBoundsFromEntity(layout))
	}

	return taleSpire
}

func taleSpireBoundsFromEntity(entity *entities.Bounds) *talespirecontracts.Bounds {
	taleSpire := &talespirecontracts.Bounds{
		Coordinates: &talespirecontracts.Vector3d{
			X: entity.Coordinates.X,
			Y: entity.Coordinates.Y,
			Z: entity.Coordinates.Z,
		},
		Rotation: entity.Rotation,
	}

	return taleSpire
}

func EntitySlabFromTaleSpire(taleSpire *talespirecontracts.Slab) *entities.Slab {
	entity := &entities.Slab{
		MagicBytes: taleSpire.MagicBytes,
		Version:    taleSpire.Version,
	}

	for _, asset := range taleSpire.Assets {
		entity.AddAsset(entityAssetFromTaleSpire(asset))
	}

	return entity
}

func entityAssetFromTaleSpire(taleSpire *talespirecontracts.Asset) *entities.Asset {
	entity := &entities.Asset{
		Id: taleSpire.Id,
	}

	for _, layout := range taleSpire.Layouts {
		entity.Layouts = append(entity.Layouts, entityBoundsFromTaleSpire(layout))
	}

	return entity
}

func entityBoundsFromTaleSpire(taleSpire *talespirecontracts.Bounds) *entities.Bounds {
	entity := &entities.Bounds{
		Coordinates: &entities.Vector3d{
			X: taleSpire.Coordinates.X,
			Y: taleSpire.Coordinates.Y,
			Z: taleSpire.Coordinates.Z,
		},
		Rotation: taleSpire.Rotation,
	}

	return entity
}
