package main

import (
  "encoding/json"
  "fmt"
  "strings"

  "github.com/bwilkins/processes"
)

type UnicornReport struct {
  Workers int `json:"workers"`
  MemoryUsed int `json:"memory_used"`
}

func findUnicorns() processes.PsList {
  entries := processes.Ps()
  unicorns := make(processes.PsList, 0, 2)
  for _, entry := range entries {
    if strings.Contains(entry.Command, "unicorn_rails") {
      unicorns = append(unicorns, entry)
    }
  }
  return unicorns
}

func main() {
  workers := 0
  memory_used := 0

  for _, unicorn := range findUnicorns() {
    if strings.Contains(unicorn.Command, "worker") {
      workers += 1
      memory_used += unicorn.ResidentMemory
    }
  }

  report := UnicornReport{workers, memory_used}

  j, _ := json.Marshal(report)
  fmt.Println(string(j))
}
