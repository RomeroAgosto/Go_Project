package main

import (
	"runtime"
	"fmt"
	"time"
    "math/rand"
)

func main() {

	//RamUsage()

    var value int
    value = getValues()
    fmt.Println(value)
}
/*
//Function that gathers the Ram Usage
func RamUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v ", m.Alloc)
        fmt.Printf("\tTotalAlloc = %v ", m.TotalAlloc)
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
*/

//Function that generates random integer values
func getValues() int {

    rand.Seed(time.Now().UnixNano())   //Using current time in nanosenconds as a seed so it changes everytime
    return rand.Int()                  //Generating random value
}
