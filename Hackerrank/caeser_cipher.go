package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
    var i int
    var ans string
    ans=""
    var temp byte
    for i=0;i<len(s);i++{
        if s[i]>='a' && s[i]<='z'{
            temp=byte((((s[i])-byte(97))+byte(k))%byte(26)+byte(97))
            ans+=string(temp)
        }else if s[i]>='A' && s[i]<='Z'{
            temp=byte((((s[i])-byte(65))+byte(k))%byte(26)+byte(65))
            ans+=string(temp)
        }else{
            ans+=string(s[i])
        }
    }
    return ans
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)
    n++
    n--
    s := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)

    fmt.Fprintf(writer, "%s\n", result)

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
