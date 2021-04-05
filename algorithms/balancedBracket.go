package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the isBalanced function below.
func isBalanced(s string) string {

    yes, no := "YES", "NO"
    for len(s) > 0 {
        t := s
        s = strings.ReplaceAll(s, "()", "")
        s = strings.ReplaceAll(s, "{}", "")
        s = strings.ReplaceAll(s, "[]", "")

        if t == s {
            return no
        }
    }
    return yes

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

    stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
    

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024*1024)

    tTemp, _ := strconv.ParseInt(readLine(reader), 10, 64)
    
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        s := readLine(reader)

        result := isBalanced(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
