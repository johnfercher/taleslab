package taleslabentities

import (
	"errors"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/math"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"math/rand"
)

type Biome struct {
	Type      biometype.BiomeType                    `json:"biome_type"`
	Reliefs   map[taleslabconsts.ElementType]*Relief `json:"reliefs"`
	StoneWall string                                 `json:"stone_wall"`
}

func (b *Biome) Print() {
	if b == nil {
		return
	}

	fmt.Printf("Biome: %s\n", b.Type)
	fmt.Printf("Stone: %s\n", b.StoneWall)
	for key, relief := range b.Reliefs {
		relief.Print(key)
	}
}

func (b *Biome) GetBuildingBlockFromElement(reliefType taleslabconsts.ElementType) (string, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return "", errors.New("unknown relief")
	}

	block := relief.BuildingBlocks
	if len(block) <= 0 {
		return "", errors.New("there is no building blocks for this relief type")
	}

	index := rand.Intn(len(block))

	return block[index], nil
}

func (b *Biome) GetPropBlockFromElement(reliefType taleslabconsts.ElementType, propType taleslabconsts.ElementType) (string, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return "", errors.New("unknown relief")
	}

	if propType == taleslabconsts.Tree {
		vegetation := relief.PropBlocks.Vegetation
		if len(vegetation.Props) <= 0 {
			return "", errors.New("there is no building blocks for this relief type")
		}

		index := math.GetRandomValue(len(vegetation.Props), fmt.Sprintf("%s-%s-vegetation", reliefType, propType))

		return vegetation.Props[index], nil
	}

	if propType == taleslabconsts.Stone {
		stones := relief.PropBlocks.Stones
		if len(stones.Props) <= 0 {
			return "", errors.New("there is no building blocks for this relief type")
		}

		index := math.GetRandomValue(len(stones.Props), fmt.Sprintf("%s-%s-stones", reliefType, propType))

		return stones.Props[index], nil
	}

	misc := relief.PropBlocks.Misc
	if len(misc.Props) <= 0 {
		return "", errors.New("there is no building blocks for this relief type")
	}

	index := math.GetRandomValue(len(misc.Props), fmt.Sprintf("%s-%s-misc", reliefType, propType))

	return misc.Props[index], nil
}

func (b *Biome) GetPropBlockWeight(reliefType taleslabconsts.ElementType, propType taleslabconsts.ElementType) (float64, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return 0, errors.New("unknown relief")
	}

	if propType == taleslabconsts.Tree {
		vegetation := relief.PropBlocks.Vegetation
		return vegetation.Weight, nil
	}

	if propType == taleslabconsts.Stone {
		stones := relief.PropBlocks.Stones
		return stones.Weight, nil
	}

	misc := relief.PropBlocks.Misc
	return misc.Weight, nil
}
