package syssleepwake

// Package name is smartconn.cc/sibolwolf/syssleepwake

import (
    "log"
    "time"
    "os/exec"
    SH "smartconn.cc/sibolwolf/syssleepwake/sleephandle"
    WH "smartconn.cc/sibolwolf/syssleepwake/wakehandle"
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
    log.Println("---------------------------------------")
    log.Printf("Current lock status is: %d\n", currlocksum)
    log.Println("audiolock:",           lockstatus["audiolock"])
    log.Println("cameralock:",          lockstatus["cameralock"])
    log.Println("storydecompresslock",  lockstatus["storydecompresslock"])
    log.Println("storydownloadlock",    lockstatus["storydownloadlock"])
    log.Println("storysynclock",        lockstatus["storysynclock"])
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
        log.Println(cntdown)
        cntdown -= 1
        if cntdown == 0 {
            cntdown = cntdownsum
            log.Println("Start sleep ...")
            // Action before sleep
            SH.SleepHandle()

            time.Sleep(time.Second * 1)

            // Action after sleep
            log.Println("Back from sleep ...")
            WH.WakeHandle()

        }

        time.Sleep(time.Second * 1)
    }
}

func Init() {
    ClearLockStatus()
    UpdateLockStatus("init", 0)
    go ContinueCnt()
}
