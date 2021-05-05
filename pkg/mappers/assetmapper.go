package mappers

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/consts"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecontracts"
)

func AssetInfoFromTalespireContract(assetId, assetType string, slab *talespirecontracts.Slab) *assetloader.AssetInfo {
	assetInfo := &assetloader.AssetInfo{
		Id:   assetId,
		Type: assetType,
	}

	for _, asset := range slab.Assets {
		for _, layout := range asset.Layouts {
			assetPart := &assetloader.AssetPart{
				Id: asset.Id,
				Dimensions: &assetloader.Dimensions{
					Width:  1,
					Length: 1,
					Height: 1,
				},
				OffsetX:  int(layout.Coordinates.X),
				OffsetY:  int(layout.Coordinates.Y),
				OffsetZ:  int(layout.Coordinates.Z),
				Rotation: int(layout.Rotation),
				Name:     "",
			}
			assetInfo.AssertParts = append(assetInfo.AssertParts, assetPart)
		}
	}
	return assetInfo
}

func TalespireContractFromAssetInfo(assetInfo *assetloader.AssetInfo) *talespirecontracts.Slab {
	slab := &talespirecontracts.Slab{
		MagicBytes:  consts.MagicBytes,
		Version:     2,
		AssetsCount: int16(len(assetInfo.AssertParts)),
	}

	for _, assetPart := range assetInfo.AssertParts {
		asset := &talespirecontracts.Asset{
			Id:           assetPart.Id,
			LayoutsCount: 0,
			Layouts:      nil,
		}
		fmt.Println("%s%s", slab, asset)
	}

	return nil
}
