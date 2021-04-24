package assetloaderv2

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

type AssetLoaderV2 interface {
	GetConstructors() (map[string]AssetInfo, error)
	GetOrnaments() (map[string]AssetInfo, error)
}

type assetLoaderV2 struct {
}

func NewAssetLoaderV2() *assetLoaderV2 {
	return &assetLoaderV2{}
}

func (self *assetLoaderV2) GetConstructors() (map[string]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/constructors_v2.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "constructors_v2"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	return assetMap, nil
}

func (self *assetLoaderV2) GetOrnaments() (map[string]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/ornaments_v2.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(assetInfos); i++ {
		assetInfos[i].Type = "ornaments_v2"
	}

	assetMap := make(map[string]AssetInfo)

	for _, assetinfo := range assetInfos {
		assetMap[assetinfo.Name] = assetinfo
	}

	return assetMap, nil
}
