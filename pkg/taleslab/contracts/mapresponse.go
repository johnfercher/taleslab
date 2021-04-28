package contracts

// MapResponse response model
type MapResponse struct {
	// Version of the TaleSpire Slab
	SlabVersion string `json:"slab_version"`
	// Size of the base64 string
	Size string `json:"size"`
	// Code to insert in the game
	Code string `json:"code"`
}

// swagger:response mapRes
type swaggMapRes struct {
	// in: body
	Map MapResponse
}
