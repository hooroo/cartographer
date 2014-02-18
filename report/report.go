package report

type SystemReport struct {
	NodeName string `json:"node_name"`
	IPAddress string `json:"ip_address"`
}

type Report struct {
	System SystemReport `json:"system"`
	Reports map[string]interface{} `json:"reports"`
}