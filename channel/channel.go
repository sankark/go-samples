package main

import "fmt"
import "time"

func worker2(done,done2 chan bool) {
    <- done
    fmt.Print("worker2...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done2 <- true
}

func worker1(done chan bool) {
    fmt.Print("worker1...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    done := make(chan bool, 1)
    done2 := make(chan bool, 1)
    go worker1(done)
    go worker2(done,done2)
    // Block until we receive a notification from the
    // worker on the channel.
    <-done2
}

