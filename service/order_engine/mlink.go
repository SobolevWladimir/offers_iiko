package order_engine

import (
	"altegra_offers/service/order_engine/mlink"

	"github.com/jmoiron/sqlx"
)

func SaveMLink(mlinks []string, order string) error {
	if err := mlink.RemoveNotExist(mlinks, order); err != nil {
		return err
	}
	return mlink.SaveMlinks(mlinks, order)
}
func TxSaveMLink(tx *sqlx.Tx, mlinks []string, order string) error {
	if err := mlink.TxRemoveNotExist(tx, mlinks, order); err != nil {
		return err
	}
	return mlink.TxSaveMlinks(tx, mlinks, order)
}
