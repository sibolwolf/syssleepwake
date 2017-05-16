package wakehandle

import (
    "log"
    "os/exec"
    KEY "smartconn.cc/liugen/input"
)

func WakeHandle() {
    KEY.Connect("readingangel")
    KEY.GetButton("power").OnPress(func() {
            log.Println("RA got a short key press event for power")
            WakeHandleAction()
        })

    KEY.GetButton("home").OnPress(func() {
            log.Println("RA got a short key press event for home")
            WakeHandleAction()
        })
}

func WakeHandleAction() {
    log.Println("#1. SleepHandle: Echo test > wake_lock ..........")
    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, cmd_wake_lock_err := cmd_wake_lock.Output()
    if cmd_wake_lock_err != nil {
        log.Println("cmd_wake_lock_err: " + cmd_wake_lock_err.Error())
    }

    log.Println("#2. SleepHandle: Open Wifi ..........")
    cmd_wifi_up := exec.Command("/bin/sh", "-c", "ifconfig wlan0 up")
    _, cmd_wifi_up_err := cmd_wifi_up.Output()
    if cmd_wifi_up_err != nil {
        log.Println("cmd_wifi_up_err: " + cmd_wifi_up_err.Error())
    }
    //log.Println(string(cmd_wifi_up_output))

    log.Println("#3. SleepHandle: Start RA ..........")
    cmd_ra_up := exec.Command("/bin/sh", "-c", "/etc/init.d/ra start")
    _, cmd_ra_up_err := cmd_ra_up.Output()
    if cmd_ra_up_err != nil {
        log.Println("cmd_ra_up_err: " + cmd_ra_up_err.Error())
    }
    //log.Println(string(cmd_ra_up_output))
}
