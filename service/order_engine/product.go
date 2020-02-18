package order_engine

import (
	"altegra_offers/lib/base"
	"altegra_offers/service/order_engine/products"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
)

type Product struct {
	Id        string      `json:"id"`
	Product   string      `json:"product"`
	Quantity  float32     `json:"quantity"`
	Comment   null.String `json:"comment"`
	Modifiers Modifiers   `json:"modifiers"`
}
type Products []Product

func toProduct(pr products.Product) (Product, error) {
	result := Product{}
	result.Id = pr.Id
	result.Product = pr.Product
	result.Quantity = pr.Quantity
	result.Comment = pr.Comment
	mods, err := FindAllModifierByProduct(pr.Id)
	result.Modifiers = mods
	return result, err
}
func (pr *Product) toDBEntity(order string) products.Product {
	result := products.Product{}
	result.Id = pr.Id
	result.Order = order
	result.Product = pr.Product
	result.Quantity = pr.Quantity
	result.Comment = pr.Comment
	return result
}
func toProducts(prs products.Products) (Products, error) {
	result := Products{}
	for _, pr := range prs {
		if prod, err := toProduct(pr); err != nil {
			return result, err
		} else {
			result = append(result, prod)
		}
	}
	return result, nil
}
func (prs *Products) toDBEntity(order string) products.Products {
	result := products.Products{}
	for _, pr := range *prs {
		result = append(result, pr.toDBEntity(order))
	}
	return result
}

func FindAllProductByOrder(order string) (Products, error) {
	result := Products{}
	prods, err := products.FindAllByOrder(order)
	if err != nil {
		return result, err
	}
	return toProducts(prods)
}
func (prs *Products) Insert(order string) error {
	for _, ps := range *prs {
		entity := ps.toDBEntity(order)
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		if err := products.Insert(&entity); err != nil {
			return err
		}
		if err := ps.Modifiers.Insert(entity.Id); err != nil {
			return err
		}
	}
	return nil
}
func (prs *Products) Save(order string) error {
	prs.CheckUUID()
	prods := prs.toDBEntity(order)
	if err := products.RemoveNotExist(&prods, order); err != nil {
		return err
	}
	if err := products.SaveExist(&prods); err != nil {
		return err
	}
	for _, pr := range *prs {
		if err := pr.Modifiers.Save(pr.Id); err != nil {
			return fmt.Errorf("(modifiers) %v", err.Error())
		}
	}
	return nil
}
func (prs *Products) TxInsert(tx *sqlx.Tx, order string) error {
	for _, ps := range *prs {
		entity := ps.toDBEntity(order)
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		if err := products.TxInsert(tx, &entity); err != nil {
			return err
		}
		if err := ps.Modifiers.TxInsert(tx, entity.Id); err != nil {
			return err
		}
	}
	return nil
}
func (prs *Products) TxSave(tx *sqlx.Tx, order string) error {
	prs.CheckUUID()
	prods := prs.toDBEntity(order)
	if err := products.TxRemoveNotExist(tx, &prods, order); err != nil {
		return err
	}
	if err := products.TxSaveExist(tx, &prods); err != nil {
		return err
	}
	return nil
}
func (prs *Products) TxSaveModifiers(tx *sqlx.Tx) error {
	for _, pr := range *prs {
		if err := pr.Modifiers.TxSave(tx, pr.Id); err != nil {
			return fmt.Errorf("(modifiers) %v", err.Error())
		}
	}
	return nil
}

func (prs *Products) GetIds() []string {
	result := []string{}
	for _, p := range *prs {
		result = append(result, p.Product)
	}
	return result
}
func (prs *Products) GetBesides(ids []string) Products {
	result := Products{}
	for _, p := range *prs {
		if !contains(ids, p.Product) {
			result = append(result, p)
		}
	}
	return result
}
func (prs *Products) GetModifierIds() []string {
	result := []string{}
	for _, p := range *prs {
		result = append(result, p.Modifiers.GetModifierIds()...)
	}
	return result
}
func (prs *Products) CheckUUID() {
	for index, p := range *prs {
		if len(p.Id) == 0 {
			p.Id = base.UUID()
			(*prs)[index] = p
		}
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
