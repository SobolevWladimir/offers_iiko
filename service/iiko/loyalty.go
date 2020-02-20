package iiko

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"offers_iiko/mentity/transport"
	"time"
)

const URL_Loyalty = "/api/0/orders/calculate_checkin_result"

func GetLoality(auth AuthData, order transport.IOrderRequest) error {

	if len(order.Organization) == 0 {
		return errors.New("для получения данных  необходим  id  организации ")
	}
	token, err := GetToken(auth)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: 6 * time.Second,
	}
	vals := url.Values{}
	vals.Add("access_token", token)
	url := url.URL{
		Scheme:   BizScheme,
		Host:     BizHost,
		Path:     URL_Loyalty,
		RawQuery: vals.Encode(),
	}
	jsonStr, err := json.Marshal(order)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		//return errors.New(string(robots))
		fmt.Println(string(robots))
		return nil
	}
	fmt.Println("json:", string(robots))

	return nil
}
