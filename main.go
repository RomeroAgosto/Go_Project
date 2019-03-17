package main

import (
	//"runtime"
	"fmt"
	"time"
    "math/rand"
)

const samp_len = 4      // "Sensor" Sample size

func main() {

	//RamUsage()

    var sample [samp_len] int
    sample = getValues()
    for i := 0; i< len(sample); i++ {
        fmt.Println(sample[i])
    }
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
func getValues() [samp_len]int {

    var tmp [samp_len] int
    rand.Seed(time.Now().UnixNano())   //Using current time in nanosenconds as a seed so it changes everytime
    for i := 0; i< len(tmp); i++ {
        tmp[i] = rand.Int()                  //Generating random value
    }
    return tmp
}
