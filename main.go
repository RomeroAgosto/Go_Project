package main

import (
	//"runtime"
	"fmt"
	"time"
    "math/rand"
    "os"
    "log"
    "bufio"
)

const samp_len = 4      // "Sensor" Sample size
const fileName = "dat1.txt"

func main() {

	//RamUsage()
    initializeFile()
    
    go generator()
    reader := bufio.NewReader(os.Stdin)
    cmdString, err := reader.ReadString('\n')
    check(err)

    fmt.Println(cmdString)    
}

//Function that runs in "background" generating samples each second
func generator() {

    i := 10
    for i>1 {                           //Infinite loop
        var sample [samp_len] int       //Array to save vaules

        sample = getValues()            //Getting sample values
        writeFile(sample)               //Writing samples to file
        time.Sleep(time.Second)         //Waiting one second
    }
}

//Function that writes to the File that will be used to save values
func writeFile(w [samp_len]int) {

    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Opens file with permission and appends new values
    check(err)
    defer f.Close()

    for _, value := range w {
        _, err = f.WriteString(fmt.Sprintf("%d\n", value))
       check(err)

    }

}

//Function that initializes file
func initializeFile() {
    f, err := os.Create(fileName)
    check(err)

    defer f.Close()
   
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
