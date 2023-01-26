package requests

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
