package wakehandle

import (
    "log"
    "os/exec"
)

func WakeHandle() {
    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, cmd_wake_lock_err := cmd_wake_lock.Output()
    if cmd_wake_lock_err != nil {
        log.Println("cmd_wake_lock_err: " + cmd_wake_lock_err.Error())
        return
    }

    cmd_wifi_up := exec.Command("/bin/sh", "-c", "ifconfig wlan0 up")
    cmd_wifi_up_output, cmd_wifi_up_err := ccmd_wifi_up.Output()
    if cmd_wifi_up_err != nil {
        log.Println("cmd_wifi_up_err: " + cmd_wifi_up_err.Error())
        return
    }
    log.Println(string(cmd_wifi_up_output))

    cmd_ra_up := exec.Command("/bin/sh", "-c", "/etc/init.d/ra start")
    cmd_ra_up_output, cmd_ra_up_err := cmd_ra_up.Output()
    if cmd_ra_up_err != nil {
        log.Println("cmd_ra_up_err: " + cmd_ra_up_err.Error())
        return
    }
    log.Println(string(cmd_ra_up_output))
}
