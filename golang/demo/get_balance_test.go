package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/openblock/openblock-api-sdk-go"
)

func TestGetBalance(t *testing.T) {
	client := apisdk.NewClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)
	resp, err := client.CompanyWallet.GetBalance(&apisdk.ParamGetBalance{
		ChainName: "Polygon",
		Currency:  "USD",
		Page:      0,
		Limit:     1,
	})
	fmt.Println("resp: ", resp)
	fmt.Println("err: ", err)
	t.Error(resp)
}
