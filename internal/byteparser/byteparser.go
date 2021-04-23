package byteparser

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func BytesFromFloat32(value float32) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesFromInt32(value int32) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesFromInt64(value int64) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesFromInt16(value int16) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesFromInt8(value int8) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BufferToHex(buf *bufio.Reader) (string, error) {
	var packet byte
	err := binary.Read(buf, binary.LittleEndian, &packet)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", packet), nil
}

func BufferString(buf *bufio.Reader, size int) (string, error) {
	packetBytes := make([]byte, size)

	n, err := buf.Read(packetBytes)
	if err != nil {
		return "", err
	}

	packetBuffer := bytes.NewReader(packetBytes)
	bufioBuffer := bufio.NewReader(packetBuffer)

	magicHex := ""
	for i := 0; i < n; i++ {
		hex, err := BufferToHex(bufioBuffer)
		if err != nil {
			return "", err
		}

		magicHex += hex + " "
	}

	return magicHex, nil
}

func BufferToInt16(buf *bufio.Reader) (int16, error) {
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

func BufferToInt8(buf *bufio.Reader) (int8, error) {
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

func BufferToFloat32(buf *bufio.Reader) (float32, error) {
	packetBytes := make([]byte, 4)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint32(packetBytes)
	float := math.Float32frombits(bits)
	return float, nil
}

func BufferToFloat64(buf *bufio.Reader) (float64, error) {
	packetBytes := make([]byte, 8)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	bits := binary.LittleEndian.Uint64(packetBytes)
	float := math.Float64frombits(bits)
	return float, nil
}

func BufferToInt32(buf *bufio.Reader) (int32, error) {
	packetBytes := make([]byte, 4)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)

	value := int32(16)
	err = binary.Read(packetBuffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}
