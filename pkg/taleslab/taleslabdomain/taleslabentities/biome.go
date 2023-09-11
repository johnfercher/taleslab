package taleslabentities

import (
	"errors"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/shared/rand"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
)

type Biome struct {
	Type      biometype.BiomeType                 `json:"biome_type"`
	Reliefs   map[elementtype.ElementType]*Relief `json:"reliefs"`
	StoneWall string                              `json:"stone_wall"`
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

func (b *Biome) GetBuildingBlockFromElement(reliefType elementtype.ElementType) (string, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return "", errors.New("unknown relief")
	}

	block := relief.BuildingBlocks
	if len(block) == 0 {
		return "", errors.New("there is no building blocks for this relief type")
	}

	index := rand.Intn(len(block))

	return block[index], nil
}

func (b *Biome) GetPropBlockFromElement(reliefType elementtype.ElementType, propType elementtype.ElementType) (string, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return "", errors.New("unknown relief")
	}

	if propType == elementtype.Tree {
		vegetation := relief.PropBlocks.Vegetation
		if len(vegetation.Props) == 0 {
			return "", errors.New("there is no building blocks for this relief type")
		}

		index := rand.DifferentIntn(len(vegetation.Props), fmt.Sprintf("%s-%s-vegetation", reliefType, propType))

		return vegetation.Props[index], nil
	}

	if propType == elementtype.Stone {
		stones := relief.PropBlocks.Stones
		if len(stones.Props) == 0 {
			return "", errors.New("there is no building blocks for this relief type")
		}

		index := rand.DifferentIntn(len(stones.Props), fmt.Sprintf("%s-%s-stones", reliefType, propType))

		return stones.Props[index], nil
	}

	misc := relief.PropBlocks.Misc
	if len(misc.Props) == 0 {
		return "", errors.New("there is no building blocks for this relief type")
	}

	index := rand.DifferentIntn(len(misc.Props), fmt.Sprintf("%s-%s-misc", reliefType, propType))

	return misc.Props[index], nil
}

func (b *Biome) GetPropBlockWeight(reliefType elementtype.ElementType, propType elementtype.ElementType) (float64, error) {
	relief, ok := b.Reliefs[reliefType]
	if !ok {
		return 0, errors.New("unknown relief")
	}

	if propType == elementtype.Tree {
		vegetation := relief.PropBlocks.Vegetation
		return vegetation.Weight, nil
	}

	if propType == elementtype.Stone {
		stones := relief.PropBlocks.Stones
		return stones.Weight, nil
	}

	misc := relief.PropBlocks.Misc
	return misc.Weight, nil
}
