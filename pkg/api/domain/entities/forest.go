package entities

type Forest struct {
	Ground    *Ground    `json:"ground,omitempty"`
	Mountains *Mountains `json:"mountains,omitempty"`
	River     *River     `json:"river,omitempty"`
	Props     *Props     `json:"props,omitempty"`
}
