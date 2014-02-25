package main

import (
  "encoding/json"
  "fmt"
  "strings"

  "github.com/bwilkins/processes"
)

type NginxReport struct {
  Workers int `json:"workers"`
  MemoryUsed int `json:"memory_used"`
}

func findNginxs() processes.PsList {
  entries := processes.Ps()
  nginxs := make(processes.PsList, 0, 2)
  for _, entry := range entries {
    if strings.Contains(entry.Command, "nginx") {
      nginxs = append(nginxs, entry)
    }
  }
  return nginxs
}

func main() {
  workers := 0
  memory_used := 0

  for _, nginx := range findNginxs() {
    if strings.Contains(nginx.Command, "worker") {
      workers += 1
      memory_used += nginx.ResidentMemory
    }
  }

  report := NginxReport{workers, memory_used}

  j, _ := json.Marshal(report)
  fmt.Println(string(j))
}
