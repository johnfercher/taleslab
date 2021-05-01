package assetloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AssetInfo struct {
	Id         []byte      `json:"id"`
	Dimensions *Dimensions `json:"dimensions"`
	OffsetZ    int         `json:"offset_z"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Length int `json:"length"`
	Height int `json:"height"`
}

type AssetLoader interface {
	GetConstructor(id string) AssetInfo
	GetProp(id string) AssetInfo
}

type assetLoader struct {
	constructors map[string]AssetInfo
	props        map[string]AssetInfo
}

func NewAssetLoader() *assetLoader {
	assetLoader := &assetLoader{}
	assetLoader.loadProps()
	assetLoader.loadConstructors()
	return assetLoader
}

func (self *assetLoader) GetConstructor(id string) AssetInfo {
	return self.constructors[id]
}

func (self *assetLoader) GetProp(id string) AssetInfo {
	return self.props[id]
}

func (self *assetLoader) loadConstructors() {
	bytes, err := ioutil.ReadFile("./config/assets/constructors.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "constructors"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	self.constructors = assetMap
}

func (self *assetLoader) loadProps() {
	bytes, err := ioutil.ReadFile("./config/assets/ornaments.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "ornaments"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	self.props = assetMap
}
