package bytecompressor

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
)

type ByteCompressor interface {
	ToBase64(byteArray []byte) (string, error)
	BufferFromBase64(stringBase64 string) (*bufio.Reader, error)
}

type byteCompressor struct {
}

func New() *byteCompressor {
	return &byteCompressor{}
}

func (self *byteCompressor) ToBase64(byteArray []byte) (string, error) {
	var buffer bytes.Buffer
	err := self.compress(&buffer, byteArray)
	if err != nil {
		return "", err
	}

	slabByteArrayCompressed := buffer.Bytes()

	stringBase64 := base64.StdEncoding.EncodeToString(slabByteArrayCompressed)
	return stringBase64, nil
}

func (self *byteCompressor) BufferFromBase64(stringBase64 string) (*bufio.Reader, error) {
	compressedBytes, err := base64.StdEncoding.DecodeString(stringBase64)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = self.uncompress(&buffer, compressedBytes)
	if err != nil {
		return nil, err
	}

	bufferBytes := buffer.Bytes()

	reader := bytes.NewReader(bufferBytes)
	bufieReader := bufio.NewReader(reader)

	return bufieReader, nil
}

func (self *byteCompressor) uncompress(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer gr.Close()

	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (self *byteCompressor) compress(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	if err != nil {
		return err
	}
	defer gw.Close()

	_, err = gw.Write(data)
	if err != nil {
		return err
	}

	return nil
}
