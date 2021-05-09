package taleslabrepositories

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"io/ioutil"
)

type propJsonRepository struct {
	props map[string]*taleslabentities.Prop
}

func NewPropRepository() (*propJsonRepository, error) {
	assetLoader := &propJsonRepository{}

	err := assetLoader.loadProps()
	if err != nil {
		return nil, err
	}

	return assetLoader, nil
}

func (self *propJsonRepository) GetProp(id string) *taleslabentities.Prop {
	return self.props[id]
}

func (self *propJsonRepository) GetProps() map[string]*taleslabentities.Prop {
	return self.props
}

func (self *propJsonRepository) loadProps() error {
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
