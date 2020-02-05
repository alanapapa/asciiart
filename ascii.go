package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, _ := os.Open("standard.txt")
    fileVal := ScanFile(file)
    arg := os.Args[1]
    for _, v := range os.Args[2:] {
        arg += " " + v
    }
    narg := strings.Split(arg, "\\n")
    for _, v := range narg {
        printLetter(string(v), fileVal)
    }
}

func printLetter(s string, fileVal []string) {
    for i := 1; i <= 8; i++ {
        for _, arg := range s {
            indexBase := int(rune(arg)-32) * 9
            fmt.Print(fileVal[indexBase+i])
        }
        fmt.Println()
    }
}

func ScanFile(arg *os.File) []string {
    var fileValue []string
    scanner := bufio.NewScanner(arg)
    for scanner.Scan() {
        lines := scanner.Text()
        fileValue = append(fileValue, lines)
    }
    return fileValue
}
