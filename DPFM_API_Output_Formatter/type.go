package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	UsageControlChain *UsageControlChain `json:"UsageControl"`
}

type UsageControlChain struct {
	UsageControlChain              string   `json:"UsageControlChain"`
	UsageControlLess               *bool    `json:"UsageControlLess"`
	Perpetual                      *bool    `json:"Perpetual"`
	Rental                         *bool    `json:"Rental"`
	Duration                       *float64 `json:"Duration"`
	DurationUnit                   *string  `json:"DurationUnit"`
	ValidityStartDate              *string  `json:"ValidityStartDate"`
	ValidityStartTime              *string  `json:"ValidityStartTime"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	ValidityEndTime                *string  `json:"ValidityEndTime"`
	DeleteAfterValidityEnd         *bool    `json:"DeleteAfterValidityEnd"`
	ServiceLabelRestriction        *string  `json:"ServiceLabelRestriction"`
	ApplicationRestriction         *string  `json:"ApplicationRestriction"`
	PurposeRestriction             *string  `json:"PurposeRestriction"`
	BusinessPartnerRoleRestriction *string  `json:"BusinessPartnerRoleRestriction"`
	DataStateRestriction           *string  `json:"DataStateRestriction"`
	NumberOfUsageRestriction       *int     `json:"NumberOfUsageRestriction"`
	NumberOfActualUsage            *int     `json:"NumberOfActualUsage"`
	IPAddressRestriction           *string  `json:"IPAddressRestriction"`
	MACAddressRestriction          *string  `json:"MACAddressRestriction"`
	ModifyIsAllowed                *bool    `json:"ModifyIsAllowed"`
	LocalLoggingIsAllowed          *bool    `json:"LocalLoggingIsAllowed"`
	RemoteNotificationIsAllowed    *string  `json:"RemoteNotificationIsAllowed"`
	DistributeOnlyIfEncrypted      *bool    `json:"DistributeOnlyIfEncrypted"`
	AttachPolicyWhenDistribute     *bool    `json:"AttachPolicyWhenDistribute"`
	PostalCode                     *string  `json:"PostalCode"`
	LocalSubRegion                 *string  `json:"LocalSubRegion"`
	LocalRegion                    *string  `json:"LocalRegion"`
	Country                        *string  `json:"Country"`
	GlobalRegion                   *string  `json:"GlobalRegion"`
	TimeZone                       *string  `json:"TimeZone"`
	CreationDate                   string   `json:"CreationDate"`
	CreationTime                   string   `json:"CreationTime"`
	LastChangeDate                 string   `json:"LastChangeDate"`
	LastChangeTime                 string   `json:"LastChangeTime"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
