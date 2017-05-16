package sleephandle

import (
    "log"
    "os/exec"
)

func SleepHandle() {
    cmd_wifi_down := exec.Command("/bin/sh", "-c", "ifconfig wlan0 down")
    output, err := cmd_wifi_down.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }
    log.Println(string(output))

    cmd_ra_down := exec.Command("/bin/sh", "-c", "/etc/init.d/ra stop")
    _, err := cmd_ra_down.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }

    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, err := cmd_wake_lock.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }

    cmd_mem_down := exec.Command("/bin/sh", "-c", "echo mem > /sys/power/state")
    _, err := cmd_mem_down.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }

    cmd_wake_unlock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_unlock")
    _, err := cmd_wake_unlock.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }

}
