package transport

type ProductItem struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Weight  string `json:"weight"`
	New     bool   `json:"new"`
	Hot     bool   `json:"hot"`
	Hit     bool   `json:"hit"`
	Caloric string `json:"caloric"`
	Vendor1 string `json:"vendor1"`
	Vendor2 string `json:"vendor2"`
}
