package main

import (
  "github.com/hooroo/cartographer/report"
  "encoding/json"
  "fmt"
  "net/http"
  "bytes"
)

func generateReports() report.Report {
  external_reports := report.ExternalReports("run")
  system_report := report.NewSystemReport()
  return report.Report{system_report, external_reports}
}

func main() {
  reports := generateReports()
  reports_json, _ := json.Marshal(reports)

  reports_json  = append(reports_json, []byte("\r\n\r\n")... )

  fmt.Println(string(reports_json))
  reader := bytes.NewReader(reports_json)
  resp, _ := http.Post("http://localhost:5050/", "application/json", reader)

  resp.Body.Close()
  return
}
