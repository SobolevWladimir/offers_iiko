package client

import (
	"altegra_offers/service/client_category_link"

	"gopkg.in/guregu/null.v3"
)

type Client struct {
	Phone     string      `db:"phone" json:"phone" valid:"required" accessfield:"phone"`
	Name      string      `db:"name" json:"name" valid:"required" accessfield:"name"`
	LastName  null.String `db:"lastname" json:"lastname" valid:"-"`
	Email     null.String `db:"email" json:"email" valid:"nullEmail, optional"`
	Sex       int         `db:"sex" json:"sex" valid:"optional" accessfield:"sex"`
	BirthDate null.String `db:"birth_date" json:"birth_date" valid:"-"  accessfield:"birth_date"`
	Bonuses   float64     `db:"bonuses" json:"bonuses" valid:"-"`
	Deleted   bool        `db:"deleted" json:"-" valid:"-"`
}
type Clients []Client
type ClientEntity struct {
	Client
	Categories []string `json:"categories"`
}
type ClientEntitys []ClientEntity

func (c *Client) ToEntity() (ClientEntity, error) {
	result := ClientEntity{}
	result.Client = *c
	catgs, err := client_category_link.FindAllCategoriesByClient(c.Phone)
	if err != nil {
		return result, err
	}
	result.Categories = catgs
	return result, nil
}
func (c *Clients) ToEntity() (ClientEntitys, error) {
	result := ClientEntitys{}
	for _, client := range *c {
		entity, err := client.ToEntity()
		if err != nil {
			return result, err
		}
		result = append(result, entity)
	}
	return result, nil
}
func (c *ClientEntity) Save() error {
	cl := c.Client
	err := Save(&cl)
	if err != nil {
		return err
	}
	return client_category_link.SaveByClient(cl.Phone, c.Categories)
}
func (c *ClientEntity) Insert() error {
	cl := c.Client
	err := Insert(&cl)
	if err != nil {
		return err
	}
	return client_category_link.InsertByClient(cl.Phone, c.Categories)
}
