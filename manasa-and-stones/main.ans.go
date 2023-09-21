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
 * Complete the 'stones' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER a
 *  3. INTEGER b
 */

func stones(n int32, a int32, b int32) []int32 {
	var possibleLastVal []int32
	var lower int32
	var upper int32
	if a >= b {
		lower = b
		upper = a
	} else {
		lower = a
		upper = b
	}
	for i := int32(0); i < n; i++ {

		possibleLastVal = append(possibleLastVal, upper*i+(lower*(n-1-i)))
	}
	fmt.Println(possibleLastVal)
	var uniqueArr []int32
	for j := 0; j < len(possibleLastVal); j++ {
		if !contains(uniqueArr, possibleLastVal[j]) {
			uniqueArr = append(uniqueArr, possibleLastVal[j])
		}
	}
	return uniqueArr

}

func contains(s []int32, num int32) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		aTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := stones(n, a, b)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
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
