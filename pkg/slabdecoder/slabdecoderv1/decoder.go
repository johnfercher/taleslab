package slabdecoderv1

import (
	"bufio"
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slab/slabv1"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
)

type DecoderV1 interface {
	Decode(slabBase64 string) (*slabv1.Slab, error)
}

type decoderV1 struct {
	slabCompressor slabcompressor.SlabCompressor
}

func NewDecoderV1(slabCompressor slabcompressor.SlabCompressor) *decoderV1 {
	return &decoderV1{
		slabCompressor: slabCompressor,
	}
}

func (self *decoderV1) Decode(slabBase64 string) (*slabv1.Slab, error) {
	slab := &slabv1.Slab{}

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

	// Bounds
	bounds, err := self.decodeBounds(reader)
	if err != nil {
		return nil, err
	}

	slab.Bounds = bounds

	return slab, nil
}

func (self *decoderV1) decodeBounds(reader *bufio.Reader) (*slabv1.Bounds, error) {
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

func (self *decoderV1) decodeAsset(reader *bufio.Reader) (*slabv1.Asset, error) {
	asset := &slabv1.Asset{}

	// Id
	idBytes, err := byteparser.BufferToBytes(reader, 16)
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

	// End of Structure 2
	_, _ = byteparser.BufferToBytes(reader, 2)

	return asset, nil
}
