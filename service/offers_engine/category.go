package offers_engine

import (
	"altegra_offers/service/offers_engine/category"
	"errors"

	"gopkg.in/guregu/null.v3"
)

type CategoryEntity struct {
	Id       string      `json:"id" valid:"uuid"`
	Name     string      `json:"name" valid:"-"`
	Parent   null.String `json:"parent" valid:"nullUuid, optional"`
	Sort     int         `json:"sort" valid:"-"`
	Division string      `json:"division" valid:"uuid, required"`
	Policys  Policys     `json:"policys" valid:"-"`
}
type Category category.Category
type Categorys category.Categorys

func FindAllCategory() (Categorys, error) {
	cat, err := category.FindAll()
	return Categorys(cat), err
}
func FindByCity(city int) (Categorys, error) {
	cat, err := category.FindByCity(city)
	return Categorys(cat), err

}
func InsertCategory(cat *Category) error {
	entity := category.Category(*cat)
	return category.Insert(&entity)
}
func FindOneCategoryById(id string) (Category, error) {
	cat, err := category.FindOneById(id)
	return Category(cat), err
}
func SaveCategory(cat *Category) error {
	entity := category.Category(*cat)
	return category.Save(&entity)
}
func RemoveCategoryById(id int) {
	category.RemoveById(id)
}
func (cs *Categorys) GetCategoryById(id int) (Category, error) {
	result := Category{}
	for _, c := range *cs {
		if c.Id == id {
			return Category(c), nil
		}
	}
	return result, errors.New("not found category")
}
func (cats *Categorys) Append(c ...*Category) {
	for _, cat := range c {
		*cats = append(*cats, category.Category(*cat))
	}
}
func (cats *Categorys) AppendCategorys(cs *Categorys) {
	for _, cat := range *cs {
		*cats = append(*cats, cat)
	}
}
func (cats *Categorys) GetChilds(target *Category) Categorys {
	result := Categorys{}
	for _, div := range *cats {
		if int(div.Parent.Int64) == target.Id {
			cat := Category(div)
			result = append(result, cats.GetChilds(&cat)...)
			result = append(result, div)
		}
	}
	return result
}

func (cats *Categorys) FilterByCity(id int) Categorys {
	result := Categorys{}
	for _, cat := range *cats {
		if cat.City == id {
			result = append(result, cat)
		}
	}
	return result
}

//Убираем  родителя у категории если он неизвестен
func (cats *Categorys) ClearNullParent() Categorys {

	result := Categorys{}
	for _, cat := range *cats {
		if _, err := cats.GetById(int(cat.Parent.Int64)); err != nil {
			cat.Parent.Valid = false
		}
		result = append(result, cat)
	}

	return result
}
func (cats *Categorys) GetById(id int) (Category, error) {
	for _, cat := range *cats {
		if cat.Id == id {
			return Category(cat), nil
		}
	}
	return Category{}, errors.New("not found")
}
func (cats *Categorys) GetIds() []int {
	result := []int{}
	for _, cat := range *cats {
		result = append(result, cat.Id)
	}
	return result
}
