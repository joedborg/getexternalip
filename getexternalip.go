package getexternalip

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetExternalIP returns your external IP address as returned by http://echoip.com.
func GetExternalIP() (string, error) {
	url := "http://echoip.com"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET %q: %v", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
