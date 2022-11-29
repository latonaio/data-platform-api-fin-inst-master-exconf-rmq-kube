package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey        string                `json:"connection_key"`
	Result               bool                  `json:"result"`
	RedisKey             string                `json:"redis_key"`
	Filepath             string                `json:"filepath"`
	APIStatusCode        int                   `json:"api_status_code"`
	RuntimeSessionID     string                `json:"runtime_session_id"`
	BusinessPartnerID    *int                  `json:"business_partner"`
	ServiceLabel         string                `json:"service_label"`
	FinInstMasterGeneral *FinInstMasterGeneral `json:"FinInstMasterGeneral,omitempty"`
	FinInstMasterBranch  *FinInstMasterBranch  `json:"FinInstMasterBranch,omitempty"`
	APISchema            string                `json:"api_schema"`
	Accepter             []string              `json:"accepter"`
	Deleted              bool                  `json:"deleted"`
}

type FinInstMasterGeneral struct {
	FinInstCountry string `json:"FinInstCountry"`
	FinInstCode    string `json:"FinInstCode"`
	ExistenceConf  bool   `json:"ExistenceConf"`
}

type FinInstMasterBranch struct {
	FinInstCountry    string `json:"FinInstCountry"`
	FinInstCode       string `json:"FinInstCode"`
	FinInstBranchCode string `json:"FinInstBranchCode"`
	FinInstFullCode   string `json:"FinInstFullCode"`
	ExistenceConf     bool   `json:"ExistenceConf"`
}
