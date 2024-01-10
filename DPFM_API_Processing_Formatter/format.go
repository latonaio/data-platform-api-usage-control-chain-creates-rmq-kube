package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-usage-control-chain-creates-rmq-kube/DPFM_API_Input_Reader"
	"strings"
)

func ConvertToUsageControlChainUpdates(
	usageControlChain dpfm_api_input_reader.UsageControlChain,
) *UsageControlChainUpdates {
	data := usageControlChain

	quoteAndJoin := func(values []*string) *string {
		if len(values) == 0 {
			return nil
		}

		var result strings.Builder

		for i, v := range values {
			result.WriteString(*v)
			if i < len(values)-1 {
				result.WriteString(`, `)
			}
		}

		concatenatedValues := result.String()
		return &concatenatedValues
	}

	return &UsageControlChainUpdates{
		UsageControlChain:              data.UsageControlChain,
		UsageControlLess:               data.UsageControlLess,
		Perpetual:                      data.Perpetual,
		Rental:                         data.Rental,
		Duration:                       data.Duration,
		DurationUnit:                   data.DurationUnit,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityStartTime:              data.ValidityStartTime,
		ValidityEndDate:                data.ValidityEndDate,
		ValidityEndTime:                data.ValidityEndTime,
		DeleteAfterValidityEnd:         data.DeleteAfterValidityEnd,
		ServiceLabelRestriction:        quoteAndJoin(data.ServiceLabelRestriction),
		ApplicationRestriction:         quoteAndJoin(data.ApplicationRestriction),
		PurposeRestriction:             quoteAndJoin(data.PurposeRestriction),
		BusinessPartnerRoleRestriction: quoteAndJoin(data.BusinessPartnerRoleRestriction),
		DataStateRestriction:           quoteAndJoin(data.DataStateRestriction),
		NumberOfUsageRestriction:       data.NumberOfUsageRestriction,
		NumberOfActuaUsage:             data.NumberOfActuaUsage,
		IPAddressRestriction:           quoteAndJoin(data.IPAddressRestriction),
		MACAddressRestriction:          quoteAndJoin(data.MACAddressRestriction),
		ModifyIsAllowed:                data.ModifyIsAllowed,
		LocalLoggingIsAllowed:          data.LocalLoggingIsAllowed,
		RemoteNotificationIsAllowed:    data.RemoteNotificationIsAllowed,
		DestributeOnlyIfEncrypted:      data.DestributeOnlyIfEncrypted,
		PostalCode:                     quoteAndJoin(data.PostalCode),
		LocalSubRegion:                 quoteAndJoin(data.LocalSubRegion),
		LocalRegion:                    quoteAndJoin(data.LocalRegion),
		Country:                        quoteAndJoin(data.Country),
		GlobalRegion:                   quoteAndJoin(data.GlobalRegion),
		TimeZone:                       quoteAndJoin(data.TimeZone),
		CreationDate:                   *data.CreationDate,
		CreationTime:                   *data.CreationTime,
		LastChangeDate:                 *data.LastChangeDate,
		LastChangeTime:                 *data.LastChangeTime,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}
