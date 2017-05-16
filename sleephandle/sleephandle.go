package sleephandle

import (
    "log"
    "os/exec"
)

func SleepHandle() {
    cmd_wifi_down := exec.Command("/bin/sh", "-c", "ifconfig wlan0 down")
    cmd_wifi_down_output, cmd_wifi_down_err := cmd_wifi_down.Output()
    if cmd_wifi_down_err != nil {
        log.Println("cmd_wifi_down_err: " + cmd_wifi_down_err.Error())
        return
    }
    log.Println(string(cmd_wifi_down_output))

    cmd_ra_down := exec.Command("/bin/sh", "-c", "/etc/init.d/ra stop")
    _, cmd_ra_down_err := cmd_ra_down.Output()
    if cmd_ra_down_err != nil {
        log.Println("cmd_ra_down_err: " + cmd_ra_down_err.Error())
        return
    }

    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, cmd_wake_lock_err := cmd_wake_lock.Output()
    if cmd_wake_lock_err != nil {
        log.Println("cmd_wake_lock_err: " + cmd_wake_lock_err.Error())
        return
    }

    cmd_mem_down := exec.Command("/bin/sh", "-c", "echo mem > /sys/power/state")
    _, cmd_mem_down_err := cmd_mem_down.Output()
    if cmd_mem_down_err != nil {
        log.Println("cmd_mem_down_err: " + cmd_mem_down_err.Error())
        return
    }

    cmd_wake_unlock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_unlock")
    _, cmd_wake_unlock_err := cmd_wake_unlock.Output()
    if cmd_wake_unlock_err != nil {
        log.Println("cmd_wake_unlock_err: " + cmd_wake_unlock_err.Error())
        return
    }

}
