package slabdecoder

import (
	"errors"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv1"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"github.com/johnfercher/taleslab/pkg/versionchecker"
)

type SlabEncoder interface {
	Encode(aggregator *slab.Aggregator) (string, error)
}

type slabEncoder struct {
	versionChecker versionchecker.VersionChecker
	encoderV1      slabdecoderv1.EncoderV1
	encoderV2      slabdecoderv2.EncoderV2
}

func NewEncoder(versionChecker versionchecker.VersionChecker, encoderV1 slabdecoderv1.EncoderV1, encoderV2 slabdecoderv2.EncoderV2) *slabEncoder {
	return &slabEncoder{
		versionChecker: versionChecker,
		encoderV1:      encoderV1,
		encoderV2:      encoderV2,
	}
}

func (self *slabEncoder) Encode(aggregator *slab.Aggregator) (string, error) {
	if aggregator.SlabV1 != nil {
		v1, err := self.encoderV1.Encode(aggregator.SlabV1)
		if err != nil {
			return "", err
		}

		return v1, nil
	}

	if aggregator.SlabV2 != nil {
		v2, err := self.encoderV2.Encode(aggregator.SlabV2)
		if err != nil {
			return "", err
		}

		return v2, nil
	}

	return "", errors.New("invalid_version")
}
