package convert_complementer

import (
	dpfm_api_input_reader "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping SZEDIの処理
func (c *ConvertComplementer) ComplementMappingSZEDI(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingSZEDI, error) {
	res, err := psdc.ConvertToMappingSZEDI(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionSZEDI(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionSZEDI, error) {
	var data []dpfm_api_processing_formatter.CodeConversionSZEDI

	for _, sZEDI := range sdc.DPFMPaymentRequisitionHeader.DPFMPaymentRequisitionItem {

		var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
		var args []interface{}

		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Payer", "PayerID", sdc.DPFMPaymentRequisitionHeader.Payer))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Payee", "PayeeID", sZEDI.Payee))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Buyer", "BuyerID", sZEDI.Buyer))

		repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?)"
		for _, v := range dataKey {
			args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
		}

		rows, err := c.db.Query(
			`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
		CodeConvertFrom, CodeConvertTo, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
		WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}

		dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
		if err != nil {
			return nil, err
		}

		datum, err := psdc.ConvertToCodeConversionSZEDI(dataQueryGets)
		if err != nil {
			return nil, err
		}

		data = append(data, *datum)
	}

	return &data, nil
}
