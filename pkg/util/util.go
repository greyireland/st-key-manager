package util

import (
	"encoding/json"
	"github.com/greyireland/log"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

var (
	httpclient = &http.Client{
		Timeout: time.Second * 15,
	}
)

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
func Exists(set []string, find string) bool {
	for _, s := range set {
		if s == find {
			return true
		}
	}
	return false
}
func NewDecimal(v string) decimal.Decimal {
	ret, _ := decimal.NewFromString(v)
	return ret
}
func UnitToEther(a decimal.Decimal, decimals int32) decimal.Decimal {
	ether := decimal.New(1, decimals)
	ret := a.Div(ether)
	return ret
}
func EtherToUnit(a decimal.Decimal, decimals int32) decimal.Decimal {
	ether := decimal.New(1, decimals)
	ret := a.Mul(ether)
	return ret
}
func NewBigInt(s string, base int) *big.Int {
	ret, _ := (&big.Int{}).SetString(s, base)
	return ret
}
func JSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
func GetDecimals(v string) int32 {
	vv, err := decimal.NewFromString(v)
	if err != nil {
		return 0
	}
	count := int32(0)
	for i := 0; i < len(v); i++ {
		vv = vv.Mul(decimal.NewFromFloat(10))
		if vv.GreaterThan(decimal.NewFromFloat(1)) {
			return count
		} else {
			count++
		}
	}
	return 0
}
