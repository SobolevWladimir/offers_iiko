package offerentity

type ActionBonuses struct {
	All       bool     `json:"all"`
	Exception []string `json:"exception"`
	Type      int      `json:"type"`
	Value     float32  `json:"value"`
}
