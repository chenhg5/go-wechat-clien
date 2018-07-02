package client

import (
	"github.com/json-iterator/go"
	"errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	server_addr string
	acid        string
}

var client = &Client{
	"", "",
}

func InitClient(addr string, acid string)  {
	(*client).server_addr = addr
	(*client).acid = acid
}

func SetAddr(addr string) *Client {
	(*client).server_addr = addr
	return client
}

func (client *Client) SetAcid(acid string) *Client {
	(*client).acid = acid
	return client
}

func (client *Client) WxappOauth(code string) (map[string]interface{}, error) {
	return GetData(post((*client).server_addr, map[string]string{
		"accountId" : (*client).acid,
		"method" : "WxappOauth",
		"jsCode" : code,
	}))
}

func (client *Client) DecodeWxappData(sessionKey string, iv string, encryptedData string) (map[string]interface{}, error) {
	data, err := post((*client).server_addr, map[string]string{
		"accountId" : (*client).acid,
		"method" : "DecodeWxappData",
		"sessionKey" : sessionKey,
		"iv" : iv,
		"encryptedData" : encryptedData,
	})

	if data["code"].(float64) != 200 {
		return map[string]interface{}{}, errors.New(data["msg"].(string))
	}

	return data["data"].(map[string]interface{}), err
}

func GetData(olddata map[string]interface{}, err error) (map[string]interface{}, error) {
	if err != nil {
		return map[string]interface{}{}, err
	}

	if olddata["code"].(float64) != 200 {
		return map[string]interface{}{}, errors.New(olddata["msg"].(string))
	}

	wechatRes := olddata["data"].(map[string]interface{})

	errcode, ok := wechatRes["errcode"]

	if ok && errcode.(float64) != 0 {
		return map[string]interface{}{}, errors.New(wechatRes["errmsg"].(string))
	}

	return olddata["data"].(map[string]interface{}), nil
}