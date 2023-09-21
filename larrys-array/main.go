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
 * Complete the 'larrysArray' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY A as parameter.
 */

func larrysArray(A []int32) string {
	// Write your code here
	lengthA := len(A)

	for i := 0; i < lengthA-2; i++ {
		num := i + 1
		numIndex := -1
		for j := i; j < lengthA; j++ {
			if A[j] == int32(num) {
				numIndex = j
				break
			}
		}

		if i == numIndex {
			continue
		}

		for {
			if numIndex < i+3 {
				temp := A[i]
				A[i] = A[i+1]
				A[i+1] = A[i+2]
				A[i+2] = temp
				numIndex--

				if numIndex == i {
					break
				}
			} else {
				start := numIndex - 2
				temp := A[start]
				A[start] = A[start+1]
				A[start+1] = A[start+2]
				A[start+2] = temp
				numIndex--
			}
		}

	}

	if A[lengthA-1]-A[lengthA-2] == A[lengthA-2]-A[lengthA-3] {
		return "YES"
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		result := larrysArray(A)

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
