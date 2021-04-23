package slabdecoder2

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
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

	// TODO: understand why this
	toSkip, _ := byteparser.BufferToInt16(reader)
	fmt.Println(toSkip)

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

	for _, bufferByte := range bufferBytes {
		fmt.Printf("0x%X ", bufferByte)
	}

	fmt.Println("")

	fmt.Println(bufferBytes)

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
	for i := 0; i < 18; i++ {
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

	return asset, nil
}
