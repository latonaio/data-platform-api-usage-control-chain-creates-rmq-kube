package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-usage-control-chain-creates-rmq-kube/sub_func_complementer"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"sync"
	"time"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var usageControlChain *dpfm_api_output_formatter.UsageControlChain
	for _, fn := range accepter {
		switch fn {
		case "UsageControlChain":
			usageControlChain = c.usageControlChainCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		UsageControlChain: usageControlChain,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var usageControlChain *dpfm_api_output_formatter.UsageControlChain
	for _, fn := range accepter {
		switch fn {
		case "UsageControlChain":
			usageControlChain = c.usageControlChainUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		UsageControlChain: usageControlChain,
	}

	return data
}

func (c *DPFMAPICaller) usageControlChainCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.UsageControlChain {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// usageControlChainの更新
	//usageControlChainData := subfuncSDC.Message.UsageControlChain
	usageControlChain := input.UsageControlChain

	currentDate := time.Now().Format("2006-01-02")
	usageControlChain.CreationDate = &currentDate
	tmpTime := "00:00:00"
	usageControlChain.CreationTime = &tmpTime

	usageControlChain.LastChangeDate = &currentDate
	usageControlChain.LastChangeTime = &tmpTime

	usageControlChainData := dpfm_api_processing_formatter.ConvertToUsageControlChainUpdates(usageControlChain)

	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": usageControlChainData, "function": "UsageControlChainUsageControlChain", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "UsageControlChain Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToUsageControlChainCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) usageControlChainUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.UsageControlChain {
	usageControlChain := input.UsageControlChain
	usageControlChainData := dpfm_api_processing_formatter.ConvertToUsageControlChainUpdates(usageControlChain)

	sessionID := input.RuntimeSessionID
	if usageControlChainIsUpdate(usageControlChainData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": usageControlChainData, "function": "UsageControlChainUsageControlChain", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "UsageControlChain Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToUsageControlChainUpdates(*usageControlChainData)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func usageControlChainIsUpdate(usageControlChain *dpfm_api_processing_formatter.UsageControlChainUpdates) bool {
	usageControlChainId := usageControlChain.UsageControlChain

	return !(usageControlChainId == "")
}
