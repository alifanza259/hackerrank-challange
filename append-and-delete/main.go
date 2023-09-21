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
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	current := strings.Split(s, "")
	desired := strings.Split(t, "")
	currentLength := len(current)
	desiredLength := len(desired)

	// Edge case
	edgeCaseStep := currentLength + desiredLength
	if k >= int32(edgeCaseStep) {
		return "Yes"
	}

	var longer string
	var n int
	if currentLength >= desiredLength {
		longer = "current"
		n = desiredLength
	}

	differences := 0
	for i := 0; i < int(n); i++ {
		char := current[i]
		if char != desired[i] {
			if longer == "current" {
				differences = (n - i) * 2
			} else {
				differences = n - i
			}
			break
		}
	}
	differences = differences + int(math.Abs(float64(currentLength)-float64(desiredLength)))

	if k >= int32(differences) {
		if longer != "current" {
			if differences%2 == 1 && k%2 == 1 {
				return "Yes"
			}
			if differences%2 == 0 && k%2 == 0 {
				return "Yes"
			}
		} else {
			return "Yes"
		}
	}

	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

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
