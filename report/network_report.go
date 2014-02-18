package report

import (
  "net"
  "strings"
)

type NetworkReport struct {
  Addresses  []string `json:"addresses"`
}

func NewNetworkReport() NetworkReport {
  addrs := make([]string, 0, 3)
  intaddrs, _ := net.InterfaceAddrs()
  for _, addr := range intaddrs {
    addr_string := strings.Split(addr.String(), "/")[0]
    addrs = append(addrs, addr_string)
  }
  return NetworkReport{addrs}
}
