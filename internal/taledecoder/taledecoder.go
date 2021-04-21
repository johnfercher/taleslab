package taledecoder

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
	"strconv"
)

func Base64ToReader(stringBase64 string) (*bufio.Reader, error) {
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

func DecodeSlab(forestBase64 string) (*model.Slab, error) {
	slab := &model.Slab{}
	reader, err := Base64ToReader(forestBase64)
	if err != nil {
		return nil, err
	}

	magicHex, err := DecodeString(reader, 4)
	if err != nil {
		return nil, err
	}
	slab.MagicHex = magicHex

	version, err := DecodeInt2(reader)
	if err != nil {
		return nil, err
	}
	slab.Version = version

	assetCount, err := DecodeInt2(reader)
	if err != nil {
		return nil, err
	}
	slab.AssetsCount = assetCount

	for i := 0; i < assetCount; i++ {
		asset, err := DecodeAsset(reader)
		if err != nil {
			return nil, err
		}

		slab.Assets = append(slab.Assets, asset)
	}

	for i := 0; i < assetCount; i++ {
		assetCount := slab.Assets[i].LayoutsCount

		for j := 0; j < assetCount; j++ {
			bounds, err := DecodeBounds(reader)
			if err != nil {
				return nil, err
			}
			slab.Assets[i].Layouts = append(slab.Assets[i].Layouts, bounds)
		}
	}

	bounds, err := DecodeBounds(reader)
	if err != nil {
		return nil, err
	}

	slab.Bounds = bounds

	return slab, nil
}

func DecodeBounds(reader *bufio.Reader) (*model.Bounds, error) {
	centerX, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	extentsX, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	extentsY, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	extentsZ, err := DecodeFloat(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := DecodeInt1(reader)
	if err != nil {
		return nil, err
	}

	// TODO: understand why this
	_, _ = DecodeString(reader, 3)

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

func DecodeAsset(reader *bufio.Reader) (*model.Asset, error) {
	asset := &model.Asset{}
	uuid, err := DecodeUuid(reader)
	if err != nil {
		return nil, err
	}
	asset.Uuid = uuid

	count, err := DecodeInt2(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

	// TODO: understand why this
	_, _ = DecodeString(reader, 2)

	return asset, nil
}

func DecodeString(buf *bufio.Reader, size int) (string, error) {
	packetBytes := make([]byte, size)

	n, err := buf.Read(packetBytes)
	if err != nil {
		return "", err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	magicHex := ""
	for i := 0; i < n; i++ {
		hex, err := DecodeHex(bufioBuffer)
		if err != nil {
			return "", err
		}

		magicHex += hex + " "
	}

	return magicHex, nil
}

func DecodeUuid(buf *bufio.Reader) (string, error) {
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

func DecodeInt1(buf *bufio.Reader) (int, error) {
	packetBytes := make([]byte, 1)

	_, err := buf.Peek(1)
	if err != nil {
		return 0, nil
	}

	_, err = buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	valueString, err := DecodeHex(bufioBuffer)
	if err != nil {
		return 0, err
	}

	valueInt, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, err
	}

	return valueInt, nil
}

func DecodeInt2(buf *bufio.Reader) (int, error) {
	packetBytes := make([]byte, 2)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	valueString, err := DecodeHex(bufioBuffer)
	if err != nil {
		return 0, err
	}

	valueInt, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, err
	}

	return valueInt, nil
}

func DecodeFloat(buf *bufio.Reader) (float32, error) {
	packetBytes := make([]byte, 4)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint32(packetBytes)
	float := math.Float32frombits(bits)
	return float, nil
}

func DecodeHex(buf *bufio.Reader) (string, error) {
	var packet byte
	err := binary.Read(buf, binary.LittleEndian, &packet)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", packet), nil
}
