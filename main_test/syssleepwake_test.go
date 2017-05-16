package main

import (
    "fmt"
    "os"
    "os/signal"
    SysSW "smartconn.cc/sibolwolf/syssleepwake"
)

func main() {
    SysSW.Init()
    SysSw.testLockStatus2()

    signalChanel := make(chan os.Signal, 1)
    signal.Notify(signalChanel, os.Interrupt)
    for {
        select {
        case <-signalChanel:
            return
        }
    }
}
