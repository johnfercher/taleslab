package slabdecoder

import (
	"bufio"
	"fmt"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
)

type Decoder interface {
	Decode(slabBase64 string) (*slab.Slab, error)
}

type decoder struct {
	slabCompressor slabcompressor.SlabCompressor
}

func NewDecoder(slabCompressor slabcompressor.SlabCompressor) *decoder {
	return &decoder{
		slabCompressor: slabCompressor,
	}
}

func (self *decoder) Decode(slabBase64 string) (*slab.Slab, error) {
	slab := &slab.Slab{}

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

func (self *decoder) decodeBounds(reader *bufio.Reader) (*slab.Bounds, error) {
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

	return &slab.Bounds{
		Coordinates: &slab.Vector3d{
			X: centerX,
			Y: centerY,
			Z: centerZ,
		},
		Rotation: rotation,
	}, nil
}

func (self *decoder) decodeAsset(reader *bufio.Reader) (*slab.Asset, error) {
	asset := &slab.Asset{}

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
