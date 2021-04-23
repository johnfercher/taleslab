package slabdecoder

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/johnfercher/taleslab/internal/gzipper"
	"github.com/johnfercher/taleslab/pkg/model"
	"math"
	"strconv"
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

	// TODO: understand why this
	toSkip, _ := decodeInt16(reader)
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

func decodeBounds(reader *bufio.Reader) (*model.Bounds, error) {
	centerX, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}

	centerY, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}

	centerZ, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}

	rotation, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}

	return &model.Bounds{
		Coordinates: &model.Vector3d{
			X: centerX,
			Y: centerY,
			Z: centerZ,
		},
		RotationNew: rotation,
	}, nil
}

func decodeAsset(reader *bufio.Reader) (*model.Asset, error) {
	asset := &model.Asset{}

	// Id
	for i := 0; i < 18; i++ {
		hex, err := decodeInt8(reader)
		if err != nil {
			return nil, err
		}

		asset.Id = append(asset.Id, hex)
	}

	// Count
	count, err := decodeInt16(reader)
	if err != nil {
		return nil, err
	}
	asset.LayoutsCount = count

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

func decodeInt8(buf *bufio.Reader) (int8, error) {
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

	value := int8(16)
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

	value := int16(16)
	err = binary.Read(packetBuffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func decodeInt32(buf *bufio.Reader) (int32, error) {
	packetBytes := make([]byte, 4)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	valueString, err := decodeHex(bufioBuffer)
	if err != nil {
		return 0, err
	}

	valueInt, err := strconv.ParseInt(valueString, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(valueInt), nil
}

func decodeInt64(buf *bufio.Reader) (int64, error) {
	packetBytes := make([]byte, 8)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	valueString, err := decodeHex(bufioBuffer)
	if err != nil {
		return 0, err
	}

	valueInt, err := strconv.ParseInt(valueString, 10, 64)
	if err != nil {
		return 0, err
	}

	return valueInt, nil
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
