package slab

import (
	"github.com/johnfercher/taleslab/pkg/slab/slabv1"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
)

var MagicBytes = []byte{206, 250, 206, 209}

type Aggregator struct {
	SlabV1 *slabv1.Slab `json:"slab_v1,omitempty`
	SlabV2 *slabv2.Slab `json:"slab_v2,omitempty`
}