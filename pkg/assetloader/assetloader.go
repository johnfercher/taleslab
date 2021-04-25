package assetloader

import (
	"encoding/json"
	"io/ioutil"
)

type AssetInfo struct {
	Id        []byte `json:"id"`
	Dimension string `json:"dimension"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type AssetLoader interface {
	GetConstructors() (map[string]AssetInfo, error)
	GetOrnaments() (map[string]AssetInfo, error)
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

func (self *assetLoader) GetOrnaments() (map[string]AssetInfo, error) {
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
