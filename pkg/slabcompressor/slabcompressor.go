package slabcompressor

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/taleslab/internal/gzipper"
)

type SlabCompressor interface {
	ByteToStringBase64(byteArray []byte) (string, error)
	StringBase64ToReader(stringBase64 string) (*bufio.Reader, error)
}

type slabCompressor struct {
}

func New() *slabCompressor {
	return &slabCompressor{}
}

func (self *slabCompressor) ByteToStringBase64(byteArray []byte) (string, error) {
	test := fmt.Sprintf("%v", byteArray)
	fmt.Printf("WRITE: ByteArray %s Write Size: %d\n", test, len(byteArray))

	var buffer bytes.Buffer
	err := gzipper.Compress(&buffer, byteArray)
	if err != nil {
		return "", err
	}

	slabByteArrayCompressed := buffer.Bytes()

	stringBase64 := base64.StdEncoding.EncodeToString(slabByteArrayCompressed)
	return stringBase64, nil
}

func (self *slabCompressor) StringBase64ToReader(stringBase64 string) (*bufio.Reader, error) {
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

	test := fmt.Sprintf("%v", bufferBytes)

	fmt.Printf("READ: ByteArray %s Size: %d\n", test, len(bufferBytes))

	reader := bytes.NewReader(bufferBytes)
	bufieReader := bufio.NewReader(reader)

	return bufieReader, nil
}
