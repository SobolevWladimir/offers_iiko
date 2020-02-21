package transport

type ACartOprionItems []ACartOprionItem
type ACartOptionItemProduct struct {
	ID      int          `json:"id"`
	Product AProductItem `json:"product"`
	Type    string       `json:"type"`
}
type ACartOprionItem struct {
	Item     ACartOptionItemProduct `json:"item"`
	Quantity int                    `json:"quantity"`
}
