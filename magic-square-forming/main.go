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
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	cost := 0

	// Center must be 5
	if s[1][1] != 5 {
		diff := math.Abs(float64(s[1][1] - 5))
		cost += int(diff)
		s[1][1] = 5
	}

	// Corner must be 2,4,6,8. Diagonal must be 2-5-8 and 4-5-6
	case1 := handleCase(s, 2, 4, 6, 8, cost)
	case2 := handleCase(s, 8, 4, 6, 2, cost)
	case3 := handleCase(s, 2, 6, 4, 8, cost)
	case4 := handleCase(s, 8, 6, 4, 2, cost)
	case5 := handleCase(s, 4, 2, 8, 6, cost)
	case6 := handleCase(s, 6, 2, 8, 4, cost)
	case7 := handleCase(s, 4, 8, 2, 6, cost)
	case8 := handleCase(s, 6, 8, 2, 4, cost)

	minCost14 := math.Min(math.Min(case1, case2), math.Min(case3, case4))
	minCost58 := math.Min(math.Min(case5, case6), math.Min(case7, case8))

	return int32(math.Min(minCost14, minCost58))
}

func findCost(provided int32, needed int32) int {
	return int(math.Abs(float64(provided - needed)))
}

func handleCase(oriS [][]int32, tl int32, tr int32, bl int32, br int32, cost int) float64 {
	s := make([][]int32, len(oriS))
	copy(s, oriS)
	for i := range oriS {
		s[i] = make([]int32, len(oriS[i]))
		copy(s[i], oriS[i])
	}

	// if tl == 4 && tr == 2 && bl ==8 {
	// fmt.Println(cost)
	// fmt.Println(tl)
	// fmt.Println(s[0][0])
	//     fmt.Println(&s)
	//     fmt.Println("diatas")
	// }

	tlCost := findCost(s[0][0], tl)
	s[0][0] = tl
	cost += tlCost

	trCost := findCost(s[0][2], tr)
	s[0][2] = tr
	cost += trCost

	blCost := findCost(s[2][0], bl)
	s[2][0] = bl
	cost += blCost

	brCost := findCost(s[2][2], br)
	s[2][2] = br
	cost += brCost

	processLastStepCost(s, &tl, &tr, &bl, &br, &cost)
	return float64(cost)
}

func processLastStepCost(s [][]int32, tl *int32, tr *int32, bl *int32, br *int32, cost *int) {
	if *br == 2 {
		if *bl == 4 {
			diff := findCost(s[2][1], 9)
			diff2 := findCost(s[1][2], 7)
			diff3 := findCost(s[1][0], 3)
			diff4 := findCost(s[0][1], 1)
			*cost += diff + diff2 + diff3 + diff4
		} else {
			diff := findCost(s[2][1], 7)
			diff2 := findCost(s[1][2], 9)
			diff3 := findCost(s[1][0], 1)
			diff4 := findCost(s[0][1], 3)
			*cost += diff + diff2 + diff3 + diff4
		}
	}

	if *br == 4 {
		if *bl == 8 {
			diff := findCost(s[2][1], 3)
			diff2 := findCost(s[1][2], 9)
			diff3 := findCost(s[1][0], 1)
			diff4 := findCost(s[0][1], 7)
			*cost += diff + diff2 + diff3 + diff4
		} else {
			diff := findCost(s[2][1], 9)
			diff2 := findCost(s[1][2], 3)
			diff3 := findCost(s[1][0], 7)
			diff4 := findCost(s[0][1], 1)
			*cost += diff + diff2 + diff3 + diff4
		}
	}

	if *br == 6 {
		if *bl == 2 {
			diff := findCost(s[2][1], 7)
			diff2 := findCost(s[1][2], 1)
			diff3 := findCost(s[1][0], 9)
			diff4 := findCost(s[0][1], 3)
			*cost += diff + diff2 + diff3 + diff4
		} else {
			diff := findCost(s[2][1], 1)
			diff2 := findCost(s[1][2], 7)
			diff3 := findCost(s[1][0], 3)
			diff4 := findCost(s[0][1], 9)
			*cost += diff + diff2 + diff3 + diff4
		}
	}

	if *br == 8 {
		if *bl == 4 {
			diff := findCost(s[2][1], 3)
			diff2 := findCost(s[1][2], 1)
			diff3 := findCost(s[1][0], 9)
			diff4 := findCost(s[0][1], 7)
			*cost += diff + diff2 + diff3 + diff4
		} else {
			diff := findCost(s[2][1], 1)
			diff2 := findCost(s[1][2], 3)
			diff3 := findCost(s[1][0], 7)
			diff4 := findCost(s[0][1], 9)
			*cost += diff + diff2 + diff3 + diff4
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

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
