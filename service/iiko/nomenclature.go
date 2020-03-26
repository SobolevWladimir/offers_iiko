package iiko

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const URL_Nomenclature = "/api/0/nomenclature/"

type Nomenclature struct {
	ProductCategories ProductCategorys `json:"productCategories"`
	Products          Products         `json:"products"`
}

func LoadNomenclature(auth AuthData) (Nomenclature, error) {
	result := Nomenclature{}
	token, err := GetToken(auth)
	if err != nil {
		return result, err
	}
	client := &http.Client{
		Timeout: 6 * time.Second,
	}
	vals := url.Values{}
	vals.Add("access_token", token)
	url := url.URL{
		Scheme:   BizScheme,
		Host:     BizHost,
		Path:     URL_Nomenclature + auth.Organization,
		RawQuery: vals.Encode(),
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return result, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(robots, &result)

	return result, err
}
func (n *Nomenclature) FindCategoryNameByProductCode(code string) (string, error) {
	prod, err := n.Products.FindProductByCode(code)
	if err != nil {
		return "", err
	}
	cat, err := n.ProductCategories.FindByID(prod.ProductCategoryId)

	return cat.Name, err
}
func (n *Nomenclature) FindProductIdByCode(code string) (string, error) {

	prod, err := n.Products.FindProductByCode(code)
	return prod.Id, err

}
