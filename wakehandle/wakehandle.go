package wakehandle

import (
    "log"
    "os/exec"
)

func WakeHandle() {
    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, err := cmd_wake_lock.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }

    cmd_wifi_up := exec.Command("/bin/sh", "-c", "ifconfig wlan0 up")
    output, err := cmd_wake_lock.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }
    log.Println(string(output))

    cmd_ra_up := exec.Command("/bin/sh", "-c", "/etc/init.d/ra start")
    _, err := cmd_ra_up.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }
}
