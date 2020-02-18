package webhook

import (
	"altegra_offers/service/webhook_type"
	"errors"
)

var buffer = make(map[string]string)

func UpdateBuffer() {

	aliases, _ := webhook_type.FindAll()

	for _, a := range aliases {
		setUrlInBufer(a.Alias)
	}
}

func FindAlias(alias string) (bool, error) {

	if _, ok := buffer[alias]; !ok {
		err := setUrlInBufer(alias)
		if err != nil {
			return false, err
		}
	}
	return true, nil

}

func setUrlInBufer(alias string) error {
	url, err := GetUrlByAliasInDB(alias)
	if err != nil {
		return err
	}
	buffer[alias] = url
	return nil
}

func GetUrlByAlias(alias string) (string, error) {
	if ok, err := FindAlias(alias); ok {
		if err != nil {
			return "", err
		}
		return buffer[alias], nil
	}
	return "", errors.New("Alias не найден")

}
