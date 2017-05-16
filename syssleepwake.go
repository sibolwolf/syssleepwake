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
var cntdown int = 300

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
            cntdown = 300   // 300s
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
            cntdown = 300
            return
        }
        fmt.Println(cntdown)
        cntdown -= 1
        if cntdown == 0 {
            cntdown = 300
            // Action before sleep
            exec.Command("/bin/sh", "-c", "ifconfig wlan0 down")
            exec.Command("/bin/sh", "-c", "/etc/init.d/ra stop")
            exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
            exec.Command("/bin/sh", "-c", "echo mem > /sys/power/state")
            exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_unlock")

            time.Sleep(time.Second * 1)

            // Action after sleep
            exec.Command("/bin/sh", "-c", "echo test > /sys/power/wake_lock")
            exec.Command("/bin/sh", "-c", "ifconfig wlan0 up")
            exec.Command("/bin/sh", "-c", "/etc/init.d/ra start")
        }

        time.Sleep(time.Second * 1)
    }
}

func Init() {
    ClearLockStatus()
    UpdateLockStatus("init", 0)
    go ContinueCnt()
}
