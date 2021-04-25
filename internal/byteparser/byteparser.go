package byteparser

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math"
)

func BytesFromByte(value byte) ([]byte, error) {
	return []byte{value}, nil
}

func BufferToBytes(buf *bufio.Reader, bytesCount int) ([]byte, error) {
	packetBytes := make([]byte, bytesCount)

	_, err := buf.Peek(bytesCount)
	if err != nil {
		return nil, err
	}

	_, err = buf.Read(packetBytes)
	if err != nil {
		return nil, err
	}

	return packetBytes, nil
}

func BytesFromInt8(value int8) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

func BytesFromInt16(value int16) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesFromUint16(value uint16) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

func BufferToUint16(buf *bufio.Reader) (uint16, error) {
	packetBytes := make([]byte, 2)

	_, err := buf.Read(packetBytes)
	if err != nil {
		return 0, err
	}

	packetBuffer := bytes.NewReader(packetBytes)

	value := uint16(16)
	err = binary.Read(packetBuffer, binary.LittleEndian, &value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func BytesFromFloat32(value float32) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

/*

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



*/
