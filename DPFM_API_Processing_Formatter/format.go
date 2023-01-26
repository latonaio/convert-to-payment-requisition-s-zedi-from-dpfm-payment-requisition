package dpfm_api_processing_formatter

import (
	"convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// データ連携基盤 Payment Requisition S Zedi と データ連携基盤 Payment Requisition との項目マッピング
func (psdc *SDC) ConvertToMappingSZEDI(sdc *dpfm_api_input_reader.SDC) (*[]MappingSZEDI, error) {
	var res []MappingSZEDI
	data := sdc.DPFMPaymentRequisitionHeader
	dataItem := sdc.DPFMPaymentRequisitionHeader.DPFMPaymentRequisitionItem

	for _, dataItem := range dataItem {
		invoiceDocument := strconv.Itoa(*dataItem.InvoiceDocument)
		paytReqnAmtInTransCrcy := strconv.FormatFloat(float64(*data.PaytReqnAmtInTransCrcy), 'f', -1, 32)

		res = append(res, MappingSZEDI{
			PayerPaymentRequisitionID:   data.PayerPaymentRequisitionID,
			PayerPaymentRequisitionItem: dataItem.PayerPaymentRequisitionItem,
			Payer:                       data.Payer,
			PayerPaymentDate:            data.PayerPaymentDate,
			Payee:                       dataItem.Payee,
			BillFromParty:               dataItem.BillFromParty,
			BillToParty:                 dataItem.BillToParty,
			Buyer:                       dataItem.Buyer,
			Seller:                      dataItem.Seller,
			IssueDateTime:               data.PayerPaymentDate,
			IssuerAssignedID:            &invoiceDocument,
			PaymentTotalAmount:          &paytReqnAmtInTransCrcy,
			TaxTotalAmount:              &paytReqnAmtInTransCrcy,
		})
	}

	return &res, nil
}

// 番号変換
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	var res []CodeConversionQueryGets

	for i := 0; true; i++ {
		pm := &requests.CodeConversionQueryGets{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionSZEDI(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionSZEDI, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionSZEDI{}

	pm.PayerID = GetStringPtr(dataMap["PayerID"].CodeConvertTo)
	pm.PayeeID = GetStringPtr(dataMap["PayeeID"].CodeConvertTo)
	pm.BuyerID, err = ParseIntPtr(GetStringPtr(dataMap["BuyerID"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionSZEDI{
		PayerID: data.PayerID,
		PayeeID: data.PayeeID,
		BuyerID: data.BuyerID,
	}

	return &res, nil
}
