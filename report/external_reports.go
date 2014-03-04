package report

import (
  "path/filepath"
  "os"
  "os/exec"
  "encoding/json"
  "fmt"
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

func ExternalReports(folder string) map[string]interface{} {
  external_reports := make(map[string]interface{})

  if filepath.IsAbs(folder) {
    external_reports_dir := folder
  } else {
    external_reports_dir := filepath.Join(whereami(), folder)
  }

  external_reports_globstr := fmt.Sprintf("%s/*", external_reports_dir)

  external_report_runners, _ := filepath.Glob(external_reports_globstr)

  for _, report_runner := range external_report_runners {
    f, _ := os.Open(report_runner)
    fstat, _ := f.Stat()
    f.Close()
    if !fstat.IsDir() && (fstat.Mode() | 0111) != 0 {
      runner_name := filepath.Base(report_runner)
      runner_output, _ := exec.Command(report_runner).Output()
      var report_json interface{}
      json.Unmarshal(runner_output, &report_json)
      external_reports[runner_name] = report_json
    }
  }

  return external_reports
}
