package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/openblock/openblock-api-sdk-go"
)

// NewCompanyWalletClient

func TestAgreeApproval(t *testing.T) {

	client := apisdk.NewCompanyWalletClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)

	resp, err := client.AgreeApproval(&apisdk.ParamAgreeApproval{
		RecordID: "0c3ce1d251414264b3614f353c89ffa21",
		Agree:    "agree",
	})

	// client := apisdk.NewClient("ffca25bb9d6c4e7cafcaf9b95851a603", "dwZSAb2XnNIRGYmGLWvNPz5CHSggqq59", 10*time.Second)

	// resp, err := client.CompanyWallet.AgreeApproval(&apisdk.ParamAgreeApproval{
	// 	RecordID: "0c3ce1d251414264b3614f353c89ffa21",
	// 	Agree:    "agree",
	// })

	fmt.Println(resp, err)
	t.Error(resp)
}
