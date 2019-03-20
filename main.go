package main

import (
	"runtime"
	"fmt"
	"time"
    "math/rand"
    "os"
    "os/exec"
    "log"
    "bufio"
    "strings"
    "io/ioutil"
    "strconv"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
)

const samp_len = 4                          // "Sensor" Sample size
const fileName = "dat1.txt"                 // Name of the file with the samples
var flagGo bool = false                     // Free to go boolean
var reader = bufio.NewReader(os.Stdin)      // Standard input (keyboard) reader
var perCPU = true                           // Per CPU percentage

func main() {

	//RamUsage()
    initializeFile()
    
    go generator()
    userInterface ()
}

//Function that gathers the Ram Usage
func RamUsage() {
    v, _ := mem.VirtualMemory()
    fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
}

//Function that gathers CPU Usage
func CPUUsage() {
    usage, err := cpu.Percent(0, perCPU)
    for i, value := range usage {
        fmt.Printf("CPU%d: %f ", i, value)
    }
    check(err)
}
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}


//Function that manages the user interface
func userInterface() {

    var input string
    var argv []string
    for flagGo == false {}
    i := 10
    for i>1 {
        printMenu()
        input = readInput()
        argv = strings.Fields(input)
        if len(argv) < 1 {                      // To avoid the program to crash in case there isn't a first word
            argv = append(argv, "error")
        } 
        switch argv[0] {
        case "all" :
            allFunc(argv)
        case "some" :
            someFunc(argv)
        /*case "average" :
            averFunc(argv)
        */default:
            fmt.Println("Unexpected input")
        }
    }
}

//Function that handles the case to show all variables
func allFunc(argv [] string) {
    //
    fmt.Println("All function")
    fileHandle, _ := os.Open(fileName)
    defer fileHandle.Close()

    fileScanner := bufio.NewScanner(fileHandle)
    tmp, err := strconv.Atoi(argv[1])
    check(err)

    fmt.Printf("|")
    for i := 1; i <= samp_len; i++ {
        fmt.Printf("      Variable %d      |", i)
    }
    fmt.Println()
    for i := 1; i <= tmp; i++ {
        if !fileScanner.Scan() {
            fmt.Println("Not Enought Data to Show")
            break;
        }
        fmt.Printf("|")
        for j :=1; j <= samp_len; j++ {
            fmt.Printf("%22s|", fileScanner.Text())
        }
        fmt.Println()
    }
    customPause()
}

//Function that handles the case to show some variables
func someFunc(argv [] string) {

}

//Function with a custom pause that waits for an "enter to continue"
func customPause() {

    fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
//Function that reads the user input
func readInput() string {
    cmdString, err := reader.ReadString('\n')
    check(err)

    return cmdString
}
//Function that prints to the console the user Menu
func printMenu() {
    clearComand()
    fmt.Println("\nWELCOME TO THE SENSOR SIMULATOR\n")
    RamUsage()
    CPUUsage()
    fmt.Println("\n")
    fmt.Println("MAIN MENU\n")
    fmt.Println("To get the N metrics for all variables write all followed by the value of N \n(Example: all 5)\n")
    fmt.Println("To get the N metrics for one or more variables write some\nfollowed by what variables 1-4 with commas in between and followed by the value of N \n(Example: some 1,3 5)\n")
    fmt.Println("To get the average for one or more variables write average\nfollowed by what variables 1-4 with commas in between \n(Example: average 1,3)\n")
    fmt.Printf("\n->")
}

//Function to clear Command line
var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows  
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func clearComand() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

//Function that runs in "background" generating samples each second
func generator() {

    flagGo = true
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

    var s string = ""
    var tmp string

    data, err := ioutil.ReadFile(fileName)
    check(err)
    tmp = string(data)

    for _, value := range w {
        s += fmt.Sprintf("%d\n", value)
    }

    tmp = s + tmp

    f, err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY, 0755) // Opens file with permission and appends new values
    check(err)
    defer f.Close()

    _, err = f.WriteString(tmp)
    check(err)



}

//Function that initializes file
func initializeFile() {
    f, err := os.Create(fileName)
    check(err)

    defer f.Close()
   
}

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
