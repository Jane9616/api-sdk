package demo

import (
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetCompanyWalletInfo(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	_, err := client.CompanyWallet.GetCompanyWalletInfo()
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(resp, err)
}
