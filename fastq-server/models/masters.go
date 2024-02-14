package models

type Countries struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}
type TimeZones struct {
	ZoneName     string `json:"zone_name"`
	Abbreviation string `json:"abbreviation"`
	CountryCode  string `json:"country_code"`
	TimeStart    string `json:"time_start"`
	GMT          string `json:"gmt_offset"`
}

type MasterResult struct {
	Countries []*Countries `json:"countries"`
	TimeZones []*TimeZones `json:"timezones"`
}
