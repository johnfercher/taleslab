package slabdecoderv2

import (
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
)

type EncoderV2 interface {
	Encode(slab *slabv2.Slab) (string, error)
}

type encodeV2 struct {
	slabCompressor slabcompressor.SlabCompressor
}

func NewEncoderV2(slabCompressor slabcompressor.SlabCompressor) *encodeV2 {
	return &encodeV2{
		slabCompressor: slabCompressor,
	}
}

func (self *encodeV2) Encode(slab *slabv2.Slab) (string, error) {
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

	// Assets
	assetsBytes, err := self.encodeAssets(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, assetsBytes...)

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

	// Assets.Layouts
	layoutsBytes, err := self.encodeAssetLayouts(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, layoutsBytes...)

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

	slabBase64, err := self.slabCompressor.ByteToStringBase64(slabByteArray)
	if err != nil {
		return "", err
	}

	return slabBase64, nil
}

func (self *encodeV2) encodeAssets(slab *slabv2.Slab) ([]byte, error) {
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

func (self *encodeV2) encodeAssetLayouts(slab *slabv2.Slab) ([]byte, error) {
	layoutsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		for _, layout := range asset.Layouts {
			// Center X
			centerX, err := byteparser.BytesFromUint16(layout.Coordinates.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerX...)

			// Center Z
			centerZ, err := byteparser.BytesFromUint16(layout.Coordinates.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerZ...)

			// Center Y
			yEncoded := EncodeY(layout.Coordinates.Y)

			//fmt.Printf("[ENCODE: %d, %d]\n", layout.Coordinates.Y, yEncoded)

			centerY, err := byteparser.BytesFromUint16(yEncoded)
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
