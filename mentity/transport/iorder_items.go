package transport

import "errors"

type IOrderItems []IOrderItem

func (p *IOrderItems) GetSiteIdById(id string) (int, error) {
	for _, item := range *p {
		if item.ID == id {
			return item.SiteId, nil
		}
	}
	return 0, errors.New(" не могу найти id сайта ")
}
func (p *IOrderItems) GetSiteIdByOrderItemId(id string) (int, error) {
	for _, item := range *p {
		if item.OrderItemId == id {
			return item.SiteId, nil
		}
	}
	return 0, errors.New(" не могу найти id сайта ")
}
