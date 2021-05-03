package slabloader

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
	"os"
	"path"
	"testing"
)

var AppBaseDir = ""

func init() {
	if AppBaseDir != "" {
		return
	}

	if AppBaseDir == "" {
		base := os.Getenv("GOPATH")
		AppBaseDir = path.Join(base, "src/github.com/johnfercher/taleslab")
	}

	err := os.Chdir(AppBaseDir)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestSlabLoader_GetSlabs(t *testing.T) {
	byteCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(byteCompressor)
	slabLoader, err := NewSlabLoader(decoder)
	if err != nil {
		log.Fatal(err)
	}

	slab := slabLoader.GetSlabs()

	fmt.Println(slab["general_store"])
}
