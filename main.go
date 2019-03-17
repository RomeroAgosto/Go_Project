package main

import (
	//"runtime"
	"fmt"
	"time"
    "math/rand"
    "os"
    "log"
)

const samp_len = 4      // "Sensor" Sample size
const fileName = "dat1.txt"

func main() {

	//RamUsage()
    var sample [samp_len] int

    initializeFile()
    f, err := os.Create(fileName)
    check(err)

    defer f.Close()

    i:=10

    for i>1 { 
        sample = getValues()
        writeFile(sample)
    }
    
}

//Function that writes to the File that will be used to save values
func writeFile(w [samp_len]int) {

    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
    check(err)
    defer f.Close()

    for _, value := range w {
        _, err = f.WriteString(fmt.Sprintf("%d\n", value))
       check(err)

    }

}

//Function that initializes file
func initializeFile() {

   
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
        tmp[i] = rand.Int()            //Generating random value
        //check(tmp[i])
    }
    return tmp
}

// Error handler
func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
