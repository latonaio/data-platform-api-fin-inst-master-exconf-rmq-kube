package requests

type FinInstMasterBranch struct {
	FinInstCountry    *string `json:"FinInstCountry"`
	FinInstCode       *string `json:"FinInstCode"`
	FinInstBranchCode *string `json:"FinInstBranchCode"`
	FinInstFullCode   *string `json:"FinInstFullCode"`
}
