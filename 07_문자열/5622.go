package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	var ans string

	sc.Scan()
	s := sc.Text()
	sum := 0
	code := [26]int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9}
	for _, c := range s {
		sum += code[c-'A'] + 1
	}
	ans = strconv.FormatInt(int64(sum), 10)
	wr.WriteString(ans)
}
