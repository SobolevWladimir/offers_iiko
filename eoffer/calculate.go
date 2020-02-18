package eoffer

import (
	"altegra_offers/lib/base"
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/client"
	"altegra_offers/service/client_category_link"
	"altegra_offers/service/coupon"
	"altegra_offers/service/offers_engine"
	"altegra_offers/service/order_engine"
	"altegra_offers/service/status"
	"errors"
	"fmt"
)

//подсчет акций
func Calculate(order, previousVersionOrder *order_engine.Order) (offerentity.Allegiance, error) {
	result := offerentity.Allegiance{}
	if order == nil {
		return result, errors.New("pointer order is null")
	}

	given := offerentity.OfferGiven{
		System: System{
			Order: order,
		},
		Order:  Order(*order),
		Client: Client(order.Client.ValueOrZero()),
	}
	if previousVersionOrder != nil {
		given.OldVersionOrder = Order(*previousVersionOrder)
		//result = previousVersionOrder.Offers
	}
	tpolicys, err := getOffersForOrder(order)
	if err != nil {
		return result, err
	}
	all_statuses, err := status.FindAll()
	if err != nil {
		return result, err
	}
	stats, err := all_statuses.GetStatusesFromStatusCode(order.Status)
	if err != nil {
		return result, err
	}
	if !offers_engine.AllowCalculte(tpolicys) {
		return CalculateOnGlobalServer(order)
	}
	order_off := offerentity.RItems{}
	for _, stat := range stats {
		off, err := offers_engine.Calculate(&given, tpolicys.FilterByStatus(stat.Code))
		if err == nil {
			order_off.SetOffers(stat.Code, off)
		}
	}
	//Если заказ новый
	if previousVersionOrder == nil {
		result.Item = order_off
		order.CalculateCart(previousVersionOrder, &result.Item)
		//result.Events = result.Item.CreateEvent(Order(*order), order.Status)
		return result, nil
	}
	order_stat, _ := all_statuses.GetStatusByCode(order.Status)
	prev_stat, _ := all_statuses.GetStatusByCode(previousVersionOrder.Status)
	//если статус заказа изменился в большею сторону
	if order_stat.Priority > prev_stat.Priority {
		result.Item = previousVersionOrder.Offers.AppendRItems(&order_off)
		order.CalculateCart(previousVersionOrder, &result.Item)
		result.Events = previousVersionOrder.OfferEvent
		//result.Events.AppendEvents(result.Item.CreateEvent(Order(*order), order.Status))
		return result, nil
	}
	//если статус заказа  не изменился
	if order.Status == previousVersionOrder.Status {
		result.Item = previousVersionOrder.Offers.ReplaceRItems(&order_off)
		order.CalculateCart(previousVersionOrder, &result.Item)
		result.Events = previousVersionOrder.OfferEvent
		cev := previousVersionOrder.Offers.GetCancelEvent(previousVersionOrder.Status, previousVersionOrder.OfferEvent)
		result.Events.AppendEvents(cev)
		//result.Events.AppendEvents(result.Item.CreateEvent(Order(*order), order.Status))
		return result, nil
	}
	//если статус заказа измнился в меньшую сторону
	if order_stat.Priority < prev_stat.Priority {
		//  получаем статусы которые нужно отменить
		dstat := all_statuses.GetBeetwenByPriority(order_stat.Priority, prev_stat.Priority)
		// из действий в предыдущем заказе удаляем статусы отмены
		items := previousVersionOrder.Offers.GetWithoutStatuses(dstat.GetCodes())
		//дальше как еслибы заказ изменился в большую сторону
		result.Item = items.AppendRItems(&order_off)
		order.CalculateCart(previousVersionOrder, &result.Item)
		result.Events = previousVersionOrder.OfferEvent
		cev := previousVersionOrder.Offers.GetCancelEventByStatuses(dstat.GetCodes(), previousVersionOrder.OfferEvent)
		result.Events.AppendEvents(cev)
		//result.Events.AppendEvents(result.Item.CreateEvent(Order(*order), order.Status))
		return result, nil
	}

	//если статус заказа измениля а приоритет нет
	if order_stat.Priority == prev_stat.Priority && order_stat.Code != prev_stat.Code {
		result.Item = previousVersionOrder.Offers.ReplaceRItems(&order_off)
		order.CalculateCart(previousVersionOrder, &result.Item)
		result.Events = previousVersionOrder.OfferEvent
		cev := previousVersionOrder.Offers.GetCancelEvent(previousVersionOrder.Status, previousVersionOrder.OfferEvent)
		result.Events.AppendEvents(cev)
		//result.Events.AppendEvents(result.Item.CreateEvent(Order(*order), order.Status))
		return result, nil
	}

	return result, nil
}
func CalculateOnGlobalServer(order *order_engine.Order) (offerentity.Allegiance, error) {
	result := offerentity.Allegiance{}
	//@fixme  сделать подсчет  на главном сервере :
	return result, errors.New("Не реализованна функционал подсчета на сервере")
}

func getOffersForOrder(order *order_engine.Order) (offers_engine.Policys, error) {

	result := offers_engine.Policys{}
	if order == nil {
		return result, errors.New("pointer order is null")
	}
	// cats, err := offers_engine.GetCategorysByDivisions(divs)
	// if err != nil {
	// return result, err
	// }

	//	return offers_engine.GetOfferByCategorys(&cats)
	return result, nil
}
func ExecAction(order *order_engine.Order) error {

	if len(order.Client.ValueOrZero()) == 0 {
		return nil
	}
	//bonuses
	bals := order.OfferEvent.Bonuses.CalculateValue(false)
	if bals != 0 {
		err := client.AppendBonuses(order.Client.ValueOrZero(), bals)
		if err != nil {
			return err
		}
	}
	order.OfferEvent.Bonuses.SetCompleted(true)
	//end bonuses
	// coupon
	coup_event := order.OfferEvent.Coupons.FilterByCompleted(false)
	coup_enabled := coup_event.FilterByStatus(true)
	coup_off := coup_event.FilterByStatus(false)

	err := coupon.SaveStasusByIds(coup_off.GetCouponIds(), false, "списан по акции")
	if err != nil {
		fmt.Println("save off:", err, coup_off.GetCouponIds())
		return err
	}
	err = coupon.SaveStasusByIds(coup_enabled.GetCouponIds(), true, "")
	if err != nil {
		return err
	}

	order.OfferEvent.Coupons.SetCompleted(true)
	//end coupon

	// client category
	client_cat, err := client_category_link.FindAllCategoriesByClient(order.Client.String)
	if err != nil {
		return err
	}
	cat_event := order.OfferEvent.ClientCategory.FilterByCompleted(false)
	cat_added := cat_event.GetAddedCategory()
	cat_removed := cat_event.GetRemovedCategory()
	client_cat = appendCategorys(client_cat, cat_added)
	client_cat = removeCategory(client_cat, cat_removed)
	client_category_link.SaveByClient(order.Client.String, client_cat)
	order.OfferEvent.ClientCategory.SetCompleted(true)
	// end client category
	return order_engine.SaveOfferEvents(order)
}
func appendCategorys(target, added []string) []string {
	result := target
	for _, cat := range added {
		if !base.ContainsStringInArray(cat, target) {
			result = append(result, cat)
		}
	}
	return result
}
func removeCategory(target, cats []string) []string {
	result := []string{}
	for _, cat := range target {
		if !base.ContainsStringInArray(cat, cats) {
			result = append(result, cat)
		}
	}
	return result
}
