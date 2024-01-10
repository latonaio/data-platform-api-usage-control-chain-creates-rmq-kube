package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-usage-control-chain-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToUsageControlChainCreates(
	subfuncSDC *sub_func_complementer.SDC,
) (*UsageControlChain, error) {
	data := subfuncSDC.Message.UsageControlChain

	usageControlChain, err := TypeConverter[*UsageControlChain](data)
	if err != nil {
		return nil, err
	}

	return usageControlChain, nil
}

func ConvertToUsageControlChainUpdates(
	usageControlChainData dpfm_api_processing_formatter.UsageControlChainUpdates,
) (*UsageControlChain, error) {
	data := usageControlChainData

	header, err := TypeConverter[*UsageControlChain](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
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
