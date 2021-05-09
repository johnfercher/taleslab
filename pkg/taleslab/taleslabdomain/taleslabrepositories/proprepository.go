package taleslabrepositories

import "github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"

type PropRepository interface {
	GetProp(id string) *taleslabentities.Prop
	GetProps() map[string]*taleslabentities.Prop
}
