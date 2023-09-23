package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'absolutePermutation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 */

func absolutePermutation(n int32, k int32) []int32 {
	// Write your code here

	// Check if 2 provided argument is possible to have absolutePermutation
	possibleCheck := float64(n) / float64(k)
	if possibleCheck != math.Trunc(possibleCheck) {
		return []int32{-1}
	}
	if k != 0 && (n/k)%2 == 1 {
		return []int32{-1}
	}

	arr := make([]int32, n)
	if k == 0 {
		for i := 1; i <= int(n); i++ {
			arr[i-1] = int32(i)
		}
	} else {
		side := "left"
		cursor := 0
		counter := int32(0)
		for i := 1; i <= int(n); i++ {
			if side == "left" {
				a := (1 + k) + (int32(cursor) * 2 * k)
				addInSide := int32(i-1) % k

				arr[i-1] = a + addInSide
				counter++

				if counter == k {
					counter = int32(0)
					side = "right"
				}
			} else {
				a := (1 + k) + (int32(cursor) * 2 * k)
				subtractSide := k
				addInSide := int32(i-1) % k

				arr[i-1] = a - subtractSide + addInSide
				counter++
				if counter == k {
					counter = int32(0)
					cursor++
					side = "left"
				}
			}
		}
	}

	return arr
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
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		result := absolutePermutation(n, k)

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
