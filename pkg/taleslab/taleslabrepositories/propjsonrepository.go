package taleslabrepositories

import (
	"encoding/json"
	"os"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
)

type propJSONRepository struct {
	props map[string]*taleslabentities.Prop
}

func NewPropRepository() taleslabrepositories.PropRepository {
	assetLoader := &propJSONRepository{}

	_ = assetLoader.loadProps()
	return assetLoader
}

func (p *propJSONRepository) GetProp(id string) *taleslabentities.Prop {
	return p.props[id]
}

func (p *propJSONRepository) GetProps() map[string]*taleslabentities.Prop {
	return p.props
}

func (p *propJSONRepository) loadProps() error {
	bytes, err := os.ReadFile("./configs/props.json")
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
