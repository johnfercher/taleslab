package taleslabrepositories

import (
	"encoding/json"
	"os"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
)

type propJSONRepository struct {
	path  string
	props map[string]*taleslabentities.Prop
}

func NewPropRepository(path ...string) (taleslabrepositories.PropRepository, error) {
	assetLoader := &propJSONRepository{}
	if len(path) != 0 {
		assetLoader.path = path[0]
	} else {
		assetLoader.path = "./configs/props.json"
	}

	err := assetLoader.loadProps()
	if err != nil {
		return nil, err
	}

	return assetLoader, nil
}

func (p *propJSONRepository) GetProp(id string) *taleslabentities.Prop {
	return p.props[id]
}

func (p *propJSONRepository) GetProps() map[string]*taleslabentities.Prop {
	return p.props
}

func (p *propJSONRepository) loadProps() error {
	bytes, err := os.ReadFile(p.path)
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
		propMap[prop.ID] = prop
	}

	p.props = propMap

	return nil
}
