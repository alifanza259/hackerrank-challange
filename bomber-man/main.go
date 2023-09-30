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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	// Write your code here

	// Bruteforce way

	arr2D := [][]string{}
	for i := 0; i < len(grid); i++ {
		splittedString := strings.Split(grid[i], "")
		arr2D = append(arr2D, splittedString)
	}

	patterns := [][][]string{}
	answer := []string{}

	patternsOccurences := []int{}
	patternCounter := 0

	if n%2 == 0 {
		arr2D = processResult(arr2D, true)
	} else {
	out:
		for i := 1; i <= int(n); i++ {
			if i <= 3 {
				if i == 3 {
					arr2D = processResult(arr2D, false)
					patterns = append(patterns, arr2D)
				}
			} else {
				if i%2 == 1 {
					arr2D = processResult(arr2D, false)
					patterns = append(patterns, arr2D)
					if samePattern(patterns[0], arr2D) {
						patternsOccurences = append(patternsOccurences, i)
						patternCounter++
						if patternCounter == 2 {
							diff := -1
							for k := 1; k < len(patternsOccurences); k++ {
								diffSample := patternsOccurences[k] - patternsOccurences[k-1]
								if k == 1 {
									diff = diffSample
								} else {
									if diff != diffSample {
										patternCounter = 0
										continue out
									}
								}
							}
							nearestStart := n / int32(diff)
							indexPattern := math.Abs(float64(int(n)-(diff*int(nearestStart)+(patternsOccurences[0]%diff)))) / 2
							arr2D = patterns[int(indexPattern)]
							break
						}
					}
					continue
				}
			}

		}
	}

	for i := 0; i < len(arr2D); i++ {
		answer = append(answer, strings.Join(arr2D[i], ""))
	}
	return answer

}

func samePattern(firstPattern [][]string, currentPattern [][]string) bool {
	for i := 0; i < len(currentPattern); i++ {
		for j := 0; j < len(currentPattern[i]); j++ {
			if firstPattern[i][j] != currentPattern[i][j] {
				return false
			}
		}
	}
	return true
}

func processResult(arr2D [][]string, isEdgeCase bool) [][]string {
	s := make([][]string, len(arr2D))
	for i := 0; i < len(arr2D); i++ {
		s[i] = make([]string, len(arr2D[i]))
	}
	for i := 0; i < len(arr2D); i++ {
		for j := 0; j < len(arr2D[i]); j++ {
			s[i][j] = "O"
		}
	}
	if isEdgeCase {
		return s
	}

	for i := 0; i < len(arr2D); i++ {
		for j := 0; j < len(arr2D[i]); j++ {
			char := arr2D[i][j]
			if char == "O" {
				if i-1 >= 0 {
					s[i-1][j] = "."
				}
				if j-1 >= 0 {
					s[i][j-1] = "."
				}

				if i+1 <= len(s)-1 {
					if arr2D[i+1][j] != "O" {
						s[i+1][j] = "."
					}
				}
				if j+1 <= len(s[i])-1 {
					if arr2D[i][j+1] != "O" {
						s[i][j+1] = "."
					}
				}
				s[i][j] = "."
			}
		}
	}

	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	// cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
