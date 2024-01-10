package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-usage-control-chain-creates-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context

	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
	db  *database.Mysql
}

func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		c:   c,
		rmq: rmq,
		db:  db,
	}
}

// Confirm returns existenceMap, allExist, err
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, ssdc *dpfm_api_output_formatter.SDC, accepter []string, l *logger.Logger) (allExist bool, errs []error) {

	return true, nil
}

func (c *ExistenceConf) getExConfMapper(serviceLabel string) (*[]ExConfMapper, error) {
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformCommonSettingsMysqlKube.data_platform_ex_conf_api_mapper_data
		WHERE ServiceLabel = ?;`, serviceLabel,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]ExConfMapper, 0)

	i := 0
	for rows.Next() {
		i++
		data := ExConfMapper{}

		err := rows.Scan(
			&data.ServiceLabel,
			&data.APIType,
			&data.APIName,
			&data.Field,
			&data.ExConfAPIServiceName,
			&data.ExConfAPIName,
			&data.ExConfAPIQueueName,
			&data.ExConfProgramFileName,
			&data.Tabletag,
			&data.TableConfirmed,
			&data.ExConfAPIType,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, data)
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_ex_conf_api_mapper_data'テーブルに対象のレコードが存在しません。")
	}

	return &res, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
func confKeyExistence(res map[string]interface{}, tableTag string) bool {
	//if res == nil {
	//	return false
	//}
	//if tableTag == "ProductMasterGeneral" {
	//	return productMasterConfKeyExistence(res, tableTag)
	//}
	//raw, ok := res["ExistenceConf"]
	//exist := fmt.Sprintf("%v", raw)
	//if ok {
	//	return strings.ToLower(exist) == "true"
	//}

	return false
}

func (c *ExistenceConf) exconfRequest(req interface{}, mapper ExConfMapper, log *logger.Logger) (bool, error) {
	//queueName, err := getQueueName(mapper)
	//if err != nil {
	//	return false, err
	//}
	//tableTag := *mapper.Tabletag
	//
	//res, err := c.rmq.SessionKeepRequest(nil, queueName, req)
	//if err != nil {
	//	return false, xerrors.Errorf("response error: %w", err)
	//}
	//res.Success()
	//exist := confKeyExistence(res.Data(), tableTag)
	//log.Info(res.Data())
	//return exist, nil
	return true, nil
}

type result struct {
	keys map[string]interface{}
}

func newResult(keys map[string]interface{}) *result {
	return &result{
		keys: keys,
	}
}

func (r *result) fail() string {
	txt := ""
	for k, v := range r.keys {
		txt = fmt.Sprintf("%s%s:%v, ", k, v)
	}
	txt = fmt.Sprintf("%s does not exist", txt)
	return txt
}

func getQueueName(mapper ExConfMapper) (string, error) {
	var err error

	if mapper.ExConfAPIQueueName == nil {
		err = xerrors.Errorf("cannot specify null queue name")
		return "", err
	}
	queueName := *mapper.ExConfAPIQueueName

	return queueName, nil
}

func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func isZero[T comparable](obj T) bool {
	var zero T
	return obj == zero
}
