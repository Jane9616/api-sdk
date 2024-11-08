package demo

import (
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetBalance(t *testing.T) {
	client := apisdk.NewClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)
	_, err := client.CompanyWallet.GetBalance(&apisdk.ParamGetBalance{
		ChainName: "Polygon",
		Currency:  "USD",
		Page:      0,
		Limit:     1,
	})
	if err != nil {
		t.Error(err)
	}
	// fmt.Println("resp: ", resp)
	// fmt.Println("err: ", err)
	// t.Error(resp)
}
