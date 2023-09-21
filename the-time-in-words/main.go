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
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

var timeStrings = []string{"", "one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eightteen",
	"nineteen",
	"twenty",
	"twenty one",
	"twenty two",
	"twenty three",
	"twenty four",
	"twenty five",
	"twenty six",
	"twenty seven",
	"twenty eight",
	"twenty nine",
	"thirteen",
	"thirteen one",
	"thirteen two",
	"thirteen three",
	"thirteen four",
	"thirteen five",
	"thirteen six",
	"thirteen seven",
	"thirteen eight",
	"thirteen nine",
	"forty",
	"forty one",
	"forty two",
	"forty three",
	"forty four",
	"forty five",
	"forty six",
	"forty seven",
	"forty eight",
	"forty nine",
	"fifty",
	"fifty one",
	"fifty two",
	"fifty three",
	"fifty four",
	"fifty five",
	"fifty six",
	"fifty seven",
	"fifty eight",
	"fifty nine",
}

func timeInWords(h int32, m int32) string {

	// Write your code here
	if m == 00 {
		return fmt.Sprintf("%s o' clock", timeStrings[h])
	}
	if m <= 30 {
		if m == 1 {
			return fmt.Sprintf("one minute past %s", timeStrings[h])
		}
		if m == 30 {
			return fmt.Sprintf("half past %s", timeStrings[h])
		}
		if m == 15 {
			return fmt.Sprintf("quarter past %s", timeStrings[h])
		}
		return fmt.Sprintf("%s minutes past %s", timeStrings[m], timeStrings[h])
	} else {
		if m == 59 {
			return fmt.Sprintf("one minute to %s", timeStrings[h])
		}
		if m == 45 {
			return fmt.Sprintf("quarter to %s", timeStrings[h+1])
		}
		return fmt.Sprintf("%s minutes to %s", timeStrings[60-m], timeStrings[h+1])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
