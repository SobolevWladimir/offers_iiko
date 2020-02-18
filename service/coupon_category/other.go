package coupon_category

func GetChilds(target *Category, cats *Categorys) Categorys {
	result := Categorys{}
	for _, div := range *cats {
		if int(div.Parent.Int64) == target.Id {
			result = append(result, GetChilds(&div, cats)...)
			result = append(result, div)
		}
	}
	return result
}
