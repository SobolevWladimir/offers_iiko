package transport

type ACartOprionItems []ACartOprionItem
type ACartOprionItem struct {
	Item     ProductItem `json:"item"`
	Quantity int         `json:"quantity"`
}
