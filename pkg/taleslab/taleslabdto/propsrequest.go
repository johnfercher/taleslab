package taleslabdto

type PropsDtoRequest struct {
	TreeDensity  int `json:"tree_density"`
	StoneDensity int `json:"stone_density"`
	MiscDensity  int `json:"misc_density"`
}
