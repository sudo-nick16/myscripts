package main

import (
	"log"
	"os/exec"
	"strings"
)

func main() {
  out, err := exec.Command("pacmd", "list-sinks").Output()
  if err != nil {
    log.Fatalf("SoundError: %v", err)
  }
  outS := strings.Split(string(out), "\n")
  var index string
  for _, line := range outS {
    if strings.Contains(line, "index:") {
      index = strings.Split(line, ":")[1]
      index = strings.TrimSpace(index)
    }
    if strings.Contains(line, "analog-output-speaker") {
      break
    }
  }
  if index == "" {
    log.Fatalf("SoundError: no index found")
  }
  err1 := exec.Command("pacmd", "set-default-sink", index).Run()
  if err1 != nil {
    log.Fatalf("SoundError setting default sink: %v", err)
  }
  log.Printf("Set default sink to %v", index)
  err2 := exec.Command("pacmd", "set-default-source", "2").Run()
  if err2 != nil {
    log.Fatalf("SoundError setting default source: %v", err)
  }
  log.Printf("Set default source to %v", index)
}
