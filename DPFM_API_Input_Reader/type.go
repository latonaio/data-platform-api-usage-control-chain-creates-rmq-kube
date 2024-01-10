package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string            `json:"connection_key"`
	Result            bool              `json:"result"`
	RedisKey          string            `json:"redis_key"`
	Filepath          string            `json:"filepath"`
	APIStatusCode     int               `json:"api_status_code"`
	RuntimeSessionID  string            `json:"runtime_session_id"`
	BusinessPartner   *int              `json:"business_partner"`
	ServiceLabel      string            `json:"service_label"`
	APIType           string            `json:"api_type"`
	UsageControlChain UsageControlChain `json:"UsageControl"`
	APISchema         string            `json:"api_schema"`
	Accepter          []string          `json:"accepter"`
	Deleted           bool              `json:"deleted"`
}
type LocationRestriction struct {
	PostalCode     []*string `json:"PostalCode"`
	LocalSubRegion []*string `json:"LocalSubRegion"`
	LocalRegion    []*string `json:"LocalRegion"`
	Country        []*string `json:"Country"`
	GlobalRegion   []*string `json:"GlobalRegion"`
	TimeZone       []*string `json:"TimeZone"`
}

type UsageControlChain struct {
	UsageControlChain              string              `json:"UsageControlChain"`
	UsageControlLess               *bool               `json:"UsageControlLess"`
	Perpetual                      *bool               `json:"Perpetual"`
	OneTime                        *bool               `json:"OneTime"`
	Rental                         *bool               `json:"Rental"`
	Duration                       *float64            `json:"Duration"`
	DurationUnit                   *string             `json:"DurationUnit"`
	ValidityStartDate              *string             `json:"ValidityStartDate"`
	ValidityStartTime              *string             `json:"ValidityStartTime"`
	ValidityEndDate                *string             `json:"ValidityEndDate"`
	ValidityEndTime                *string             `json:"ValidityEndTime"`
	DeleteAfterValidityEnd         *bool               `json:"DeleteAfterValidityEnd"`
	ServiceLabelRestriction        []*string           `json:"ServiceLabelRestriction"`
	ApplicationRestriction         []*string           `json:"ApplicationRestriction"`
	PurposeRestriction             []*string           `json:"PurposeRestriction"`
	BusinessPartnerRoleRestriction []*string           `json:"BusinessPartnerRoleRestriction"`
	DataStateRestriction           []*string           `json:"DataStateRestriction"`
	NumberOfUsageRestriction       *int                `json:"NumberOfUsageRestriction"`
	NumberOfActuaUsage             *int                `json:"NumberOfActuaUsage"`
	IPAddressRestriction           []*string           `json:"IPAddressRestriction"`
	MACAddressRestriction          []*string           `json:"MACAddressRestriction"`
	ModifyIsAllowed                *bool               `json:"ModifyIsAllowed"`
	LocationRestriction            LocationRestriction `json:"LocationRestriction"`
	LocalLoggingIsAllowed          *bool               `json:"LocalLoggingIsAllowed"`
	RemoteNotificationIsAllowed    *string             `json:"RemoteNotificationIsAllowed"`
	DestributeOnlyIfEncrypted      *bool               `json:"DestributeOnlyIfEncrypted"`
	PostalCode                     []*string           `json:"PostalCode"`
	LocalSubRegion                 []*string           `json:"LocalSubRegion"`
	LocalRegion                    []*string           `json:"LocalRegion"`
	Country                        []*string           `json:"Country"`
	GlobalRegion                   []*string           `json:"GlobalRegion"`
	TimeZone                       []*string           `json:"TimeZone"`
	CreationDate                   *string             `json:"CreationDate"`
	CreationTime                   *string             `json:"CreationTime"`
	LastChangeDate                 *string             `json:"LastChangeDate"`
	LastChangeTime                 *string             `json:"LastChangeTime"`
	IsMarkedForDeletion            *bool               `json:"IsMarkedForDeletion"`
}
