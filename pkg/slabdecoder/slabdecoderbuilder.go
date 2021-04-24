package slabdecoder

import (
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv1"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"github.com/johnfercher/taleslab/pkg/versionchecker"
)

type SlabDecoderBuilder interface {
	Build() SlabDecoder
}

type slabDecoderBuilder struct {
}

func NewSlabDecoderBuilder() *slabDecoderBuilder {
	return &slabDecoderBuilder{}
}

func (self *slabDecoderBuilder) Build() SlabDecoder {
	slabCompressor := slabcompressor.New()
	versionChecker := versionchecker.NewVersionChecker(slabCompressor)

	decoderV1 := slabdecoderv1.NewDecoderV1(slabCompressor)
	decoderV2 := slabdecoderv2.NewDecoderV2(slabCompressor)

	return NewDecoder(versionChecker, decoderV1, decoderV2)
}
