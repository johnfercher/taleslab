package proploader

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"io/ioutil"
)

type PropLoader interface {
	GetProp(id string) *taleslabentities.Prop
	GetProps() map[string]*taleslabentities.Prop
}

type propLoader struct {
	props map[string]*taleslabentities.Prop
}

func NewPropLoader() (*propLoader, error) {
	assetLoader := &propLoader{}

	err := assetLoader.loadProps()
	if err != nil {
		return nil, err
	}

	return assetLoader, nil
}

func (self *propLoader) GetProp(id string) *taleslabentities.Prop {
	return self.props[id]
}

func (self *propLoader) GetProps() map[string]*taleslabentities.Prop {
	return self.props
}

func (self *propLoader) loadProps() error {
	bytes, err := ioutil.ReadFile("./config/assets/props.json")
	if err != nil {
		return err
	}

	props := []*taleslabentities.Prop{}

	err = json.Unmarshal(bytes, &props)
	if err != nil {
		return err
	}

	propMap := make(map[string]*taleslabentities.Prop)

	for _, prop := range props {
		propMap[prop.Id] = prop
	}

	self.props = propMap

	return nil
}
