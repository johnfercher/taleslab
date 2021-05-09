package taleslabentities

import "github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"

type Biome struct {
	BiomeType    taleslabconsts.BiomeType                `json:"biome_type"`
	GroundBlocks map[taleslabconsts.ElementType][]string `json:"ground_blocks"`
	PropBlocks   map[taleslabconsts.ElementType][]string `json:"prop_blocks"`
	StoneWall    string                                  `json:"stone_wall"`
}
