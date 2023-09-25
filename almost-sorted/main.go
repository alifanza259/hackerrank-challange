package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'almostSorted' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

type foo []int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func almostSorted(arr []int32) {
	// Write your code here
	arrCopy := make(foo, len(arr))
	copy(arrCopy, arr)

	sort.Sort(arrCopy)

	counter := 0
	diff := []int32{}
	diffValue := []int32{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arrCopy[i] {
			counter++

			diffValue = append(diffValue, int32(arr[i]))
			diff = append(diff, int32(i))
		}
	}

	if counter == 2 {
		fmt.Println("yes")
		min := math.Min(float64(diff[0]), float64(diff[1]))
		max := math.Max(float64(diff[0]), float64(diff[1]))
		fmt.Printf("swap %d %d", int(min)+1, int(max)+1)
		return
	} else if counter == 0 {
		fmt.Println("yes")
		return
	}

	for i := int32(1); i < int32(len(diff)); i++ {
		if (diff[i]-diff[i-1] != 1) || ((diffValue[i-1] - diffValue[i]) < 0) {
			if diff[i]-diff[i-1] == 2 {
				i++
				continue
			}
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
	fmt.Printf("reverse %d %d", diff[0]+1, diff[len(diff)-1]+1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	almostSorted(arr)
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
