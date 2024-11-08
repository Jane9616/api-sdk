package demo

import (
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetCompanyWalletHDWalletAddress(t *testing.T) {
	client := apisdk.NewClient("APIKey", "APISecret", 10*time.Second)
	_, err := client.CompanyWallet.GetCompanyWalletHDWalletAddress(&apisdk.ParamGetCompanyWalletHDWalletAddress{
		HDWalletID: "HDWalletID",
	})
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(resp, err)
}
