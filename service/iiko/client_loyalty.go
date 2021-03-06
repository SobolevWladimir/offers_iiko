package iiko

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"offers_iiko/lib/log"
	"offers_iiko/mentity/offerentity"
	"offers_iiko/mentity/transport"
	"time"
)

const URL_Loyalty = "/api/0/orders/calculate_checkin_result"

func GetLoality(auth AuthData, order transport.IOrderRequest, tprod TableProduct) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	if len(order.Organization) == 0 {
		return result, errors.New("для получения данных  необходим  id  организации ")
	}
	token, err := GetToken(auth)
	if err != nil {
		return result, err
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
		return result, err
	}
	log.LogIIKOReques(jsonStr)
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	log.LogIIKOResponce(robots)

	if resp.StatusCode != 200 {
		return result, errors.New(string(robots))
	}

	check_result := CheckinResult{}
	err = json.Unmarshal(robots, &check_result)
	if err != nil {
		return result, err
	}
	actions, err := check_result.GetActons(order, tprod)
	if err != nil {
		return result, err
	}
	if actions.ContainsiByType(offerentity.TypeUpsale) {
		return actions, nil

	}
	return offerentity.Actions{}, nil
}
