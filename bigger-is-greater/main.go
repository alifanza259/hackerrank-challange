package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {
    // Write your code here
    wArr := strings.Split(w, "")
    
    possible := false
    maxSwappedIndex := 0
    minSwappedIndex := 0
    out:
    for i:=len(wArr)-1;i>0;i-- {
        for j:=i-1;j>=0;j-- {
            if wArr[i] >  wArr[j] {

                maxSwappedIndex = j
                minSwappedIndex = i
                possible=true
                break out
            }
        }
    }
    if possible == false {
        return "no answer"
    }
    
    out2:
    for {
        for i:=minSwappedIndex-1;i>maxSwappedIndex+1; i--{
            for j:=i-1;j>maxSwappedIndex;j--{
                if wArr[i] > wArr[j]{
                maxSwappedIndex = j
                minSwappedIndex = i
                continue out2
                }
            }
        }
        break
    }
    
    temp  := wArr[minSwappedIndex]
    wArr[minSwappedIndex] = wArr[maxSwappedIndex]
    wArr[maxSwappedIndex]=temp
    
    for i:=len(wArr)-1;i>maxSwappedIndex;i--{
        for j := i-1;j>maxSwappedIndex;j--{
            if wArr[i] <  wArr[j] {
                temp  := wArr[i]
                wArr[i] = wArr[j]
                wArr[j]=temp
            }
        }
    }
    
    return strings.Join(wArr,"")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    T := int32(TTemp)

    for TItr := 0; TItr < int(T); TItr++ {
        w := readLine(reader)

        result := biggerIsGreater(w)

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
