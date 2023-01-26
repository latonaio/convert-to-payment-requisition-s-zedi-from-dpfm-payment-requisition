package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-payment-requisition-s-zedi-from-dpfm-payment-requisition/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToSZEDI(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]SZEDI, error) {
	var sZEDIs []SZEDI
	mappingSZEDI := psdc.MappingSZEDI
	codeConversionSZEDI := psdc.CodeConversionSZEDI

	for i := range *mappingSZEDI {
		sZEDI := SZEDI{}

		data, err := json.Marshal(mappingSZEDI)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &sZEDI)
		if err != nil {
			return nil, err
		}

		sZEDI.PayerID = (*codeConversionSZEDI)[i].PayerID
		sZEDI.PayeeID = (*codeConversionSZEDI)[i].PayeeID
		sZEDI.BuyerID = (*codeConversionSZEDI)[i].BuyerID
	}

	return &sZEDIs, nil
}
