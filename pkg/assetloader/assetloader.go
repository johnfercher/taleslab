package assetloader

import (
	"encoding/json"
	"io/ioutil"
)

type AssetInfo struct {
	Id         []byte      `json:"id"`
	Dimensions *Dimensions `json:"dimensions"`
	OffsetZ    uint8       `json:"offset_z"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
}

type Dimensions struct {
	Width  uint8 `json:"width"`
	Length uint8 `json:"length"`
	Height uint8 `json:"height"`
}

type AssetLoader interface {
	GetConstructors() (map[string]AssetInfo, error)
	GetProps() (map[string]AssetInfo, error)
}

type assetLoader struct {
}

func NewAssetLoader() *assetLoader {
	return &assetLoader{}
}

func (self *assetLoader) GetConstructors() (map[string]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/constructors.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "constructors"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	return assetMap, nil
}

func (self *assetLoader) GetProps() (map[string]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/ornaments.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "ornaments"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	return assetMap, nil
}
