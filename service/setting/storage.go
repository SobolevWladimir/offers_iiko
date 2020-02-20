package setting

import "fmt"

var iiko_settings IikoSettings

func UpdataStorage() error {
	entitys, err := FindAll()
	if err != nil {
		return err
	}
	iiko_settings = entitys
	return nil
}

func GetSettingIikoBySity(id int) (IikoSetting, error) {
	for _, item := range iiko_settings {
		if item.SityId == id {
			return item, nil
		}

	}
	return IikoSetting{}, fmt.Errorf(" не могу найти настройки iiko  для города(id: %v)", id)
}
