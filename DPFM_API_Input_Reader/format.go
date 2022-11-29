package dpfm_api_input_reader

import (
	"data-platform-api-fin-inst-master-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *GeneralSDC) ConvertToFinInstMasterGeneral() *requests.FinInstMasterGeneral {
	data := sdc.FinInstMasterGeneral
	return &requests.FinInstMasterGeneral{
		FinInstCountry: data.FinInstCountry,
		FinInstCode:    data.FinInstCode,
	}
}

func (sdc *BranchSDC) ConvertToFinInstMasterBranch() *requests.FinInstMasterBranch {
	data := sdc.FinInstMasterBranch
	return &requests.FinInstMasterBranch{
		FinInstCountry:    data.FinInstCountry,
		FinInstCode:       data.FinInstCode,
		FinInstBranchCode: data.FinInstBranchCode,
		FinInstFullCode:   data.FinInstFullCode,
	}
}
