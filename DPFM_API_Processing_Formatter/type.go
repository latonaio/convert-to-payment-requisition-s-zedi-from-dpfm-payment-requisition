package dpfm_api_processing_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MappingSZEDI        *[]MappingSZEDI        `json:"MappingSZEDI"`
	CodeConversionSZEDI *[]CodeConversionSZEDI `json:"CodeConversionSZEDI"`
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

// 項目マッピング変換
type MappingSZEDI struct {
	PayerPaymentRequisitionID   int     `json:"PayerPaymentRequisitionID"`
	PayerPaymentRequisitionItem int     `json:"PayerPaymentRequisitionItem"`
	Payer                       *int    `json:"Payer"`
	PayerPaymentDate            *string `json:"PayerPaymentDate"`
	Payee                       *int    `json:"Payee"`
	BillFromParty               *int    `json:"BillFromParty"`
	BillToParty                 *int    `json:"BillToParty"`
	Buyer                       *int    `json:"Buyer"`
	Seller                      *int    `json:"Seller"`
	IssueDateTime               *string `json:"IssueDateTime"`
	IssuerAssignedID            *string `json:"IssuerAssignedID"`
	PaymentTotalAmount          *string `json:"PaymentTotalAmount"`
	TaxTotalAmount              *string `json:"TaxTotalAmount"`
}

// 番号変換
type CodeConversionKey struct {
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionQueryGets struct {
	CodeConversionID  int    `json:"CodeConversionID"`
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	CodeConvertTo     string `json:"CodeConvertTo"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionSZEDI struct {
	PayerID *string `json:"PayerID"`
	PayeeID *string `json:"PayeeID"`
	BuyerID *int    `json:"BuyerID"`
}

// 個別処理
