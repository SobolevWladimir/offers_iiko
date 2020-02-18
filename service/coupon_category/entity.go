package coupon_category

import (
	"errors"

	"gopkg.in/guregu/null.v3"
)

type Category struct {
	Id      int      `db:"id" json:"id" valid:"-"`
	Name    string   `db:"name" json:"name" valid:"required" `
	Parent  null.Int `db:"parent" json:"parent" valid:"-" `
	City    int      `db:"city" json:"city" valid:"-"`
	Sort    int      `db:"sort" json:"sort" valid:"-"`
	Deleted bool     `db:"deleted" json:"-" valid:"-"`
}
type Categorys []Category

// Нийти категорию по id
func (cats *Categorys) GetCategoryById(id int) (Category, error) {
	for _, cat := range *cats {
		if cat.Id == id {
			return cat, nil
		}
	}
	return Category{}, errors.New("category not found")
}

func (cats *Categorys) GetCategorysByParent(id int64) Categorys {
	result := Categorys{}
	for _, c := range *cats {
		if c.Parent.Int64 == id {
			result = append(result, c)
		}
	}
	return result
}

//очищает от незвестных родителей
func (cats *Categorys) ClearNullParent() Categorys {
	result := Categorys{}
	for _, cat := range *cats {
		if _, err := cats.GetById(int(cat.Parent.Int64)); err != nil {
			cat.Parent.Int64 = 0
		}
		result = append(result, cat)
	}

	return result
}
func (cats *Categorys) GetById(id int) (Category, error) {
	for _, cat := range *cats {
		if cat.Id == id {
			return cat, nil
		}
	}
	return Category{}, errors.New("not found")
}

//возвращает идентификаторы категорий
func (cats *Categorys) GetIds() []int {
	result := []int{}
	for _, cat := range *cats {
		result = append(result, cat.Id)
	}
	return result
}
