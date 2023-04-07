package models

type Clients struct {
	Id           string
	Dt           uint64
	ClientId     uint64
	Type         string
	SubmitId     string
	Referer      string
	Os           string
	LeadSource   string
	CreativeName string
	Country      string
}

type ClientUuid struct {
	ClientId uint64
	Uuid     string
}

type ApiResponse struct {
	Id           string `json:"id"`
	Dt           string `json:"dt"`
	ClientId     string `json:"client_id"`
	Type         string `json:"type"`
	SubmitId     string `json:"submit_id"`
	Referer      string `json:"referer"`
	Os           string `json:"os"`
	LeadSource   string `json:"lead_source"`
	CreativeName string `json:"creative_name"`
	Country      string `json:"country"`
}
