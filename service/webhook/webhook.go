package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"

	"altegra_offers/config"
	"altegra_offers/lib/log"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

var type_setting config.ConfigWebhook

func Send(alias string, data interface{}) {

	url, err := GetUrlByAlias(alias)
	if err != nil {
		log.ServiceError("webhook/Send/GetUrlByAlias ", err.Error())

	}
	if url != "" {
		go SendByUrl(url, data)
	}

}
func SendByUrl(url string, data interface{}) {

	send, interval := getIntervalAndSend()
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.ServiceError("webhook/Send/SendByUrl", err.Error())

	}

	responseCode, err := reSend(url, bytes.NewBufferString(string(jsonData)))

	if err != nil {

		log.ServiceError("webhook/Send/reSend ( webhook.go 47)", err.Error())
	}

	if responseCode == http.StatusOK {
		return
	}

	for i := 0; i < send; {

		timer := time.NewTimer(time.Minute * time.Duration(interval))
		select {
		case <-timer.C:
			i++
			responseCode, err := reSend(url, bytes.NewBufferString(string(jsonData)))
			if err != nil {
				log.ServiceError("webhook/Send/reSend (webhook.go 62)", err.Error())
			}
			if responseCode == http.StatusOK {
				i = send
				return
			}
		}

	}

}
func reSend(url string, body io.Reader) (int, error) {

	if checkURL(url) {
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			return 0, err
		}
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		return resp.StatusCode, nil
	}
	return http.StatusBadRequest, errors.New("Невалидный URL-адрес")

}
func getIntervalAndSend() (int, int) {

	send, err := strconv.Atoi(setting.ConfigWebhook.Send)
	if err != nil {
	}
	interval, err := strconv.Atoi(setting.ConfigWebhook.Interval)
	if err != nil {
	}
	return send, interval
}
func statusSucsess(status int) bool {
	var statuses = strings.Trim(setting.ConfigWebhook.SucsessResponse, " ")
	sts := strings.Split(statuses, ",")
	stringstatus := string(status)
	for _, rsp := range sts {

		if rsp == stringstatus {
			return true
		}
	}

	return false
}
func checkURL(url string) bool {

	haveHTTP, err := regexp.MatchString("http://", url)
	if err != nil {
		log.ServiceError("webhook/Send/checkURL", err.Error())

	}
	if !haveHTTP {
		url = "http://" + url
	}
	if !govalidator.IsURL(url) {
		return false
	}

	return true
}
