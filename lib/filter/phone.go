package filter

import (
	"regexp"
	"strings"
)

//
func Phone(phone string) string {
	var cellphone string
	re := regexp.MustCompile("[+,0-9]+")
	str := re.FindAllString(phone, -1)
	for _, i := range str {
		cellphone += i
	}
	return strings.Replace(cellphone, " ", "", -1)
}
