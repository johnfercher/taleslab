package taleslabrepositories_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"

	"github.com/stretchr/testify/assert"
)

func TestNewPropRepository_WhenCantCreate_ShouldReturnError(t *testing.T) {
	// Act
	sut, err := taleslabrepositories.NewPropRepository()

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, sut)
}

func TestNewPropRepository_WhenCanCreate_ShouldReturnRepository(t *testing.T) {
	// Act
	sut, err := taleslabrepositories.NewPropRepository(buildPropsPath())

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, sut)
	assert.Equal(t, "*taleslabrepositories.propJSONRepository", fmt.Sprintf("%T", sut))
}

func TestPropJSONRepository_GetProps(t *testing.T) {
	// Arrange
	sut, _ := taleslabrepositories.NewPropRepository(buildPropsPath())

	// Act
	props := sut.GetProps()

	// Assert
	assert.NotNil(t, props)
	assert.Equal(t, len(getMappedProps()), len(props), "different quantity mappeds x map returned")
}

func TestPropJSONRepository_GetProp(t *testing.T) {
	// Arrange
	mappedProps := getMappedProps()
	sut, _ := taleslabrepositories.NewPropRepository(buildPropsPath())

	// Act & Assert
	for key, _ := range mappedProps {
		assert.NotNil(t, sut.GetProp(key), fmt.Sprintf("ornament not loaded %s", key))
	}

	loadedProps := sut.GetProps()
	for key, _ := range loadedProps {
		assert.True(t, mappedProps[key], fmt.Sprintf("ornament not mapped %s", key))
	}
}

func buildPropsPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.Replace(dir, "/pkg/taleslab/taleslabrepositories", "", 1)
	return path.Join(dir, "/configs/props.json")
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
		"lava":                                true,
		"lava_dust":                           true,
		"orange_fish":                         true,
		"one_small_rock-z0":                   true,
		"one_small_rock-z1":                   true,
		"coconut":                             true,
		"dead_ground_small":                   true,
		"lava_cold":                           true,
		"water":                               true,
		"coral":                               true,
	}
}
