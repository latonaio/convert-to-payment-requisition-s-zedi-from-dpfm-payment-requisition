package convert_complementer

import (
	dpfm_api_input_reader "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Processing_Formatter"
)

func (c *ConvertComplementer) SetValue(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var sZEDI *[]dpfm_api_output_formatter.SZEDI
	var err error

	sZEDI, err = dpfm_api_output_formatter.ConvertToSZEDI(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		SZEDI: *sZEDI,
	}

	return osdc, nil
}
