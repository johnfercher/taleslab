package slabdecoder2

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/internal/gzipper"
	"github.com/johnfercher/taleslab/pkg/slabv2"
)

func Decode(slabBase64 string) (*slabv2.Slab, error) {
	slab := &slabv2.Slab{}
	reader, err := base64ToReader(slabBase64)
	if err != nil {
		return nil, err
	}

	// Magic Bytes
	magicBytes, err := byteparser.BufferToBytes(reader, 4)
	if err != nil {
		return nil, err
	}

	slab.MagicBytes = append(slab.MagicBytes, magicBytes...)

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

	// TODO: understand why this
	_, _ = byteparser.BufferToInt16(reader)

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

	reader := bytes.NewReader(bufferBytes)
	bufieReader := bufio.NewReader(reader)

	return bufieReader, nil
}

func decodeBounds(reader *bufio.Reader) (*slabv2.Bounds, error) {
	centerX, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}

	return &slabv2.Bounds{
		Coordinates: &slabv2.Vector3d{
			X: centerX,
			Y: centerY,
			Z: centerZ,
		},
		Rotation: rotation,
	}, nil
}

func decodeAsset(reader *bufio.Reader) (*slabv2.Asset, error) {
	asset := &slabv2.Asset{}

	// Id
	idBytes, err := byteparser.BufferToBytes(reader, 18)
	if err != nil {
		return nil, err
	}

	asset.Id = append(asset.Id, idBytes...)

	// Count
	count, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

	return asset, nil
}
