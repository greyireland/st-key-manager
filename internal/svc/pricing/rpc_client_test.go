package pricing

import (
	"fmt"
	"solgateway/pkg/jsonx"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	ret := make([]UserInfo, 0)
	got := GetUserInfo()
	for i := 0; i < len(got); i++ {
		if got[i].Role > 0 {
			ret = append(ret, got[i])
		}
	}
	fmt.Println(jsonx.JSON(ret))
}
