# data-platform-api-fin-inst-master-exconf-rmq-kube
data-platform-api-fin-inst-master-exconf-rmq-kube は、データ連携基盤において、API で 金融機関マスタの存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、金融機関マスタの存在確認が行われます。

* data-platform-fin-inst-master-general-data.sql（データ連携基盤 金融機関マスタ - 一般データ）
* data-platform-fin-inst-master-branch-data.sql（データ連携基盤 金融機関マスタ - 支店データ）

## caller.go による存在性確認
Input で取得されたファイルに基づいて、caller.go で、 API がコールされます。
caller.go の 以下の箇所が、指定された API をコールするソースコードです。

```
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

endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}
```

## Input
data-platform-api-fin-inst-master-exconf-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "ORDERS",
	"FinInstMasterGeneral": {
		"FinInstCountry": JP,
		"FinInstCode": "0001",
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
}

```

## Output
data-platform-api-fin-inst-master-exconf-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。プラントの対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"service_label": "ORDERS",
	"FinInstMasterGeneral": {
		"FinInstCountry": "JP",
		"FinInstCode": "0001",
		"ExistenceConf": true
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": [
		"All"
	],
	"order_id": null,
	"deleted": false
}
```