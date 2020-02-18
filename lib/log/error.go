package log

import syslog "log"

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
