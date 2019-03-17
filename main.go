package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	RamUsage()
}

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