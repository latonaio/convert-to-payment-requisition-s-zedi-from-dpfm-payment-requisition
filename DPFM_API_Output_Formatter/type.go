package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	SZEDI []SZEDI `json:"SZEDI"`
}

type SZEDI struct {
	PayerPaymentRequisitionID   int     `json:"PayerPaymentRequisitionID"`
	PayerPaymentRequisitionItem int     `json:"PayerPaymentRequisitionItem"`
	Payer                       *int    `json:"Payer"`
	PayerPaymentDate            *string `json:"PayerPaymentDate"`
	Payee                       *int    `json:"Payee"`
	BillFromParty               *int    `json:"BillFromParty"`
	BillToParty                 *int    `json:"BillToParty"`
	Buyer                       *int    `json:"Buyer"`
	Seller                      *int    `json:"Seller"`
	SubsetSpecifiedID           *string `json:"SubsetSpecifiedID"`
	BusinessProcessSpecifiedID  *string `json:"BusinessProcessSpecifiedID"`
	ExchangedDocumentID         *string `json:"ExchangedDocumentID"`
	IssueDateTime               *string `json:"IssueDateTime"`
	IssuerAssignedID            *string `json:"IssuerAssignedID"`
	PayerID                     *string `json:"PayerID"`
	PayeeID                     *string `json:"PayeeID"`
	BuyerName                   *string `json:"BuyerName"`
	BuyerID                     *int    `json:"BuyerID"`
	PaymentTotalAmount          *string `json:"PaymentTotalAmount"`
	BalanceOutReasonDescription *string `json:"BalanceOutReasonDescription"`
	BalanceOutCalculatedAmount  *string `json:"BalanceOutCalculatedAmount"`
	TaxCalculatedAmount1        *string `json:"TaxCalculatedAmount1"`
	TaxCalculatedRate1          *string `json:"TaxCalculatedRate1"`
	TaxCalculatedAmount2        *string `json:"TaxCalculatedAmount2"`
	TaxCalculatedRate2          *string `json:"TaxCalculatedRate2"`
	TaxTotalAmount              *string `json:"TaxTotalAmount"`
	Content                     *string `json:"Content"`
}
