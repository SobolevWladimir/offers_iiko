package order_engine

import (
	"altegra_offers/service/order_engine/pmodifiers"

	"github.com/jmoiron/sqlx"
)

type Modifier struct {
	Id       string  `json:"id"`
	Modifier string  `json:"modifier"`
	Quantity float32 `json:"quantity"`
}
type Modifiers []Modifier

func toModifier(m pmodifiers.Modifier) Modifier {
	result := Modifier{}
	result.Id = m.Id
	result.Modifier = m.Modifier
	result.Quantity = m.Quantity
	return result
}
func (m *Modifier) toDBEntity(product string) pmodifiers.Modifier {
	result := pmodifiers.Modifier{}
	result.Id = m.Id
	result.Modifier = m.Modifier
	result.Quantity = m.Quantity
	result.Product = product
	return result
}
func (ms *Modifiers) toDBEntity(product string) pmodifiers.Modifiers {
	result := pmodifiers.Modifiers{}
	for _, mod := range *ms {
		result = append(result, mod.toDBEntity(product))
	}
	return result
}
func toModifiers(ms *pmodifiers.Modifiers) Modifiers {
	result := Modifiers{}
	for _, m := range *ms {
		result = append(result, toModifier(m))
	}
	return result
}
func FindAllModifierByProduct(product string) (Modifiers, error) {
	mods, err := pmodifiers.FindAllByProduct(product)
	return toModifiers(&mods), err
}
func (ms *Modifiers) Insert(product string) error {
	entitys := ms.toDBEntity(product)
	return pmodifiers.InsertArray(&entitys)
}
func (ms *Modifiers) Save(product string) error {
	entitys := ms.toDBEntity(product)
	if err := pmodifiers.RemoveNotExist(&entitys, product); err != nil {
		return err
	}
	return pmodifiers.SaveExist(&entitys)
}
func (ms *Modifiers) TxInsert(tx *sqlx.Tx, product string) error {
	entitys := ms.toDBEntity(product)
	return pmodifiers.TxInsertArray(tx, &entitys)
}
func (ms *Modifiers) TxSave(tx *sqlx.Tx, product string) error {
	entitys := ms.toDBEntity(product)
	if err := pmodifiers.TxRemoveNotExist(tx, &entitys, product); err != nil {
		return err
	}
	return pmodifiers.SaveExist(&entitys)
}

//получить список идентификаторов модификаторов
func (m *Modifiers) GetModifierIds() []string {
	result := []string{}
	for _, mod := range *m {
		result = append(result, mod.Modifier)
	}
	return result
}
