package versionchecker

import (
	"github.com/johnfercher/taleslab/internal/byteparser"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
)

type VersionChecker interface {
	GetVersion(stringBase64 string) (int16, error)
}

type versionChecker struct {
	slabCompressor slabcompressor.SlabCompressor
}

func NewVersionChecker(slabCompressor slabcompressor.SlabCompressor) *versionChecker {
	return &versionChecker{
		slabCompressor: slabCompressor,
	}
}

func (self *versionChecker) GetVersion(stringBase64 string) (int16, error) {
	reader, err := self.slabCompressor.StringBase64ToReader(stringBase64)
	if err != nil {
		return 0, err
	}

	// Magic Bytes
	_, err = byteparser.BufferToBytes(reader, 4)
	if err != nil {
		return 0, err
	}

	// Version
	version, err := byteparser.BufferToInt16(reader)
	if err != nil {
		return 0, err
	}

	return version, nil
}
