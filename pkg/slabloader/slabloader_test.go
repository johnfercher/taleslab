package slabloader

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
	"os"
	"testing"
)

func init() {
	os.Chdir("../../")
	_, _ = os.Getwd()
}

func TestSlabLoader_GetSlabs(t *testing.T) {
	byteCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(byteCompressor)
	slabLoader := NewSlabLoader(decoder)

	slab, err := slabLoader.GetSlabs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab["general_store"])
}
