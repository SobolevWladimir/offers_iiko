package webhook

import "altegra_offers/lib/tablename"

func GetUrlByAliasInDB(alias string) (string, error) {
	var result string
	db := connect()

	err := db.Get(&result, `SELECT `+tablename.Webhook+`.url FROM `+tablename.Webhook+`, `+tablename.Webhook_type+` where `+tablename.Webhook+`.type=`+tablename.Webhook_type+`.id AND `+tablename.Webhook_type+`.alias='`+alias+`'`)
	if err != nil {
		return err.Error(), err

	}
	return result, nil

}
