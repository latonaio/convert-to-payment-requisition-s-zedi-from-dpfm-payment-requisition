package requests

type CodeConversionSZEDI struct {
	PayerID *string `json:"PayerID"`
	PayeeID *string `json:"PayeeID"`
	BuyerID *int    `json:"BuyerID"`
}
