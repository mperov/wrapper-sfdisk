package sfdisk

import (
    "os/exec"
)

func Run(deviceName string) string {
    return exec.Command("/bin/sfdisk", deviceName)
}
