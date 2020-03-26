package iiko

import (
	"fmt"
	iiko_manager "offers_iiko/service/iiko"
	"offers_iiko/service/setting"
	"time"
)

type NomenclatureWather struct {
	Interval time.Duration
}

func (n *NomenclatureWather) Start() {
	ticker := time.NewTicker(n.Interval)
	n.UpdataStorage()
	go func() {
		for range ticker.C {
			n.UpdataStorage()
		}
	}()

}
func (n *NomenclatureWather) UpdataStorage() {
	set := setting.GetAllSettingIiko()
	err := iiko_manager.UpdataStorageNomenclature(conventIikoSettingToMap(set))
	if err != nil {
		fmt.Println("Ошибка загрузки номенклатуры  iiko", time.Now().Format("2006_01_02_15_04_05"))
	}

}

func conventIikoSettingToMap(set setting.IikoSettings) map[int]iiko_manager.AuthData {
	result := make(map[int]iiko_manager.AuthData)
	for _, s := range set {
		if s.UserID.IsZero() || s.UserSecret.IsZero() {
			continue
		}
		if s.Organization.IsZero() {
			fmt.Println("iiko demon:  поле организация пустое для города ", s.SityId)
			continue
		}
		result[s.SityId] = iiko_manager.AuthData{
			UserId:       s.UserID.ValueOrZero(),
			UserSecret:   s.UserSecret.ValueOrZero(),
			Organization: s.Organization.ValueOrZero(),
		}

	}
	return result
}
