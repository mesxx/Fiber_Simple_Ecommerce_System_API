package helpers

import (
	"errors"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

var sc snap.Client
var cc coreapi.Client

func CreateSnapTransactionURL(snapReq *snap.Request) (string, error) {
	sc.New(example.SandboxServerKey1, midtrans.Sandbox)

	// START create transaction urrl
	res, err := sc.CreateTransactionUrl(snapReq)
	if err != nil {
		return "", errors.New(err.Error())
	}
	// END create transaction urrl

	return res, nil
}

func GetCoreAPITransactionData(orderId string) (*coreapi.TransactionStatusResponse, error) {
	cc.New(example.SandboxServerKey1, midtrans.Sandbox)

	// START check transaction
	res, err := cc.CheckTransaction(orderId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	// END check transaction

	return res, nil
}

func GenerateSnapReq(fname string, lname string, email string, amount int) *snap.Request {
	// START Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       fname,
		LName:       lname,
		Phone:       "08xxx",
		Address:     "Dummy Baker Street 97th",
		City:        "Dummy Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}
	// END Initiate Customer address

	// START Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-GO-ID-" + example.Random(),
			GrossAmt: int64(amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    fname,
			LName:    lname,
			Email:    email,
			Phone:    "08xxx",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
	// END Initiate Snap Request

	return snapReq
}
