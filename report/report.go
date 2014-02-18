package report

type Report struct {
  System SystemReport `json:"system"`
  Reports map[string]interface{} `json:"reports"`
}
