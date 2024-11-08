package demo

import (
	"testing"
	"time"

	apisdk "api-sdk"
	// "github.com/tidwall/gjson"
)

func TestGetApprovals(t *testing.T) {
	client := apisdk.NewClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)
	_, err := client.CompanyWallet.GetApprovals(&apisdk.ParamGetApprovals{
		Page:   1,
		Limit:  1,
		Status: "ING",
	})

	// fmt.Println(resp, err)
	if err != nil {
		t.Error(err)
	}

	// "github.com/stretchr/testify/assert
	// assert.Equal(t, 2, result)

	// t.Error(resp)
}
