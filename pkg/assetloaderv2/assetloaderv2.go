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
	GetConstructors() ([]AssetInfo, error)
	GetOrnaments() ([]AssetInfo, error)
}

type assetLoaderV2 struct {
}

func NewAssetLoaderV2() *assetLoaderV2 {
	return &assetLoaderV2{}
}

func (self *assetLoaderV2) GetConstructors() ([]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/constructors_v2.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	return assetInfos, nil
}

func (self *assetLoaderV2) GetOrnaments() ([]AssetInfo, error) {
	bytes, err := ioutil.ReadFile("./config/assets/ornaments_v2.json")
	if err != nil {
		return nil, err
	}

	assetInfos := []AssetInfo{}

	err = json.Unmarshal(bytes, &assetInfos)
	if err != nil {
		return nil, err
	}

	return assetInfos, nil
}
