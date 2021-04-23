package slabdecoder

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"github.com/johnfercher/taleslab/internal/gzipper"
	"github.com/johnfercher/taleslab/pkg/model"
	"math"
)

func Decode(slabBase64 string) (*model.Slab, error) {
	slab := &model.Slab{}
	reader, err := base64ToReader(slabBase64)
	if err != nil {
		return nil, err
	}

	// Magic Hex
	for i := 0; i < 4; i++ {
		magicHex, err := decodeHex(reader)
		if err != nil {
			return nil, err
		}

		slab.MagicHex = append(slab.MagicHex, magicHex)
	}

	// Version
	version, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}
	slab.Version = version

	// Assets Count
	assetCount, err := decodeInt16(reader)
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

func decodeBounds(reader *bufio.Reader) (*model.Bounds, error) {
	centerX, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsX, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsY, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	extentsZ, err := decodeFloat32(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := decodeInt8(reader)
	if err != nil {
		return nil, err
	}

	// TODO: understand why this
	_, _ = decodeString(reader, 3)

	return &model.Bounds{
		Center: &model.Vector3{
			X: centerX,
			Y: centerY,
			Z: centerZ,
		},
		Extents: &model.Vector3{
			X: extentsX,
			Y: extentsY,
			Z: extentsZ,
		},
		Rotation: rotation,
	}, nil
}

func decodeAsset(reader *bufio.Reader) (*model.Asset, error) {
	asset := &model.Asset{}

	// Uuid
	uuid, err := decodeUuid(reader)
	if err != nil {
		return nil, err
	}
	asset.Uuid = uuid

	// Count
	count, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

	// End of Structure 2
	_, _ = decodeString(reader, 2)

	return asset, nil
}

func decodeString(buf *bufio.Reader, size int) (string, error) {
	packetBytes := make([]byte, size)

	n, err := buf.Read(packetBytes)
	if err != nil {
		return "", err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	magicHex := ""
	for i := 0; i < n; i++ {
		hex, err := decodeHex(bufioBuffer)
		if err != nil {
			return "", err
		}

		magicHex += hex + " "
	}

	return magicHex, nil
}

func decodeUuid(buf *bufio.Reader) (string, error) {
	packetBytes := make([]byte, 16)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return "", err
	}

	id, err := uuid.FromBytes(packetBytes)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func decodeInt8(buf *bufio.Reader) (int8, error) {
	packetBytes := make([]byte, 1)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)

	value := int8(0)
	err = binary.Read(packetBuffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func decodeInt16(buf *bufio.Reader) (int16, error) {
	packetBytes := make([]byte, 2)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)

	value := int16(0)
	err = binary.Read(packetBuffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func decodeFloat32(buf *bufio.Reader) (float32, error) {
	packetBytes := make([]byte, 4)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint32(packetBytes)
	float := math.Float32frombits(bits)
	return float, nil
}

func decodeHex(buf *bufio.Reader) (string, error) {
	var packet byte
	err := binary.Read(buf, binary.LittleEndian, &packet)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", packet), nil
}
