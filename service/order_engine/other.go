package order_engine

import (
	"altegra_offers/lib/base"
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/coupon"
	"altegra_offers/service/coupon_category"
	"altegra_offers/service/order_engine/orders"
	"altegra_offers/service/order_engine/products"
	"fmt"
	"strconv"

	"time"

	"github.com/huandu/go-sqlbuilder"
)

func GetClietnProductCount(client string, filters offerentity.OfferFilterValues) (int, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	t_order := orders.GetTableName()
	t_product := products.GetTableName()
	sb.Select(fmt.Sprintf(`count("%v"."id")`, t_order))
	sb.From(t_order)
	sb.Join(t_product, fmt.Sprintf(`%v.id=%v.order`, t_order, t_product))
	params := []string{}
	for _, fil := range filters {
		switch fil.Field {
		case "date_create":
			value, err := time.Parse("02-01-2006", fil.Value)
			if err != nil {
				return 0, err
			}
			params = append(params, sb.And(fmt.Sprintf(`%s.created%v%s`, sqlbuilder.Escape(t_order), fil.Operator, sb.Args.Add(value))))

		}

	}
	params = append(params, sb.Equal(t_order+".client", client))
	sb.Where(params...)
	sql, args := sb.Build()
	db := connect()
	row := db.QueryRow(sql, args...)
	var res int
	err := row.Scan(&res)
	return res, err
}
func GetLastActivityClient(client string, filters offerentity.OfferFilterValues) (time.Time, error) {

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("MAX(last_update)")
	sb.From(orders.GetTableName())
	sb.Where(sb.Equal("client", client))
	sql, args := sb.Build()
	db := connect()
	row := db.QueryRow(sql, args...)
	var res time.Time
	err := row.Scan(&res)
	return res, err
}
func GetCountOrdersClient(client string, filters offerentity.OfferFilterValues) (int, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select(`COUNT("id")`)
	sb.From(orders.GetTableName())
	t_order := orders.GetTableName()
	params := []string{}
	for _, fil := range filters {
		switch fil.Field {
		case "date_create":
			{
				value, err := time.Parse("02-01-2006", fil.Value)
				if err != nil {
					return 0, err
				}
				params = append(params, sb.And(fmt.Sprintf(`%s.created%v%s`, sqlbuilder.Escape(t_order), fil.Operator, sb.Args.Add(value))))
			}
		case "coupon":
			params = append(params, sb.And(fmt.Sprintf(`%s.coupon%v%s`, sqlbuilder.Escape(t_order), fil.Operator, sb.Args.Add(fil.Value))))
		case "coupon_category":
			{
				cats, err := coupon_category.FindAll()
				if err != nil {
					return 0, err
				}
				fil_value, _ := strconv.Atoi(fil.Value)
				tcat, err := cats.GetById(fil_value)
				if err != nil {
					return 0, err
				}
				child := coupon_category.GetChilds(&tcat, &cats)
				child = append(child, tcat)

				coups, err := coupon.FindByCategorys(child.GetIds())
				if err != nil {
					return 0, err
				}
				names := coups.GetNames()
				params = append(params, sb.In("coupon", base.StringToInterface(names)...))
			}
		default:
			fmt.Println(" order_engine/other.go  there is no implementation for the field:", fil.Field)
		}

	}
	params = append(params, sb.Equal(t_order+".client", client))
	sb.Where(params...)
	sql, args := sb.Build()
	db := connect()
	row := db.QueryRow(sql, args...)
	var res int
	err := row.Scan(&res)
	return res, err
}
