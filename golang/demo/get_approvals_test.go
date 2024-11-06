package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/openblock/openblock-api-sdk-go"
)

func TestGetApprovals(t *testing.T) {
	client := apisdk.NewClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)
	resp, err := client.CompanyWallet.GetApprovals(&apisdk.ParamGetApprovals{
		Page:   1,
		Limit:  1,
		Status: "ING",
	})
	fmt.Println(resp, err)
	t.Error(resp)
}
