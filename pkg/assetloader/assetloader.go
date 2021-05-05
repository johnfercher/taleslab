package assetloader

import (
	"encoding/json"
	"io/ioutil"
)

type AssetInfo struct {
	Id          string       `json:"id"`
	AssertParts []*AssetPart `json:"asset_parts"`
	Type        string       `json:"type"`
}

type AssetPart struct {
	Id         []byte      `json:"id"`
	Dimensions *Dimensions `json:"dimensions"`
	OffsetX    int         `json:"offset_x"`
	OffsetY    int         `json:"offset_y"`
	OffsetZ    int         `json:"offset_z"`
	Rotation   int         `json:"rotation"`
	Name       string      `json:"name"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Length int `json:"length"`
	Height int `json:"height"`
}

type AssetLoader interface {
	GetConstructor(id string) *AssetInfo
	GetProp(id string) *AssetInfo
	GetConstructors() map[string]*AssetInfo
	GetProps() map[string]*AssetInfo
}

type assetLoader struct {
	constructors map[string]*AssetInfo
	props        map[string]*AssetInfo
}

func NewAssetLoader() (*assetLoader, error) {
	assetLoader := &assetLoader{}

	err := assetLoader.loadProps()
	if err != nil {
		return nil, err
	}

	err = assetLoader.loadConstructors()
	if err != nil {
		return nil, err
	}

	return assetLoader, nil
}

func (self *assetLoader) GetConstructor(id string) *AssetInfo {
	return self.constructors[id]
}

func (self *assetLoader) GetConstructors() map[string]*AssetInfo {
	return self.constructors
}

func (self *assetLoader) GetProp(id string) *AssetInfo {
	return self.props[id]
}

func (self *assetLoader) GetProps() map[string]*AssetInfo {
	return self.props
}

func (self *assetLoader) loadConstructors() error {
	bytes, err := ioutil.ReadFile("./config/assets/constructors.json")
	if err != nil {
		return err
	}

	assetInfos := []*AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "constructors"
	}

	assetMap := make(map[string]*AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Id] = assetinfo
	}

	self.constructors = assetMap

	return nil
}

func (self *assetLoader) loadProps() error {
	bytes, err := ioutil.ReadFile("./config/assets/ornaments.json")
	if err != nil {
		return err
	}

	assetInfos := []*AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "ornaments"
	}

	assetMap := make(map[string]*AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Id] = assetinfo
	}

	self.props = assetMap

	return nil
}
