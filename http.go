package client

import (
	"strings"
	"net/http"
	"io/ioutil"
)

func post(url string, data map[string]string) (map[string]interface{}, error) {

	payload := strings.NewReader(makeFormData(data, "boundaryhere"))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=boundaryhere")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return map[string]interface{}{}, err
	}

	var bodyJson map[string]interface{}
	json.Unmarshal(body, &bodyJson)

	return bodyJson, nil
}

func makeFormData(data map[string]string, boundary string) string {
	var formStr = ""
	for k,v := range data {
		formStr += "--" + boundary + `\r\nContent-Disposition: form-data; name="` + k + `"\r\n\r\n` + v + "\r\n"
	}
	formStr += "--boundaryhere--"
	return formStr
}