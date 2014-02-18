package main

import (
	"github.com/bwilkins/cartographer/report"
	"encoding/json"
	"fmt"
)

func main() {
	arb_reps := make(map[string]interface{})
	sys := report.SystemReport{"node1", "127.0.0.1"}
	rep := report.Report{sys, arb_reps}

	strrep, _ := json.Marshal(rep)

	fmt.Println(string(strrep))
}