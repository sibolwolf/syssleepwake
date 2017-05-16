package main

import (
    "log"
    "os"
    "os/signal"
    "time"
    SysSW "smartconn.cc/sibolwolf/syssleepwake"
)

func SysSWTest() {
    // Nothing to do
    SysSW.UpdateLockStatus("audiolock", 0)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 1)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 0)
    time.Sleep(time.Second * 1)
    SysSW.UpdateLockStatus("storysynclock", 1)
    time.Sleep(time.Second * 1)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 0)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
}

func main() {
    log.Println("Hello, SysSW")
    SysSW.Init()
    go SysSWTest()

    signalChanel := make(chan os.Signal, 1)
    signal.Notify(signalChanel, os.Interrupt)
    for {
        select {
        case <-signalChanel:
            return
        }
    }
}
