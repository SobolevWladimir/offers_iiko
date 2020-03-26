package iiko

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type AuthData struct {
	UserId       string
	UserSecret   string
	Organization string
}

const URL_Auth = "/api/0/auth/access_token"

func GetToken(auth AuthData) (string, error) {
	if len(auth.UserId) == 0 {
		return "", errors.New("user_id  для подключения к iiko пустой")
	}
	if len(auth.UserSecret) == 0 {
		return "", errors.New("user_secret  для подключения к iiko пустой")
	}
	client := &http.Client{
		Timeout: 6 * time.Second,
	}
	vals := url.Values{}
	vals.Add("user_id", auth.UserId)
	vals.Add("user_secret", auth.UserSecret)
	url := url.URL{
		Scheme:   BizScheme,
		Host:     BizHost,
		Path:     URL_Auth,
		RawQuery: vals.Encode(),
	}
	resp, err := client.Get(url.String())
	if err != nil {
		return "", errors.New("Ошибка подключения к iiko biz")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("auth iiko biz: " + err.Error())
		return "", errors.New("Ошибка получения токена  iiko biz(неверный логин или пароль)")
	}
	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result string
	json.Unmarshal(robots, &result)
	return result, nil
}
