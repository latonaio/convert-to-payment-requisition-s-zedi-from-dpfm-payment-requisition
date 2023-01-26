package dpfm_api_input_reader

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
	ConnectionKey                string                       `json:"connection_key"`
	Result                       bool                         `json:"result"`
	RedisKey                     string                       `json:"redis_key"`
	Filepath                     string                       `json:"filepath"`
	APIStatusCode                int                          `json:"api_status_code"`
	RuntimeSessionID             string                       `json:"runtime_session_id"`
	BusinessPartnerID            *int                         `json:"business_partner"`
	ServiceLabel                 string                       `json:"service_label"`
	DPFMPaymentRequisitionHeader DPFMPaymentRequisitionHeader `json:"DPFMPaymentRequisition"`
	APISchema                    string                       `json:"api_schema"`
	Accepter                     []string                     `json:"accepter"`
	Deleted                      bool                         `json:"deleted"`
}

type DPFMPaymentRequisitionHeader struct {
	PayerPaymentRequisitionID       int                          `json:"PayerPaymentRequisitionID"`
	Payer                           *int                         `json:"Payer"`
	PayerPaymentDate                *string                      `json:"PayerPaymentDate"`
	PaymentReqnStatus               *string                      `json:"PaymentReqnStatus"`
	AcceptanceNoByFinInst           *string                      `json:"AcceptanceNoByFinInst"`
	PaytReqnAmtInTransCrcy          *float32                     `json:"PaytReqnAmtInTransCrcy"`
	Currency                        *string                      `json:"Currency"`
	PaymentMethod                   *string                      `json:"PaymentMethod"`
	PayerPostingDate                *string                      `json:"PayerPostingDate"`
	PayerHouseBank                  *string                      `json:"PayerHouseBank"`
	PayerHouseBankAccount           *string                      `json:"PayerHouseBankAccount"`
	PayerFinInstCountry             *string                      `json:"PayerFinInstCountry"`
	PayerFinInstCode                *string                      `json:"PayerFinInstCode"`
	PayerFinInstBranchCode          *string                      `json:"PayerFinInstBranchCode"`
	PayerFinInstFullCode            *string                      `json:"PayerFinInstFullCode"`
	PayerFinInstSWIFTCode           *string                      `json:"PayerFinInstSWIFTCode"`
	PayerInternalFinInstCustomerID  *int                         `json:"PayerInternalFinInstCustomerID"`
	PayerInternalFinInstAccountID   *int                         `json:"PayerInternalFinInstAccountID"`
	PayerFinInstControlKey          *string                      `json:"PayerFinInstControlKey"`
	PayerFinInstAccount             *string                      `json:"PayerFinInstAccount"`
	PayerFinInstAccountName         *string                      `json:"PayerFinInstAccountName"`
	PayerFinInstName                *string                      `json:"PayerFinInstName"`
	PayerFinInstBranchName          *string                      `json:"PayerFinInstBranchName"`
	PayerFinInstCustomerIDByFinInst *string                      `json:"PayerFinInstCustomerIDByFinInst"`
	PaymentRequisitionIsCanceled    *bool                        `json:"PaymentRequisitionIsCanceled"`
	CreationDateTime                *string                      `json:"CreationDateTime"`
	ChangedOnDateTime               *string                      `json:"ChangedOnDateTime"`
	DPFMPaymentRequisitionItem      []DPFMPaymentRequisitionItem `json:"DPFMPaymentRequisitionItem"`
}

type DPFMPaymentRequisitionItem struct {
	PayerPaymentRequisitionID        int      `json:"PayerPaymentRequisitionID"`
	PayerPaymentRequisitionItem      int      `json:"PayerPaymentRequisitionItem"`
	Payer                            *int     `json:"Payer"`
	PayerPaymentDate                 *string  `json:"PayerPaymentDate"`
	SupplyChainRelationshipID        *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	Payee                            *int     `json:"Payee"`
	Buyer                            *int     `json:"Buyer"`
	Seller                           *int     `json:"Seller"`
	InvoiceDocument                  *int     `json:"InvoiceDocument"`
	PayeeHouseBank                   *string  `json:"PayeeHouseBank"`
	PayeeHouseBankAccount            *string  `json:"PayeeHouseBankAccount"`
	PayeeFinInstCountry              *string  `json:"PayeeFinInstCountry"`
	PayeeFinInstCode                 *string  `json:"PayeeFinInstCode"`
	PayeeFinInstBranchCode           *string  `json:"PayeeFinInstBranchCode"`
	PayeeFinInstFullCode             *string  `json:"PayeeFinInstFullCode"`
	PayeeFinInstSWIFTCode            *string  `json:"PayeeFinInstSWIFTCode"`
	PaytReqnItemAmtInTransCrcy       *float32 `json:"PaytReqnItemAmtInTransCrcy"`
	PayeeInternalFinInstCustomerID   *int     `json:"PayeeInternalFinInstCustomerID"`
	PayeeInternalFinInstAccountID    *int     `json:"PayeeInternalFinInstAccountID"`
	PayeeFinInstControlKey           *string  `json:"PayeeFinInstControlKey"`
	PayeeFinInstAccount              *string  `json:"PayeeFinInstAccount"`
	PayeeFinInstAccountName          *string  `json:"PayeeFinInstAccountName"`
	PayeeFinInstName                 *string  `json:"PayeeFinInstName"`
	PayeeFinInstBranchName           *string  `json:"PayeeFinInstBranchName"`
	PayerAccountingDocument          *int     `json:"PayerAccountingDocument"`
	PayerAccountingDocumentItem      *int     `json:"PayerAccountingDocumentItem"`
	CreationDateTime                 *string  `json:"CreationDateTime"`
	ChangedOnDateTime                *string  `json:"ChangedOnDateTime"`
}
