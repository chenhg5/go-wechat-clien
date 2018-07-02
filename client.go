package client

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	server_addr string
}

var client *Client

func InitClient(addr string)  {
	(*client).server_addr = addr
}

func SetAddr(addr string) *Client {
	(*client).server_addr = addr
	return client
}

func (client *Client) WxappOauth(code string) (map[string]interface{}, error) {
	return post((*client).server_addr, map[string]string{
		"accountId" : "0",
		"method" : "WxappOauth",
		"jsCode" : code,
	})
}