package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var act_id = "e202102251931481"
var header = http.Header{
	"Accept":          []string{"application/json, text/plain, */*"},
	"Accept-Language": []string{"en-US,en;q=0.5"},
	"Connection":      []string{"keep-alive"},
	"Origin":          []string{"https://webstatic-sea.mihoyo.com"},
}

func request(client *http.Client, url string, method string, header http.Header, params map[string]string, body io.Reader) (map[string]interface{}, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create new request error: %v", err)
	}
	req.Header = header
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	json, err := Bodytojson(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("body to json error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: %v - %v (%v)", resp.StatusCode, resp.Status, json["message"])
	}

	return json, nil
}

func GetReward(client *http.Client) ([]interface{}, error) {
	params := map[string]string{"lang": "en-us", "act_id": act_id}
	resp, err := request(client, "https://hk4e-api-os.mihoyo.com/event/sol/home", "GET", header, params, nil)
	if err != nil {
		return nil, err
	}
	return resp["data"].(map[string]interface{})["awards"].([]interface{}), nil
}

func GetStatus(client *http.Client) (bool, float64, error) {
	header := header
	header.Add("Cache-Control", "max-age=0")
	header.Add("Referer", fmt.Sprintf("https://webstatic-sea.mihoyo.com/ys/event/signin-sea/index.html?act_id=%v&lang=en-us", act_id))
	params := map[string]string{"lang": "en-us", "act_id": act_id}
	resp, err := request(client, "https://hk4e-api-os.mihoyo.com/event/sol/info", "GET", header, params, nil)
	if err != nil {
		return true, 0, err
	}
	if val, exist := resp["data"]; exist {
		return val.(map[string]interface{})["is_sign"].(bool), val.(map[string]interface{})["total_sign_day"].(float64), nil
	}
	return true, 0, fmt.Errorf("response data not found: %v", resp["message"])
}

func ClaimReward(client *http.Client) (map[string]interface{}, error) {
	header := header
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Referer", fmt.Sprintf("https://webstatic-sea.mihoyo.com/ys/event/signin-sea/index.html?act_id=%v&lang=en-us", act_id))
	params := map[string]string{"lang": "en-us"}
	data := map[string]string{"act_id": act_id}
	json, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("create json data error: %v", err)
	}
	resp, err := request(client, "https://hk4e-api-os.mihoyo.com/event/sol/sign", "POST", header, params, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
