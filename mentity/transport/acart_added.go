package transport

type ACartAdded struct {
	DoubleMeat int              `json:"doubleMeat"`
	DoubleVeg  int              `json:"doubleVeg"`
	Noodle     int              `json:"noodle"`
	Source     int              `json:"source"`
	Toping     ACartOprionItems `json:"toping"`
}
