package taleslabmappers

import (
	talespirecontracts2 "github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func TaleSpireSlabFromAssets(assets taleslabentities.Assets) *talespirecontracts2.Slab {
	uniqueAssets := getUniqueAssets(assets)
	taleSpire := &talespirecontracts2.Slab{
		MagicBytes:  taleslabconsts.MagicBytes,
		Version:     taleslabconsts.SlabVersion,
		AssetsCount: int16(len(uniqueAssets)),
	}

	for _, uniqueAsset := range uniqueAssets {
		layouts := getBoundFromAsset(uniqueAsset.Id, assets)
		taleSpireAsset := &talespirecontracts2.Asset{
			Id:           uniqueAsset.Id,
			LayoutsCount: int16(len(layouts)),
			Layouts:      layouts,
		}

		taleSpire.Assets = append(taleSpire.Assets, taleSpireAsset)
	}

	return taleSpire
}

func getUniqueAssets(assets taleslabentities.Assets) map[string]*taleslabentities.Asset {
	uniqueAssets := make(map[string]*taleslabentities.Asset)

	for _, asset := range assets {
		if uniqueAssets[string(asset.Id)] == nil {
			uniqueAssets[string(asset.Id)] = asset
		}
	}

	return uniqueAssets
}

func getBoundFromAsset(id []byte, assets taleslabentities.Assets) []*talespirecontracts2.Bounds {
	bounds := []*talespirecontracts2.Bounds{}

	for _, asset := range assets {
		if string(id) == string(asset.Id) {
			bound := &talespirecontracts2.Bounds{
				Coordinates: &talespirecontracts2.Vector3d{
					X: uint16(asset.Coordinates.X),
					Y: uint16(asset.Coordinates.Y),
					Z: uint16(asset.Coordinates.Z),
				},
				Rotation: uint16(asset.Rotation),
			}
			bounds = append(bounds, bound)
		}
	}

	return bounds
}

func AssetsFromTaleSpireSlab(taleSpire *talespirecontracts2.Slab) taleslabentities.Assets {
	assets := taleslabentities.Assets{}

	for _, asset := range taleSpire.Assets {
		for _, layout := range asset.Layouts {
			entity := &taleslabentities.Asset{
				Id: asset.Id,
				Coordinates: &taleslabentities.Vector3d{
					X: int(layout.Coordinates.X),
					Y: int(layout.Coordinates.Y),
					Z: int(layout.Coordinates.Z),
				},
			}

			assets = append(assets, entity)
		}
	}

	return assets
}

/*func entityAssetFromTaleSpire(taleSpire *talespirecontracts.Asset) *taleslabentities.Asset {
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
}*/
