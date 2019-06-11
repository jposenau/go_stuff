package main

import (
    "fmt"
    _ "net/http/pprof"
    "sync"
    "time"
)

func main() {
    chana := make(chan int)
    chanb := make(chan int)

    go func() {
        for i := 0; i < 1000; i++ {
            chana <- 100 * i
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            chanb <- i
        }
    }()

    time.Sleep(time.Microsecond * 300)

    acount := 0
    bcount := 0
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        for {
            select {
            case <-chana:
                acount++
            case <-chanb:
                bcount++
            }
            if acount == 1000 || bcount == 1000 {
                fmt.Println("finish one acount, bcount", acount, bcount)
                break
            }
        }
        wg.Done()
    }()

    wg.Wait()
}