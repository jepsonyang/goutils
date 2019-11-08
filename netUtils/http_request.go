package netUtils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
* http请求
**/
func Request(url string, headers map[string]string, bodyByte []byte, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k,v := range headers {
			req.Header.Add(k, v)
		}
	}

	client := http.Client{}
	rsp, err := client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return rspBody, nil
}