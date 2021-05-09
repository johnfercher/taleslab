package taleslabdto

// GenerationCountDtoResponse request model
// swagger:model
type GenerationCountDtoResponse struct {
	Count uint64 `json:"count"`
}

// Response from API
// swagger:response swaggCountRes
// nolint:deadcode,unused
type swaggCountRes struct {
	// in: body
	Count GenerationCountDtoResponse
}
