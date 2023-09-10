package grid

import "github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"

type River struct {
	Start              *taleslabentities.Vector3d
	End                *taleslabentities.Vector3d
	HeightCutThreshold int
}
