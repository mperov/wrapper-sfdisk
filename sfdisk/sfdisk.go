package sfdisk

import (
    "os/exec"
    "bytes"
)

func Run(deviceName string) string {
    cmd := exec.Command("/sbin/sfdisk", deviceName)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    summary := stdout.String()
    if err != nil {
        summary = stderr.String()
    }
    return summary
}
