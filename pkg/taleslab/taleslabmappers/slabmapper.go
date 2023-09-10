package taleslabmappers

import (
	"bytes"

	"github.com/johnfercher/talescoder/pkg/models"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func TaleSpireSlabFromAssets(assets taleslabentities.Assets) *models.Slab {
	uniqueAssets := getUniqueAssets(assets)
	taleSpire := &models.Slab{
		MagicBytes:  taleslabconsts.MagicBytes,
		Version:     taleslabconsts.SlabVersion,
		AssetsCount: int16(len(uniqueAssets)),
	}

	for _, uniqueAsset := range uniqueAssets {
		layouts := getLayoutFromAsset(uniqueAsset.ID, assets)
		taleSpireAsset := &models.Asset{
			Id:           uniqueAsset.ID,
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
		if uniqueAssets[string(asset.ID)] == nil {
			uniqueAssets[string(asset.ID)] = asset
		}
	}

	return uniqueAssets
}

func getLayoutFromAsset(id []byte, assets taleslabentities.Assets) []*models.Layout {
	bounds := []*models.Layout{}

	for _, asset := range assets {
		if bytes.Equal(id, asset.ID) {
			bound := &models.Layout{
				Coordinates: &models.Vector3d{
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

func AssetsFromTaleSpireSlab(taleSpire *models.Slab) taleslabentities.Assets {
	assets := taleslabentities.Assets{}

	for _, asset := range taleSpire.Assets {
		for _, layout := range asset.Layouts {
			entity := &taleslabentities.Asset{
				ID: asset.Id,
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
