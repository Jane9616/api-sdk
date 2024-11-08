package demo

import (
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

func TestGetTxHistory(t *testing.T) {
	client := apisdk.NewClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)
	_, err := client.CompanyWallet.GetTxHistory(&apisdk.ParamGetTxHistory{
		ChainName: "Polygon",
		Page:      1,
		Limit:     20,
		OrderBy:   "tx_time",
		Asc:       0,
	})
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(resp, err)
	// t.Error(resp)
}
