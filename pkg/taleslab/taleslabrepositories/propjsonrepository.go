package taleslabrepositories

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"io/ioutil"
)

type propJsonRepository struct {
	props map[string]*taleslabentities.Prop
}

func NewPropRepository() taleslabrepositories.PropRepository {
	assetLoader := &propJsonRepository{}

	_ = assetLoader.loadProps()
	return assetLoader
}

func (self *propJsonRepository) GetProp(id string) *taleslabentities.Prop {
	return self.props[id]
}

func (self *propJsonRepository) GetProps() map[string]*taleslabentities.Prop {
	return self.props
}

func (self *propJsonRepository) loadProps() error {
	bytes, err := ioutil.ReadFile("./docs/configs/props.json")
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
