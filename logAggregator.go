package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "time"
    "io/ioutil"
    "encoding/json"
    "strings"
    "bufio"
)

type Processes struct {
    Processes []Process `json:"commands"`
}

type Process struct {
    Command string `json:"cmd"`
    Prefix string `json:"prefix"`
}

func run(prefix string, command string, channel chan string) {
    split := strings.Fields(command)
    cmd := exec.Command("")
    if(len(split) > 1) {
        cmd = exec.Command(split[0], split[1:len(split)]...)
    } else {
        cmd = exec.Command(command)
    }
    cmdReader, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(cmdReader)
    go func() {
        for scanner.Scan() {
            channel <- prefix + " " + scanner.Text()
        }
    }()
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }

    time.Sleep(2 * time.Second)
    go run (prefix, command, channel)
}

func out(channel chan string) {
    for {
        msg := <- channel
        fmt.Println(msg)
    }
}

func main() {
    var c chan string = make(chan string, 5)
    jsonFile, err := os.Open("logAggregator.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var processes Processes
    json.Unmarshal(byteValue, &processes)
    for i := 0; i < len(processes.Processes); i++ {
        go run(processes.Processes[i].Prefix, processes.Processes[i].Command, c)
        fmt.Println("Command " + processes.Processes[i].Command)
        fmt.Println("Prefix " + processes.Processes[i].Prefix)
    }

    go out(c)

    var input string
    fmt.Scanln(&input)
}
