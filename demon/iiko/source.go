package iiko

import "time"

func (o *Object) Start() {
	nWather := NomenclatureWather{
		Interval: time.Second * 600,
	}
	nWather.Start()

}
