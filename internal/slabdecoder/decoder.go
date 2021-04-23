package slabdecoder

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/internal/gzipper"
	"github.com/johnfercher/taleslab/pkg/slabv1"
)

func Decode(slabBase64 string) (*slabv1.Slab, error) {
	slab := &slabv1.Slab{}
	reader, err := base64ToReader(slabBase64)
	if err != nil {
		return nil, err
	}

	// Magic Hex
	for i := 0; i < 4; i++ {
		magicHex, err := byteparser.BufferToByte(reader)
		if err != nil {
			return nil, err
		}

		slab.MagicHex = append(slab.MagicHex, magicHex)
	}

	// Version
	version, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}
	slab.Version = version

	// Assets Count
	assetCount, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}
	slab.AssetsCount = assetCount

	// Assets
	i := int16(0)
	for i = 0; i < assetCount; i++ {
		asset, err := decodeAsset(reader)
		if err != nil {
			return nil, err
		}

		slab.Assets = append(slab.Assets, asset)
	}

	// Assets.Layouts
	i = int16(0)
	for i = 0; i < assetCount; i++ {
		layoutsCount := slab.Assets[i].LayoutsCount

		j := int16(0)
		for j = 0; j < layoutsCount; j++ {
			bounds, err := decodeBounds(reader)
			if err != nil {
				return nil, err
			}
			slab.Assets[i].Layouts = append(slab.Assets[i].Layouts, bounds)
		}
	}

	// Bounds
	bounds, err := decodeBounds(reader)
	if err != nil {
		return nil, err
	}

	slab.Bounds = bounds

	return slab, nil
}

func base64ToReader(stringBase64 string) (*bufio.Reader, error) {
	compressedBytes, err := base64.StdEncoding.DecodeString(stringBase64)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = gzipper.Uncompress(&buffer, compressedBytes)
	if err != nil {
		return nil, err
	}

	bufferBytes := buffer.Bytes()

	fmt.Println(bufferBytes)

	reader := bytes.NewReader(bufferBytes)
	bufieReader := bufio.NewReader(reader)

	return bufieReader, nil
}

func decodeBounds(reader *bufio.Reader) (*slabv1.Bounds, error) {
	centerX, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsX, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsY, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsZ, err := byteparser.BufferToFloat32(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := byteparser.BufferToInt8(reader)
	if err != nil {
		return nil, err
	}

	// TODO: understand why this
	_, _ = byteparser.BufferToBytes(reader, 3)

	return &slabv1.Bounds{
		Center: &slabv1.Vector3f{
			X: centerX,
			Y: centerY,
			Z: centerZ,
		},
		Extents: &slabv1.Vector3f{
			X: extentsX,
			Y: extentsY,
			Z: extentsZ,
		},
		Rotation: rotation,
	}, nil
}

func decodeAsset(reader *bufio.Reader) (*slabv1.Asset, error) {
	asset := &slabv1.Asset{}

	// Id
	for i := 0; i < 16; i++ {
		hex, err := byteparser.BufferToByte(reader)
		if err != nil {
			return nil, err
		}

		asset.Id = append(asset.Id, hex)
	}

	// Count
	count, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

	// End of Structure 2
	_, _ = byteparser.BufferToBytes(reader, 2)

	return asset, nil
}
