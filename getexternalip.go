package getexternalip

import (
	"io/ioutil"
	"net/http"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func getExternalIP() string {
	resp, err := http.Get("http://echoip.com")
	errorCheck(err)
	body, err := ioutil.ReadAll(resp.Body)
	errorCheck(err)
	return string(body)
}
