package syssleepwake

// Package name is smartconn.cc/sibolwolf/syssleepwake

import (
    "fmt"
    "time"
    "os/exec"
)

/*
cameralock
audiolock
storydownloadlock
storydecompresslock
storysynclock
*/

var lockstatus = map[string]int{
    "audiolock": 0,
    "cameralock": 0,
    "storydecompresslock": 0,
    "storydownloadlock": 0,
    "storysynclock": 0,
}

var initlocksum int = 0
var currlocksum int = 0
var lastlocksum int = 0
var cntdownsum int = 50
var cntdown int = 50

func ShowLockStatus() {
    fmt.Println("---------------------------------------")
    fmt.Printf("Current lock status is: %d\n", currlocksum)
    fmt.Println("audiolock:",           lockstatus["audiolock"])
    fmt.Println("cameralock:",          lockstatus["cameralock"])
    fmt.Println("storydecompresslock",  lockstatus["storydecompresslock"])
    fmt.Println("storydownloadlock",    lockstatus["storydownloadlock"])
    fmt.Println("storysynclock",        lockstatus["storysynclock"])
}

func ClearLockStatus() {
    for k := range lockstatus {
        lockstatus[k] = 0
    }
}

func UpdateLockStatus(key string, value int) {
    if key == "init" {
        currlocksum = value
    } else {
        lockstatus[key] = value
        locksum := 0
        for _, v := range lockstatus {
            locksum += v
        }
        currlocksum = locksum
    }

    if currlocksum != lastlocksum {

        ShowLockStatus()
        lastlocksum = currlocksum
    }

    time.Sleep(500)
}

func ContinueCnt(){
    for {
        if currlocksum != initlocksum {
            cntdown = cntdownsum   // 300s
        }

        if currlocksum == initlocksum {
            SleepWakeHandle()
        }
    }
}

func SleepWakeHandle() {
    // When system go to sleep, there is something need to do
    for cntdown > 0{
        if currlocksum != initlocksum {
            cntdown = cntdownsum
            return
        }
        fmt.Println(cntdown)
        cntdown -= 1
        if cntdown == 0 {
            cntdown = cntdownsum
            fmt.Println("Start sleep ...")
            // Action before sleep
            fmt.Println(exec.Command("/bin/sh", "-c", "ifconfig wlan0 down").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "/etc/init.d/ra stop").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "echo mem > /sys/power/state").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_unlock").Output())

            time.Sleep(time.Second * 1)

            // Action after sleep
            fmt.Println("Back from sleep ...")
            fmt.Println(exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "ifconfig wlan0 up").Output())
            fmt.Println(exec.Command("/bin/sh", "-c", "/etc/init.d/ra start").Output())
        }

        time.Sleep(time.Second * 1)
    }
}

func Init() {
    ClearLockStatus()
    UpdateLockStatus("init", 0)
    go ContinueCnt()
}
