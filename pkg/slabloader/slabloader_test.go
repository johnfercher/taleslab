package slabloader

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"os"
	"testing"
)

func init() {
	os.Chdir("../../")
	_,_ = os.Getwd()
}

func TestSlabLoader_GetSlabs(t *testing.T) {
	slabLoader := NewSlabLoader(slabdecoder.NewDecoder(slabcompressor.New()))
	slab,err := slabLoader.GetSlabs()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(slab["general_store"])
}