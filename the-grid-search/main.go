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
 * Complete the 'gridSearch' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY G
 *  2. STRING_ARRAY P
 */

func gridSearch(G []string, P []string) string {
	// Write your code here
	requiredStrings := map[string][]struct {
		row   int
		index int
	}{}
	for i := range P {
		requiredStrings[P[i]] = []struct {
			row   int
			index int
		}{}
	}

	for i := 0; i < len(G); i++ {
		for j := 0; j < len(P); j++ {
			indices := getIndicesOf(P[j], G[i], true)
			if len(indices) != 0 {
				for _, val := range indices {
					requiredStrings[P[j]] = append(requiredStrings[P[j]], struct {
						row   int
						index int
					}{row: i, index: val})
				}
			}
		}
	}

	result := []string{}
	cursor := 0
	keyIndex := P[cursor]
	for i := 0; i < len(requiredStrings[keyIndex]); i++ {
		if cursor+1 == len(P) {
			result = append(result, "YES")
		} else {
			row := requiredStrings[keyIndex][i].row
			index := requiredStrings[keyIndex][i].index
			res := recursiveCheck(P, requiredStrings, cursor+1, false, row, index)
			result = append(result, res)
		}
	}

	for _, val := range result {
		if val == "YES" {
			return "YES"
		}
	}
	return "NO"
}

func recursiveCheck(P []string, requiredStrings map[string][]struct {
	row   int
	index int
}, cursor int, firstCursor bool, row int, index int) string {
	keyIndex := P[cursor]
	if len(requiredStrings[keyIndex]) != 0 {
		for i := 0; i < len(requiredStrings[keyIndex]); i++ {
			if firstCursor == false {
				curr := requiredStrings[keyIndex][i]
				if row+1 == curr.row && index == curr.index {
					if cursor+1 == len(P) {
						return "YES"
					} else {
						row := requiredStrings[keyIndex][i].row
						index := requiredStrings[keyIndex][i].index
						return recursiveCheck(P, requiredStrings, cursor+1, false, row, index)
					}
				}
			}
		}
	}
	return "NO"
}

func getIndicesOf(searchStr string, str string, caseSensitive bool) []int {
	searchStrLen := len(searchStr)
	if searchStrLen == 0 {
		return []int{}
	}
	startIndex := 0
	indices := []int{}
	if caseSensitive == false {
		searchStr = strings.ToLower(searchStr)
	}
	for {
		index := indexAt(str, searchStr, startIndex)
		if index == -1 {
			break
		}
		indices = append(indices, index)
		startIndex = index + 1
	}
	return indices
}

func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
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

		RTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		R := int32(RTemp)

		_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		// C := int32(CTemp)

		var G []string

		for i := 0; i < int(R); i++ {
			GItem := readLine(reader)
			G = append(G, GItem)
		}

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		rTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		r := int32(rTemp)

		_, err = strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		// c := int32(cTemp)

		var P []string

		for i := 0; i < int(r); i++ {
			PItem := readLine(reader)
			P = append(P, PItem)
		}

		result := gridSearch(G, P)

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
