package taleslabrepositories

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

var AppBaseDir = ""
var rawProps []taleslabentities.Prop

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

	_, err = os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	propsBytes, err := ioutil.ReadFile("./config/assets/props.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(propsBytes, &rawProps)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestNewPropRepository(t *testing.T) {
	// Act
	sut, err := NewPropRepository()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, sut)
	assert.Equal(t, "*taleslabrepositories.propJsonRepository", fmt.Sprintf("%T", sut))
}

func TestAssetLoader_GetProps(t *testing.T) {
	// Arrange
	sut, _ := NewPropRepository()

	// Act
	props := sut.GetProps()

	// Assert
	assert.NotNil(t, props)
	assert.Equal(t, len(rawProps), len(props), "different quantity array loaded x map returned")
	assert.Equal(t, len(getMappedProps()), len(props), "different quantity mappeds x map returned")

	for i := 0; i < len(rawProps); i++ {
		for j := 0; j < len(rawProps); j++ {
			if i != j {
				assert.NotEqual(t, rawProps[i].Id, rawProps[j].Id, fmt.Sprintf("repeated ornaments ids, name: %s", rawProps[i].Id))
				assert.NotEqual(t, rawProps[i].Id, rawProps[j].Id, fmt.Sprintf("repeated ornaments names, id %s", rawProps[i].Id))
			}
		}
	}
}

func TestAssetLoader_GetProp(t *testing.T) {
	// Arrange
	mappedProps := getMappedProps()
	sut, _ := NewPropRepository()

	// Act & Assert
	for mappedPropKey := range mappedProps {
		assert.NotNil(t, sut.GetProp(mappedPropKey), fmt.Sprintf("ornament not loaded %s", mappedPropKey))
	}

	loadedProps := sut.GetProps()
	for loadedPropKey := range loadedProps {
		assert.True(t, mappedProps[loadedPropKey], fmt.Sprintf("ornament not mapped %s", loadedPropKey))
	}
}

func getMappedProps() map[string]bool {
	return map[string]bool{
		"dead_tree_big":                       true,
		"snow_dead_tree_big":                  true,
		"coconut_tree_small":                  true,
		"coconut_tree_big":                    true,
		"pine_tree_big":                       true,
		"rectangle_bush_small":                true,
		"stone_big":                           true,
		"cactus_small":                        true,
		"cactus_big":                          true,
		"snow_pine_tree_big":                  true,
		"snow_stone_big":                      true,
		"snow_stone_small":                    true,
		"big_stone_wall":                      true,
		"big_snow_stone_wall":                 true,
		"one_big_tree":                        true,
		"two_big_tree":                        true,
		"three_big_tree":                      true,
		"bull_skull":                          true,
		"down_rib":                            true,
		"up_rib":                              true,
		"rose":                                true,
		"floor_bush":                          true,
		"small_fern":                          true,
		"big_fern":                            true,
		"big_leaves_bug":                      true,
		"small_leaves_bug":                    true,
		"many_ground_red_leaves":              true,
		"few_ground_red_leaves":               true,
		"small_dead_bush":                     true,
		"big_dead_bush":                       true,
		"tree_trunk":                          true,
		"two_floor_pine_tree":                 true,
		"two_floor_pine_tree_with_trunk":      true,
		"snow_two_floor_pine_tree":            true,
		"snow_two_floor_pine_tree_with_trunk": true,
		"up_skeleton":                         true,
		"fire":                                true,
		"dead_tree_tiny":                      true,
		"dead_tree_small":                     true,
		"cut_tree_with_moss":                  true,
		"down_skeleton":                       true,
		"bush_tree":                           true,
		"pile_of_skulls":                      true,
		"small_green_crystal":                 true,
		"curtains":                            true,
		"ground_nature_big":                   true,
		"ground_nature_with_stones_big":       true,
		"ground_nature_small":                 true,
		"ground_sand_small":                   true,
		"ground_snow_small":                   true,
		"mud_small":                           true,
		"mud_with_feather_small":              true,
		"clay_with_feather_small":             true,
		"dead_land":                           true,
		"cavern_one_rock":                     true,
	}
}
