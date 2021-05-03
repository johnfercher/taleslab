package assetloader

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

var AppBaseDir = ""
var rawProps []AssetInfo
var rawConstructors []AssetInfo

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

	ornamentBytes, err := ioutil.ReadFile("./config/assets/ornaments.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(ornamentBytes, &rawProps)
	if err != nil {
		log.Fatal(err.Error())
	}

	constructorsBytes, err := ioutil.ReadFile("./config/assets/constructors.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(constructorsBytes, &rawConstructors)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestNewAssetLoader(t *testing.T) {
	// Act
	sut, err := NewAssetLoader()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, sut)
	assert.Equal(t, "*assetloader.assetLoader", fmt.Sprintf("%T", sut))
}

func TestAssetLoader_GetProps(t *testing.T) {
	// Arrange
	sut, _ := NewAssetLoader()

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

func TestAssetLoader_GetConstructors(t *testing.T) {
	// Arrange
	sut, _ := NewAssetLoader()

	// Act
	constructors := sut.GetConstructors()

	// Assert
	assert.NotNil(t, constructors)
	assert.Equal(t, len(rawConstructors), len(constructors), "different quantity array loaded x map returned")
	assert.Equal(t, len(getMappedConstructors()), len(constructors), "different quantity mappeds x map returned")

	for i := 0; i < len(rawConstructors); i++ {
		for j := 0; j < len(rawConstructors); j++ {
			if i != j {
				assert.NotEqual(t, rawConstructors[i].Id, rawConstructors[j].Id, fmt.Sprintf("repeated constructors ids, name: %s", rawConstructors[i].Id))
				assert.NotEqual(t, rawConstructors[i].Id, rawConstructors[j].Id, fmt.Sprintf("repeated constructors names, id %s", rawConstructors[i].Id))
			}
		}
	}
}

func TestAssetLoader_GetProp(t *testing.T) {
	// Arrange
	mappedProps := getMappedProps()
	sut, _ := NewAssetLoader()

	// Act & Assert
	for mappedPropKey := range mappedProps {
		assert.NotNil(t, sut.GetProp(mappedPropKey), fmt.Sprintf("ornament not loaded %s", mappedPropKey))
	}

	for loadedPropKey := range sut.GetProps() {
		assert.NotNil(t, mappedProps[loadedPropKey], fmt.Sprintf("ornament not mapped %s", loadedPropKey))
	}
}

func TestAssetLoader_GetConstructor(t *testing.T) {
	// Arrange
	mappedConstructors := getMappedConstructors()
	sut, _ := NewAssetLoader()

	// Act & Assert
	for mappedConstructorKey := range mappedConstructors {
		assert.NotNil(t, sut.GetConstructor(mappedConstructorKey), fmt.Sprintf("constructor not loaded %s", mappedConstructorKey))
	}

	for loadConstructorKey := range sut.GetConstructors() {
		assert.NotNil(t, mappedConstructors[loadConstructorKey], fmt.Sprintf("constructor not mapped %s", loadConstructorKey))
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
	}
}

func getMappedConstructors() map[string]bool {
	return map[string]bool{
		"ground_nature_big":             true,
		"ground_nature_with_stones_big": true,
		"ground_nature_small":           true,
		"ground_sand_small":             true,
		"ground_snow_small":             true,
		"mud_small":                     true,
		"mud_with_feather_small":        true,
		"clay_with_feather_small":       true,
	}
}
