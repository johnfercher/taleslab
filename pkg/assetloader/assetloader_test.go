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

	os.Chdir(AppBaseDir)
	_, _ = os.Getwd()

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
	assert.Equal(t, len(rawProps), len(props), "repeated ornaments")

	for i := 0; i < len(rawProps); i++ {
		for j := 0; j < len(rawProps); j++ {
			if i != j {
				assert.NotEqual(t, rawProps[i].Id, rawProps[j].Id, fmt.Sprintf("repeated ornaments ids, name: %s", rawProps[i].Name))
				assert.NotEqual(t, rawProps[i].Name, rawProps[j].Name, fmt.Sprintf("repeated ornaments names, id %s", rawProps[i].Id))
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
	assert.Equal(t, len(rawConstructors), len(constructors), "repeated constructors")

	for i := 0; i < len(rawConstructors); i++ {
		for j := 0; j < len(rawConstructors); j++ {
			if i != j {
				assert.NotEqual(t, rawConstructors[i].Id, rawConstructors[j].Id, fmt.Sprintf("repeated constructors ids, name: %s", rawConstructors[i].Name))
				assert.NotEqual(t, rawConstructors[i].Name, rawConstructors[j].Name, fmt.Sprintf("repeated constructors names, id %s", rawConstructors[i].Id))
			}
		}
	}
}
