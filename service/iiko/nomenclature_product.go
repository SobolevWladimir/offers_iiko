package iiko

type Product struct {
	Code              string `json:"code"`
	Name              string `json:"name"`
	ProductCategoryId string `json:"productCategoryId"`
}
type Products []Product
