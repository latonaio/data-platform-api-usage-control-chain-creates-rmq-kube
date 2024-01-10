package requests

type UsageControlChain struct {
	UsageControlChain              string   `json:"UsageControlChain"`
	UsageControlLess               *bool    `json:"UsageControlLess"`
	Perpetual                      *bool    `json:"Perpetual"`
	OneTime                        *bool    `json:"OneTime"`
	Rental                         *bool    `json:"Rental"`
	Duration                       *float64 `json:"Duration"`
	DurationUnit                   *string  `json:"DurationUnit"`
	ValidityStartDate              *string  `json:"ValidityStartDate"`
	ValidityStartTime              *string  `json:"ValidityStartTime"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	ValidityEndTime                *string  `json:"ValidityEndTime"`
	DeleteAfterValidityEnd         *bool    `json:"DeleteAfterValidityEnd"`
	ServiceLabelRestriction        *bool    `json:"ServiceLabelRestriction"`
	ApplicationRestriction         *bool    `json:"ApplicationRestriction"`
	PurposeRestriction             *bool    `json:"PurposeRestriction"`
	BusinessPartnerRoleRestriction *bool    `json:"BusinessPartnerRoleRestriction"`
	DataStateRestriction           *bool    `json:"DataStateRestriction"`
	NumberOfUsageRestriction       *int     `json:"NumberOfUsageRestriction"`
	NumberOfActuaUsage             *int     `json:"NumberOfActuaUsage"`
	IPAddressRestriction           *bool    `json:"IPAddressRestriction"`
	MACAddressRestriction          *bool    `json:"MACAddressRestriction"`
	ModifyIsAllowed                *bool    `json:"ModifyIsAllowed"`
	LocalLoggingIsAllowed          *bool    `json:"LocalLoggingIsAllowed"`
	RemoteNotificationIsAllowed    *bool    `json:"RemoteNotificationIsAllowed"`
	DestributeOnlyIfEncrypted      *bool    `json:"DestributeOnlyIfEncrypted"`
	PostalCode                     *string  `json:"PostalCode"`
	LocalSubRegion                 *string  `json:"LocalSubRegion"`
	LocalRegion                    *string  `json:"LocalRegion"`
	Country                        *string  `json:"Country"`
	GlobalRegion                   *string  `json:"GlobalRegion"`
	TimeZone                       *string  `json:"TimeZone"`
	CreationDate                   *string  `json:"CreationDate"`
	CreationTime                   *string  `json:"CreationTime"`
	LastChangeDate                 *string  `json:"LastChangeDate"`
	LastChangeTime                 *string  `json:"LastChangeTime"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
