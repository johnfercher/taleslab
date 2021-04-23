package slabencoder2

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

	// Assets.Layouts
	layoutsBytes, err := encodeAssetLayouts(slab)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, layoutsBytes...)

	// End of Structure 2
	slabByteArray = append(slabByteArray, 0, 0)

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
		for _, assetIdHex := range asset.Id {
			byte, err := byteparser.BytesFromInt8(assetIdHex)
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

func encodeAssetLayouts(slab *model.Slab) ([]byte, error) {
	layoutsArray := []byte{}

	// For
	for _, asset := range slab.Assets {
		for _, layout := range asset.Layouts {
			// Center X
			centerX, err := byteparser.BytesFromInt16(layout.Coordinates.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerX...)

			// Center Y
			centerY, err := byteparser.BytesFromInt16(layout.Coordinates.Y)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerY...)

			// Center Z
			centerZ, err := byteparser.BytesFromInt16(layout.Coordinates.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerZ...)

			// Rotation
			rotationNew, err := byteparser.BytesFromInt16(layout.RotationNew)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, rotationNew...)
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
