package offers_engine

import (
	"altegra_offers/lib/base"
	"altegra_offers/service/offers_engine/template"
	"encoding/json"
	"errors"
)

type Template struct {
	Id      string  `json:"id"`
	SetRule SetRule `json:"setrule"`
	Sort    int     `json:"sort"`
}
type Templates []Template

func toTemplate(tem *template.Template) Template {
	result := Template{}
	result.Id = tem.Id
	result.Sort = tem.Sort
	rule := SetRule{}
	json.Unmarshal([]byte(tem.SetRule), &rule)
	result.SetRule = rule
	return result
}
func toTemplates(tems *template.Templates) Templates {
	result := Templates{}
	for _, tem := range *tems {
		result = append(result, toTemplate(&tem))
	}
	return result
}
func (tem *Template) toDBEntity() (template.Template, error) {
	result := template.Template{}
	result.Id = tem.Id
	result.Sort = tem.Sort
	rule, err := json.Marshal(tem.SetRule)
	result.SetRule = string(rule)
	return result, err
}

func FindAllTemplates() (Templates, error) {
	tems, err := template.FindAll()
	if err != nil {
		return Templates{}, err
	}
	return toTemplates(&tems), nil
}
func FindTemplateById(id string) (Template, error) {
	tem, err := template.FindOneById(id)
	if err != nil {
		return Template{}, err
	}
	return toTemplate(&tem), nil
}
func (tem *Template) Insert() error {
	entity, err := tem.toDBEntity()
	if err != nil {
		return err
	}
	if len(entity.Id) == 0 {
		entity.Id = base.UUID()
	}
	return template.Insert(&entity)
}
func (tem *Template) Save() error {
	entity, err := tem.toDBEntity()
	if len(entity.Id) == 0 {
		return errors.New("id is empty")
	}
	if err != nil {
		return err
	}
	return template.Save(&entity)
}
func RemoveTemplateById(id string) error {
	return template.RemoveById(id)
}
