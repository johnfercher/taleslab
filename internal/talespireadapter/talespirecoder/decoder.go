package talespirecoder

import (
	"bufio"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/byteparser"
	talespirecontracts2 "github.com/johnfercher/taleslab/internal/talespireadapter/talespirecontracts"
)

type Decoder interface {
	Decode(slabBase64 string) (*talespirecontracts2.Slab, error)
}

type decoder struct {
	slabCompressor bytecompressor.ByteCompressor
}

func NewDecoder(byteCompressor bytecompressor.ByteCompressor) Decoder {
	return &decoder{
		slabCompressor: byteCompressor,
	}
}

func (self *decoder) Decode(slabBase64 string) (*talespirecontracts2.Slab, error) {
	slab := &talespirecontracts2.Slab{}

	reader, err := self.slabCompressor.BufferFromBase64(slabBase64)
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

func (self *decoder) decodeBounds(reader *bufio.Reader) (*talespirecontracts2.Bounds, error) {
	centerX, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := byteparser.BufferToUint16(reader)
	if err != nil {
		return nil, err
	}

	return &talespirecontracts2.Bounds{
		Coordinates: &talespirecontracts2.Vector3d{
			X: DecodeX(centerX),
			Y: DecodeY(centerY),
			Z: DecodeZ(centerZ),
		},
		Rotation: rotation,
	}, nil
}

func (self *decoder) decodeAsset(reader *bufio.Reader) (*talespirecontracts2.Asset, error) {
	asset := &talespirecontracts2.Asset{}

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
