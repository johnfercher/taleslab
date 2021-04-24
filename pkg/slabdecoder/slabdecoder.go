package slabdecoder

import (
	"errors"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv1"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"github.com/johnfercher/taleslab/pkg/versionchecker"
)

type SlabDecoder interface {
	Decode(slabBase64 string) (*slab.Aggregator, error)
}

type slabDecoder struct {
	versionChecker versionchecker.VersionChecker
	decoderV1      slabdecoderv1.DecoderV1
	decoderV2      slabdecoderv2.DecoderV2
}

func NewDecoder(versionChecker versionchecker.VersionChecker, decoderV1 slabdecoderv1.DecoderV1, decoderV2 slabdecoderv2.DecoderV2) *slabDecoder {
	return &slabDecoder{
		versionChecker: versionChecker,
		decoderV1:      decoderV1,
		decoderV2:      decoderV2,
	}
}

func (self *slabDecoder) Decode(slabBase64 string) (*slab.Aggregator, error) {
	version, err := self.versionChecker.GetVersion(slabBase64)
	if err != nil {
		return nil, err
	}

	if version == 1 {
		v1, err := self.decoderV1.Decode(slabBase64)
		if err != nil {
			return nil, err
		}

		return &slab.Aggregator{SlabV1: v1}, nil
	}

	if version == 2 {
		v2, err := self.decoderV2.Decode(slabBase64)
		if err != nil {
			return nil, err
		}

		return &slab.Aggregator{SlabV2: v2}, nil
	}

	return nil, errors.New("invalid_version")
}
