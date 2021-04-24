package slabdecoder

import (
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv1"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"github.com/johnfercher/taleslab/pkg/versionchecker"
)

type SlabEncoderBuilder interface {
	Build() SlabEncoder
}

type slabEncoderBuilder struct {
}

func NewSlabEncoderBuilder() *slabEncoderBuilder {
	return &slabEncoderBuilder{}
}

func (self *slabEncoderBuilder) Build() SlabEncoder {
	slabCompressor := slabcompressor.New()
	versionChecker := versionchecker.NewVersionChecker(slabCompressor)

	encoderV1 := slabdecoderv1.NewEncoderV1(slabCompressor)
	encoderV2 := slabdecoderv2.NewEncoderV2(slabCompressor)
	return NewEncoder(versionChecker, encoderV1, encoderV2)
}
