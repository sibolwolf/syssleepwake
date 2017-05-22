package wakehandle

import (
    "log"
    "os/exec"
    "time"
    "strings"
    KEY "smartconn.cc/liugen/input"
)

var wakestatus int



func WakeJudgment() {
    cmd := exec.Command("/bin/sh", "-c", "sysint getSysUnlockStatus")
    bytes, err := cmd.Output()
    if err != nil {
        log.Println("Powerbutton get a fault when wake up: " + err.Error())
    } else {
        sysunlockstatus := strings.TrimSpace(string(bytes))
        if sysunlockstatus == "test" {
            WakeHandleAction()
        }
    }
}
// Power button trig wake up
func WakeHandlePowerButton() {
    log.Println("----------------------------------------------------")
    log.Println("Power button trig wake up")
    var deconnectpower func()
    KEY.Connect("readingangel")
    deconnectpower = KEY.GetButton("power").OnPress(func() {
            log.Println("RA got a short key press event for power")
            WakeJudgment()
            deconnectpower()
        })
}

func WakeHandle() {
    go WakeHandlePowerButton()
}

func WakeHandleAction() {
    log.Println("Back from sleep ...")
    log.Println("#1. WakeHandle: Echo test > wake_lock ..........")
    cmd_wake_lock := exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
    _, cmd_wake_lock_err := cmd_wake_lock.Output()
    if cmd_wake_lock_err != nil {
        log.Println("cmd_wake_lock_err: " + cmd_wake_lock_err.Error())
    }


    log.Println("#2. WakeHandle: Open Wifi ..........")
    cmd_wifi_up := exec.Command("/bin/sh", "-c", "ifconfig wlan0 up")
    _, cmd_wifi_up_err := cmd_wifi_up.Output()
    if cmd_wifi_up_err != nil {
        log.Println("cmd_wifi_up_err: " + cmd_wifi_up_err.Error())
    }
    //log.Println(string(cmd_wifi_up_output))

    // Sleep 1 Second
    time.Sleep(time.Second * 1)
    log.Println("#3. WakeHandle: Start RA ..........")
    cmd_ra_up := exec.Command("/bin/sh", "-c", "/etc/init.d/ra start")
    _, cmd_ra_up_err := cmd_ra_up.Output()
    if cmd_ra_up_err != nil {
        log.Println("cmd_ra_up_err: " + cmd_ra_up_err.Error())
    }
    //log.Println(string(cmd_ra_up_output))
}
