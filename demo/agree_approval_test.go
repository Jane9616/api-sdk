package demo

import (
	"testing"
	"time"

	apisdk "github.com/OpenBlockResource/openblock-api-sdk-go"
)

// NewCompanyWalletClient

func TestAgreeApproval(t *testing.T) {

	client := apisdk.NewCompanyWalletClient("544eb8f5b5c748cab4adce144851ecdc", "bOiYCXLd6WTDpvodD6Zd9q3ezfsb5ccI", 10*time.Second)

	_, err := client.AgreeApproval(&apisdk.ParamAgreeApproval{
		RecordID: "0c3ce1d251414264b3614f353c89ffa21",
		Agree:    "agree",
	})

	// client := apisdk.NewClient("ffca25bb9d6c4e7cafcaf9b95851a603", "dwZSAb2XnNIRGYmGLWvNPz5CHSggqq59", 10*time.Second)

	// resp, err := client.CompanyWallet.AgreeApproval(&apisdk.ParamAgreeApproval{
	// 	RecordID: "0c3ce1d251414264b3614f353c89ffa21",
	// 	Agree:    "agree",
	// })
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(resp, err)
	// t.Error(resp)
}
