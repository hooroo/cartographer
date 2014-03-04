package main

import (
  "github.com/hooroo/cartographer/report"
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
)

func whereami() string {
  exec_name := os.Args[0]
  find_me_cmd := exec.Command("which", exec_name)
  where_am_i, _ := find_me_cmd.Output()
  where_am_i_dir := filepath.Dir(string(where_am_i))
  if filepath.IsAbs(where_am_i_dir) {
    return where_am_i_dir
  }
  cwd, _ := os.Getwd()

  return filepath.Join(cwd, where_am_i_dir)
}



func main() {
  arb_reps := make(map[string]interface{})

  plugins_dir := filepath.Join(whereami(), "run")
  plugins_globstr := fmt.Sprintf("%s/*", plugins_dir)

  plugins, _ := filepath.Glob(plugins_globstr)

  for _, plugin := range plugins {
    f, _ := os.Open(plugin)
    fstat, _ := f.Stat()
    f.Close()
    if !fstat.IsDir() && (fstat.Mode() | 0111) != 0 {
      plugin_name := filepath.Base(plugin)
      plugin_output, _ := exec.Command(plugin).Output()
      var plugin_json interface{}
      json.Unmarshal(plugin_output, &plugin_json)
      arb_reps[plugin_name] = plugin_json
    }
  }

  sys := report.NewSystemReport()
  rep := report.Report{sys, arb_reps}
  strrep, _ := json.Marshal(rep)

  fmt.Println(string(strrep))
}
