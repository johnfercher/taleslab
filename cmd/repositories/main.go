package main

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
)

func main() {
	biomes := taleslabrepositories.NewBiomeRepository()

	biome := biomes.GetBiome(biometype.Beach)
	biome.Print()
}
