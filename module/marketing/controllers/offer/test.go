package offer

import (
	"altegra_offers/mentity"
	"altegra_offers/service/offers_engine"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Actions = append(Actions, &testAction)
}

var testAction mentity.Action = mentity.Action{
	Name:   "offer_test",
	Label:  "test",
	Path:   "/test",
	Method: "get",
	Handler: func(c *gin.Context) {
		entity := offers_engine.Policy{
			Id:       0,
			Name:     "политика",
			Status:   "сcылка на  код статуса",
			Category: 1,
			SetRules: offers_engine.SetRules{
				offers_engine.SetRule{
					Description: "Описание набора правил",
					Rules: offers_engine.Rules{
						offers_engine.Rule{
							Filters: offers_engine.RuleFilters{
								offers_engine.RuleFilter{},
							},
						},
					},
				},
			},
			Actions: offers_engine.Actions{
				offers_engine.Action{},
			},
		}

		c.JSON(http.StatusOK, entity)
	},
}
