package main

import (
	"bufio"
	"os"
	"strconv"
)

// fastIO
var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	var ans string

	sc.Scan()
	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		}
		n, _ := strconv.ParseInt(sc.Text(), 10, 32)
		sc.Scan()
		s := []byte(sc.Text())
		for _, i := range s {
			for j := int64(0); j < n; j++ {
				ans += string(i)
			}
		}
		ans += "\n"
	}
	wr.WriteString(ans)
}
