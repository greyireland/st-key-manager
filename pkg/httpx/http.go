package httpx

import (
	"encoding/json"
	"github.com/greyireland/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	httpclient = &http.Client{
		Timeout: time.Second * 15,
	}
)

func Get(url string, ret interface{}) (err error) {
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Warn(err.Error())
		return
	}
	res, err := httpclient.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	defer res.Body.Close()
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		log.Warn("unmarshal error", "err", err)
		return
	}
	return
}
func Post(url string, body string, ret interface{}) (err error) {
	method := "POST"

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Warn(err.Error())
		return
	}
	res, err := httpclient.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	defer res.Body.Close()
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		log.Warn("unmarshal error", "err", err)
		return
	}
	return
}
func PostH(url string, header http.Header, body string, ret interface{}) (err error) {
	method := "POST"

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	req.Header = header
	if err != nil {
		log.Warn(err.Error())
		return
	}
	res, err := httpclient.Do(req)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	defer res.Body.Close()
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn(err.Error())
		return
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		log.Warn("unmarshal error", "err", err)
		return
	}
	log.Info("post success", "url", url, "body", body, "code", res.StatusCode, "resp", string(resp))
	return
}
