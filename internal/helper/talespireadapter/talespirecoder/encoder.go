package talespirecoder

import (
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	"github.com/johnfercher/taleslab/internal/helper/byteparser"
	"github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecontracts"
)

type Encoder interface {
	Encode(slab *talespirecontracts.Slab) (string, error)
}

type encoder struct {
	slabCompressor bytecompressor.ByteCompressor
}

func NewEncoder(byteCompressor bytecompressor.ByteCompressor) Encoder {
	return &encoder{
		slabCompressor: byteCompressor,
	}
}

func (self *encoder) Encode(slab *talespirecontracts.Slab) (string, error) {
	slabByteArray := []byte{}

	// Magic Hex
	slabByteArray = append(slabByteArray, slab.MagicBytes...)

	// Version
	version, err := byteparser.BytesFromInt16(slab.Version)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, version...)

	// AssetsCount
	assetsCount, err := byteparser.BytesFromInt16(slab.AssetsCount)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, assetsCount...)

	// assets
	assetsBytes, err := self.encodeAssets(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, assetsBytes...)

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

	// assets.Layouts
	layoutsBytes, err := self.encodeAssetLayouts(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, layoutsBytes...)

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

	slabBase64, err := self.slabCompressor.ToBase64(slabByteArray)
	if err != nil {
		return "", err
	}

	return slabBase64, nil
}

func (self *encoder) encodeAssets(slab *talespirecontracts.Slab) ([]byte, error) {
	assetsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		// Id
		for _, assetIdHex := range asset.Id {
			byte, err := byteparser.BytesFromByte(assetIdHex)
			if err != nil {
				return nil, err
			}
			assetsArray = append(assetsArray, byte...)
		}

		// Count
		layoutsCount, err := byteparser.BytesFromInt16(asset.LayoutsCount)
		if err != nil {
			return nil, err
		}

		assetsArray = append(assetsArray, layoutsCount...)
	}

	return assetsArray, nil
}

func (self *encoder) encodeAssetLayouts(slab *talespirecontracts.Slab) ([]byte, error) {
	layoutsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		for _, layout := range asset.Layouts {
			// Center X
			centerX, err := byteparser.BytesFromUint16(EncodeX(layout.Coordinates.X))
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerX...)

			// Center Z
			centerZ, err := byteparser.BytesFromUint16(EncodeZ(layout.Coordinates.Z))
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerZ...)

			// Center Y
			centerY, err := byteparser.BytesFromUint16(EncodeY(layout.Coordinates.Y))
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerY...)

			// Rotation
			rotation, err := byteparser.BytesFromUint16(layout.Rotation)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, rotation...)
		}
	}

	return layoutsArray, nil
}
