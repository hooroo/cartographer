package report

type SystemReport struct {
  NodeName string `json:"node_name"`
  IPAddress string `json:"ip_address"`
}

func nodeName() string {

  return "name1"
}

func ipAddress() string {
  return "127.0.0.1"
}

func NewSystemReport() SystemReport {
  return SystemReport{nodeName(), ipAddress()}
}