package sfdisk

import (
    "os/exec"
)

func Run(deviceName string) string {
    exec.Command("/bin/sfdisk", deviceName)
    return "done"
}
