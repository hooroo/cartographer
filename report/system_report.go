package report

import (
  "os/exec"
  "bytes"
  "strings"
  "net"
  "regexp"
)

type SystemReport struct {
  NodeName string `json:"node_name"`
  IPAddress string `json:"ip_address"`
  Network NetworkReport `json:"network"`
}

func nodeName() string {
  name_cmd := exec.Command("uname", "-n")
  name, _ := name_cmd.Output()
  return string(bytes.TrimSpace(name))
}

func ipRegex() (reg *regexp.Regexp) {
  reg, _ = regexp.Compile(`\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}`)
  return
}

func ipAddress() string {
  addrs, _ := net.InterfaceAddrs()
  for _, addr := range addrs {
    addr_string := strings.Split(addr.String(), "/")[0]
    if ipRegex().Match([]byte(addr_string)) && addr_string != "127.0.0.1" {
      return addr_string
    }
  }
  return "127.0.0.1"
}

func NewSystemReport() SystemReport {
  return SystemReport{nodeName(), ipAddress(), NewNetworkReport()}
}
