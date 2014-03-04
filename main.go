package main

import (
  "github.com/hooroo/cartographer/report"
  "encoding/json"
  "fmt"
  "net/http"
  "bytes"
  "flag"
)

func generateReports(external_reports_dir string) report.Report {
  external_reports := report.ExternalReports(external_reports_dir)
  system_report := report.NewSystemReport()
  return report.Report{system_report, external_reports}
}

func main() {
  var external_reports_dir = flag.String("external-reports_dir", "run", "Directory in which to look for external reports")
  var http_report_to = flag.String("http-report-to", "", "URL to send report to")
  var print_to_stdout = flag.Bool("print-to-stdout", true, "Print output to stdout?")
/*  var daemonise = flag.Bool("daemonise", false, "Process becomes long-running (instead of run-once and exit)")*/
  flag.Parse()

  reports := generateReports(*external_reports_dir)
  reports_json, _ := json.Marshal(reports)

  if *print_to_stdout {
    fmt.Println(string(reports_json))
  }

  if *http_report_to != "" {
    reader := bytes.NewReader(reports_json)
    resp, _ := http.Post(*http_report_to, "application/json", reader)
    resp.Body.Close()
  }
  return
}
