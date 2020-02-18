package client_category_link

type ClientCategoryLink struct {
	Client   string `db:"client"`
	Category string `db:"category"`
}
type ClientCategoryLinks []ClientCategoryLink

func (cs *ClientCategoryLinks) GetCategories() []string {
	result := []string{}
	for _, c := range *cs {
		result = append(result, c.Category)
	}
	return result
}
