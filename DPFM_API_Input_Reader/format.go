package dpfm_api_input_reader

import (
	"convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToDPFMPaymentRequisitionHeader() *requests.DPFMPaymentRequisitionHeader {
	data := sdc.DPFMPaymentRequisitionHeader
	return &requests.DPFMPaymentRequisitionHeader{
		PayerPaymentRequisitionID:       data.PayerPaymentRequisitionID,
		Payer:                           data.Payer,
		PayerPaymentDate:                data.PayerPaymentDate,
		PaymentReqnStatus:               data.PaymentReqnStatus,
		AcceptanceNoByFinInst:           data.AcceptanceNoByFinInst,
		PaytReqnAmtInTransCrcy:          data.PaytReqnAmtInTransCrcy,
		Currency:                        data.Currency,
		PaymentMethod:                   data.PaymentMethod,
		PayerPostingDate:                data.PayerPostingDate,
		PayerHouseBank:                  data.PayerHouseBank,
		PayerHouseBankAccount:           data.PayerHouseBankAccount,
		PayerFinInstCountry:             data.PayerFinInstCountry,
		PayerFinInstCode:                data.PayerFinInstCode,
		PayerFinInstBranchCode:          data.PayerFinInstBranchCode,
		PayerFinInstFullCode:            data.PayerFinInstFullCode,
		PayerFinInstSWIFTCode:           data.PayerFinInstSWIFTCode,
		PayerInternalFinInstCustomerID:  data.PayerInternalFinInstCustomerID,
		PayerInternalFinInstAccountID:   data.PayerInternalFinInstAccountID,
		PayerFinInstControlKey:          data.PayerFinInstControlKey,
		PayerFinInstAccount:             data.PayerFinInstAccount,
		PayerFinInstAccountName:         data.PayerFinInstAccountName,
		PayerFinInstName:                data.PayerFinInstName,
		PayerFinInstBranchName:          data.PayerFinInstBranchName,
		PayerFinInstCustomerIDByFinInst: data.PayerFinInstCustomerIDByFinInst,
		PaymentRequisitionIsCanceled:    data.PaymentRequisitionIsCanceled,
		CreationDateTime:                data.CreationDateTime,
		ChangedOnDateTime:               data.ChangedOnDateTime,
	}
}

func (sdc *SDC) ConvertToDPFMPaymentRequisitionItem(num int) *requests.DPFMPaymentRequisitionItem {
	dataHeader := sdc.DPFMPaymentRequisitionHeader
	data := sdc.DPFMPaymentRequisitionHeader.DPFMPaymentRequisitionItem[num]
	return &requests.DPFMPaymentRequisitionItem{
		PayerPaymentRequisitionID:        dataHeader.PayerPaymentRequisitionID,
		PayerPaymentRequisitionItem:      data.PayerPaymentRequisitionItem,
		Payer:                            dataHeader.Payer,
		PayerPaymentDate:                 dataHeader.PayerPaymentDate,
		SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		Payee:                            data.Payee,
		Buyer:                            data.Buyer,
		Seller:                           data.Seller,
		InvoiceDocument:                  data.InvoiceDocument,
		PayeeHouseBank:                   data.PayeeHouseBank,
		PayeeHouseBankAccount:            data.PayeeHouseBankAccount,
		PayeeFinInstCountry:              data.PayeeFinInstCountry,
		PayeeFinInstCode:                 data.PayeeFinInstCode,
		PayeeFinInstBranchCode:           data.PayeeFinInstBranchCode,
		PayeeFinInstFullCode:             data.PayeeFinInstFullCode,
		PayeeFinInstSWIFTCode:            data.PayeeFinInstSWIFTCode,
		PaytReqnItemAmtInTransCrcy:       data.PaytReqnItemAmtInTransCrcy,
		PayeeInternalFinInstCustomerID:   data.PayeeInternalFinInstCustomerID,
		PayeeInternalFinInstAccountID:    data.PayeeInternalFinInstAccountID,
		PayeeFinInstControlKey:           data.PayeeFinInstControlKey,
		PayeeFinInstAccount:              data.PayeeFinInstAccount,
		PayeeFinInstAccountName:          data.PayeeFinInstAccountName,
		PayeeFinInstName:                 data.PayeeFinInstName,
		PayeeFinInstBranchName:           data.PayeeFinInstBranchName,
		PayerAccountingDocument:          data.PayerAccountingDocument,
		PayerAccountingDocumentItem:      data.PayerAccountingDocumentItem,
		CreationDateTime:                 data.CreationDateTime,
		ChangedOnDateTime:                data.ChangedOnDateTime,
	}
}
