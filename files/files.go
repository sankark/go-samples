package main

import (
    "os"
    "io/ioutil"
	"fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    f, err := os.Create("/tmp/dat")
    check(err)

    defer f.Close()

    f.WriteString("writes\n")

    f.Sync()

    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

}

