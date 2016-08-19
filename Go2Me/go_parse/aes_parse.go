package go_parse

import (
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/sinlov/fastEncryptDecode"
)

func byte2String(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func simple_http_get(urlStr string, params map[string]string) (string, error) {
	u, _ := url.Parse(urlStr)
	q := u.Query()
	if params != nil {
		for k, v := range params {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	resByte, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	result := byte2String(resByte)
	return result, nil
}

func AES_Http_Get(url string, key string) (string, error) {
	httpRes, err := simple_http_get(url, nil)
	if err != nil {
		return "", err
	}
	deStr, err := fastEncryptDecode.AES_CBC_PKCS7_Decrypt(httpRes, key)
	if err != nil {
		return "", err
	}
	return deStr, nil
}