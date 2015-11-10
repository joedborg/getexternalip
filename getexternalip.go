package getexternalip

import (
	"io/ioutil"
	"net/http"
)

func GetExternalIP() (string, error) {
	resp, err := http.Get("http://echoip.com")
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
