package log

import (
	"fmt"
	"io/ioutil"
	syslog "log"
	"net/url"
	"path"
	"time"
)

var PATH = ""

func ServiceError(service string, err string) {
	println(service, err)
}
func AccessError(err string) {
	println("shaker/access", err)
}
func AccessFatal(err string) {
	syslog.Fatal(err)
}
func ControllerError(modul string, controller string, err string) {
	syslog.Println(modul+"/"+controller+": ", err)
}
func MEntityError(err error) {
	syslog.Println(err.Error())
}
func LogIIKOReques(data []byte) {
	if len(PATH) == 0 {
		return
	}
	ur, err := url.Parse(PATH)
	if err != nil {
		fmt.Println("error parse dump path")
		return
	}
	t := time.Now()
	filename := "offers_iiko_request" + t.Format("2006_01_02_15_04_05") + ".json"
	ur.Path = path.Join(ur.Path, filename)
	ioutil.WriteFile(ur.String(), data, 777)
}
func LogIIKOResponce(data []byte) {
	if len(PATH) == 0 {
		return
	}
	ur, err := url.Parse(PATH)
	if err != nil {
		fmt.Println("error parse dump path")
		return
	}
	t := time.Now()
	filename := "offers_iiko_responce" + t.Format("2006_01_02_15_04_05") + ".json"
	ur.Path = path.Join(ur.Path, filename)
	ioutil.WriteFile(ur.String(), data, 777)

}
