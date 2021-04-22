package slabencoder

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
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
	version, err := int16ToByteArray(slab.Version)
	if err != nil {
		return "", err
	}

	slabByteArray = append(slabByteArray, version...)

	// AssetsCount
	assetsCount, err := int16ToByteArray(slab.AssetsCount)
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
		layoutsCount, err := int16ToByteArray(asset.LayoutsCount)
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
			centerX, err := float32ToByteArray(layout.Center.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerX...)

			// Center Y
			centerY, err := float32ToByteArray(layout.Center.Y)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerY...)

			// Center Z
			centerZ, err := float32ToByteArray(layout.Center.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, centerZ...)

			// Extent X
			extentX, err := float32ToByteArray(layout.Extents.X)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentX...)

			// Extent Y
			extentY, err := float32ToByteArray(layout.Extents.Y)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentY...)

			// Extent Z
			extentZ, err := float32ToByteArray(layout.Extents.Z)
			if err != nil {
				return nil, err
			}

			layoutsArray = append(layoutsArray, extentZ...)

			// Rotation
			rotation, err := int8ToByteArray(layout.Rotation)
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
	centerX, err := float32ToByteArray(slab.Bounds.Center.X)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerX...)

	// Center Y
	centerY, err := float32ToByteArray(slab.Bounds.Center.Y)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerY...)

	// Center Z
	centerZ, err := float32ToByteArray(slab.Bounds.Center.Z)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, centerZ...)

	// Extent X
	extentX, err := float32ToByteArray(slab.Bounds.Extents.X)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentX...)

	// Extent Y
	extentY, err := float32ToByteArray(slab.Bounds.Extents.Y)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentY...)

	// Extent Z
	extentZ, err := float32ToByteArray(slab.Bounds.Extents.Z)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, extentZ...)

	// Rotation
	rotation, err := int8ToByteArray(slab.Bounds.Rotation)
	if err != nil {
		return nil, err
	}

	boundsArray = append(boundsArray, rotation...)

	// End of Structure 3
	boundsArray = append(boundsArray, 255, 255, 255)

	return boundsArray, nil
}

func float32ToByteArray(value float32) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func int16ToByteArray(value int16) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func int8ToByteArray(value int8) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
