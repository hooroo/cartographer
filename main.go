package main

import (
  "github.com/bwilkins/cartographer/report"
  "encoding/json"
  "fmt"
)

func main() {
  arb_reps := make(map[string]interface{})
  sys := report.NewSystemReport()
  rep := report.Report{sys, arb_reps}
  strrep, _ := json.Marshal(rep)

  fmt.Println(string(strrep))
}
