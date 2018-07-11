package main

import (
    "fmt"
    "gitee.com/johng/gf/g/os/gtime"
)

func main() {
    formats := []string{
        "2006-01-02 15:04:05.000",
        "Mon Jan _2 15:04:05 MST 2006",
        "Time is: 03:04:05 PM",
        "2006-01-02T15:04:05.000000000Z07:00 MST",
    }
    t := gtime.Now()
    for _, f := range formats {
        fmt.Println(f)
        fmt.Println(t.Layout(f))
        fmt.Println()
    }
}
