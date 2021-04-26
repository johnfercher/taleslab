package services

import (
	"errors"
	"github.com/johnfercher/taleslab/internal/gridhelper"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
	"time"
)

type SlabBuilder interface {
	SetBiome(biome entities.Biome) SlabBuilder
	SetGround(ground *entities.Ground) SlabBuilder
	SetMountains(mountains *entities.Mountains) SlabBuilder
	SetRiver(river *entities.River) SlabBuilder
	SetProps(props *entities.Props) SlabBuilder
	Build() (string, error)
}

type slabBuilder struct {
	loader        assetloader.AssetLoader
	encoder       slabdecoder.Encoder
	slabGenerated *slab.Slab
	biome         entities.Biome
	props         *entities.Props
	ground        *entities.Ground
	mountains     [][][]uint16
	hasRiver      bool
}

func New(loader assetloader.AssetLoader, encoder slabdecoder.Encoder) *slabBuilder {
	return &slabBuilder{
		loader:        loader,
		encoder:       encoder,
		slabGenerated: &slab.Slab{},
		biome:         entities.ForestBiome,
	}
}

func (self *slabBuilder) SetBiome(biome entities.Biome) SlabBuilder {
	self.biome = biome
	return self
}

func (self *slabBuilder) SetGround(ground *entities.Ground) SlabBuilder {
	self.ground = ground
	return self
}

func (self *slabBuilder) SetMountains(mountains *entities.Mountains) SlabBuilder {
	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(mountains.RandComplexity) + mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(mountains.RandComplexity) + mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			mountainX := rand.Intn(mountains.RandX) + mountains.MinX

			rand.Seed(time.Now().UnixNano())
			mountainY := rand.Intn(mountains.RandY) + mountains.MinY

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(mountains.RandHeight) + mountains.MinHeight)

			generatedMountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			self.mountains = append(self.mountains, generatedMountain)
		}
	}

	return self
}

func (self *slabBuilder) SetRiver(river *entities.River) SlabBuilder {
	if river != nil {
		self.hasRiver = river.HasRiver
	}

	return self
}

func (self *slabBuilder) SetProps(props *entities.Props) SlabBuilder {
	self.props = props
}

func (self *slabBuilder) Build() (string, error) {
	constructors, err := self.loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	props, err := self.loader.GetProps()
	if err != nil {
		log.Fatalln(err)
	}

	slabGenerated := &slab.Slab{
		MagicBytes: slab.MagicBytes,
		Version:    2,
	}

	if self.ground == nil {
		return "", errors.New("Ground not define")
	}

	world := gridhelper.TerrainGenerator(self.ground.Width, self.ground.Length, 2.0, 2.0, self.ground.TerrainComplexity)

	if self.mountains != nil {
		for _, mountain := range self.mountains {
			world = gridhelper.BuildTerrain(world, mountain)
		}
	}

	if self.hasRiver {
		world = gridhelper.DigRiver(world)
	}

	if self.props != nil {
		gridStones := gridhelper.GenerateRandomGridPositions(self.ground.Width, self.ground.Length, self.props.PropsDensity)
		gridTrees := gridhelper.GenerateExclusiveRandomGrid(self.ground.Width, gridStones)
	}

}

func (self *slabBuilder) addLayout(asset *slab.Asset, x, y, z uint16) {
	layout := &slab.Bounds{
		Coordinates: &slab.Vector3d{
			X: x,
			Y: y,
			Z: z,
		},
		Rotation: y / 41,
	}

	asset.Layouts = append(asset.Layouts, layout)
	asset.LayoutsCount++
}

type client struct {
}

func (self *client) ServiceQualquer() {
	forest := entities.Forest{}

	builder := New()

	builder.Init().SetBiome(forest.Biome).
		SetGround(forest.Ground)
}
