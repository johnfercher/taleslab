package slabencoder

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/internal/gzipper"
	"github.com/johnfercher/taleslab/pkg/model"
)

func Encode(slab *model.Slab) (string, error) {
	slabByteArray := []byte{}

	// Magic Hex
	for _, magicHex := range slab.MagicHex {
		byte, err := hex.DecodeString(magicHex)
		if err != nil {
			return "", err
		}
		slabByteArray = append(slabByteArray, byte...)
	}

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
	assetsBytes, err := encodeAssets(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, assetsBytes...)

	// Assets.Layouts
	layoutsBytes, err := encodeAssetLayouts(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, layoutsBytes...)

	// Bounds
	boundsBytes, err := encodeBounds(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, boundsBytes...)

	fmt.Println(slabByteArray)

	var buffer bytes.Buffer
	err = gzipper.Compress(&buffer, slabByteArray)
	if err != nil {
		return "", err
	}

	slabByteArrayCompressed := buffer.Bytes()

	slabBase64 := base64.StdEncoding.EncodeToString(slabByteArrayCompressed)

	return slabBase64, nil
}

func encodeAssets(slab *model.Slab) ([]byte, error) {
	assetsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		// Uuid
		id, err := uuid.Parse(asset.Uuid)
		if err != nil {
			return nil, err
		}

		idBytes, err := id.MarshalBinary()
		if err != nil {
			return nil, err
		}
		assetsArray = append(assetsArray, idBytes...)

		// Count
		layoutsCount, err := byteparser.BytesFromInt16(asset.LayoutsCount)
		if err != nil {
			return nil, err
		}

		assetsArray = append(assetsArray, layoutsCount...)

		// End of Structure 2
		assetsArray = append(assetsArray, 0, 0)
	}

	return assetsArray, nil
}

func encodeAssetLayouts(slab *model.Slab) ([]byte, error) {
	layoutsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		for _, layout := range asset.Layouts {
			// Center X
			centerX, err := byteparser.BytesFromFloat32(layout.Center.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerX...)

			// Center Y
			centerY, err := byteparser.BytesFromFloat32(layout.Center.Y)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerY...)

			// Center Z
			centerZ, err := byteparser.BytesFromFloat32(layout.Center.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerZ...)

			// Extent X
			extentX, err := byteparser.BytesFromFloat32(layout.Extents.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentX...)

			// Extent Y
			extentY, err := byteparser.BytesFromFloat32(layout.Extents.Y)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentY...)

			// Extent Z
			extentZ, err := byteparser.BytesFromFloat32(layout.Extents.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentZ...)

			// Rotation
			rotation, err := byteparser.BytesFromInt8(layout.Rotation)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, rotation...)

			// End of Structure 3
			layoutsArray = append(layoutsArray, 0, 0, 0)
		}
	}

	return layoutsArray, nil
}

func encodeBounds(slab *model.Slab) ([]byte, error) {
	boundsArray := []byte{}

	// Center X
	centerX, err := byteparser.BytesFromFloat32(slab.Bounds.Center.X)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerX...)

	// Center Y
	centerY, err := byteparser.BytesFromFloat32(slab.Bounds.Center.Y)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerY...)

	// Center Z
	centerZ, err := byteparser.BytesFromFloat32(slab.Bounds.Center.Z)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerZ...)

	// Extent X
	extentX, err := byteparser.BytesFromFloat32(slab.Bounds.Extents.X)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentX...)

	// Extent Y
	extentY, err := byteparser.BytesFromFloat32(slab.Bounds.Extents.Y)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentY...)

	// Extent Z
	extentZ, err := byteparser.BytesFromFloat32(slab.Bounds.Extents.Z)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentZ...)

	// Rotation
	rotation, err := byteparser.BytesFromInt8(slab.Bounds.Rotation)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, rotation...)

	// End of Structure 3
	boundsArray = append(boundsArray, 255, 255, 255)

	return boundsArray, nil
}
