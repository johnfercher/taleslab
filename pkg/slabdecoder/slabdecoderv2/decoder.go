package slabdecoderv2

import (
	"bufio"
	"fmt"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
)

type DecoderV2 interface {
	Decode(slabBase64 string) (*slabv2.Slab, error)
}

type decoderV2 struct {
	slabCompressor slabcompressor.SlabCompressor
}

func NewDecoderV2(slabCompressor slabcompressor.SlabCompressor) *decoderV2 {
	return &decoderV2{
		slabCompressor: slabCompressor,
	}
}

func (self *decoderV2) Decode(slabBase64 string) (*slabv2.Slab, error) {
	slab := &slabv2.Slab{}

	reader, err := self.slabCompressor.StringBase64ToReader(slabBase64)
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
		asset, err := self.decodeAsset(reader)
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
			bounds, err := self.decodeBounds(reader)
			if err != nil {
				return nil, err
			}
			slab.Assets[i].Layouts = append(slab.Assets[i].Layouts, bounds)
		}
	}

	return slab, nil
}

func (self *decoderV2) decodeBounds(reader *bufio.Reader) (*slabv2.Bounds, error) {
	centerX, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	oldY, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	centerY := DecodeY(oldY)
	//fmt.Printf("[DECODE: %d, %d]\n", oldY, centerY)

	rotation, err := byteparser.BufferToUint16(reader)
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

func (self *decoderV2) decodeAsset(reader *bufio.Reader) (*slabv2.Asset, error) {
	asset := &slabv2.Asset{}

	// Id
	idBytes, err := byteparser.BufferToBytes(reader, 18)
	if err != nil {
		return nil, err
	}

	fmt.Println(idBytes)

	asset.Id = append(asset.Id, idBytes...)

	// Count
	count, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

	return asset, nil
}
