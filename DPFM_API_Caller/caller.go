package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-fin-inst-master-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-fin-inst-master-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["FinInstMasterGeneral"]
	if ok {
		input := &dpfm_api_input_reader.GeneralSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confFinInstMasterGeneral(input)
		goto endProcess
	}
	_, ok = input["FinInstMasterBranch"]
	if ok {
		input := &dpfm_api_input_reader.BranchSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfFinInstMasterBranch(input)
		goto endProcess
	}

	err = xerrors.Errorf("can not get exconf check target")
endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confFinInstMasterGeneral(input *dpfm_api_input_reader.GeneralSDC) *dpfm_api_output_formatter.FinInstMasterGeneral {
	exconf := dpfm_api_output_formatter.FinInstMasterGeneral{
		ExistenceConf: false,
	}
	if input.FinInstMasterGeneral.FinInstCode == nil {
		return &exconf
	}
	if input.FinInstMasterGeneral.FinInstCountry == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.FinInstMasterGeneral{
		FinInstCountry: *input.FinInstMasterGeneral.FinInstCountry,
		FinInstCode:    *input.FinInstMasterGeneral.FinInstCode,
		ExistenceConf:  false,
	}

	rows, err := e.db.Query(
		`SELECT FinInstName
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_master_general_data 
		WHERE (finInstCountry, finInstCode) = (?, ?);`, exconf.FinInstCountry, exconf.FinInstCode,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfFinInstMasterBranch(input *dpfm_api_input_reader.BranchSDC) *dpfm_api_output_formatter.FinInstMasterBranch {
	exconf := dpfm_api_output_formatter.FinInstMasterBranch{
		ExistenceConf: false,
	}
	if input.FinInstMasterBranch.FinInstCode == nil {
		return &exconf
	}
	if input.FinInstMasterBranch.FinInstCountry == nil {
		return &exconf
	}
	if input.FinInstMasterBranch.FinInstBranchCode == nil {
		return &exconf
	}
	if input.FinInstMasterBranch.FinInstFullCode == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.FinInstMasterBranch{
		FinInstCountry:    *input.FinInstMasterBranch.FinInstCountry,
		FinInstCode:       *input.FinInstMasterBranch.FinInstCode,
		FinInstBranchCode: *input.FinInstMasterBranch.FinInstBranchCode,
		FinInstFullCode:   *input.FinInstMasterBranch.FinInstFullCode,
		ExistenceConf:     false,
	}
	rows, err := e.db.Query(
		`SELECT FinInstCode 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_fin_inst_master_branch_data 
		WHERE (finInstCountry, finInstCode, finInstBranchCode, finInstFullCode) = (?, ?, ?, ?);`,
		exconf.FinInstCountry, exconf.FinInstCode, exconf.FinInstBranchCode, exconf.FinInstFullCode,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
