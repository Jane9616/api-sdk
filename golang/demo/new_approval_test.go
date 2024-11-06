package demo

import (
	"fmt"
	"testing"
	"time"

	apisdk "github.com/openblock/openblock-api-sdk-go"
)

func TestNewApproval(t *testing.T) {
	client := apisdk.NewClient("cc5eebae6b5a4049aff24c884e969ade", "gVuzsNFHdHXsw7AL0filahHQe0u0yW1P", 10*time.Second)
	resp, err := client.CompanyWallet.NewApproval(&apisdk.ParamNewApproval{
		Action:     "TRANSACTION",
		HDWalletID: "36978258926b4536b5e9974c0ee962e1",
		TXInfo: apisdk.TXInfo{
			To:                   "0x27D25bBDA38F681c34A514E74D4e193BC1c2f19E",
			From:                 "0x8288F8c2FC473663E84E263b4359d19AcF6b6a79",
			Value:                "0.00001",
			GasLimit:             "21000",
			GasPrice:             "0",
			MaxFeePerGas:         "243.2473634",
			MaxPriorityFeePerGas: "52.11339668",
			Chain:                "Polygon",
			TransactionType:      "native",
		},
	})
	fmt.Println(resp, err)
	t.Error(resp)
}
