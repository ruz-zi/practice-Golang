package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	var ans string

	sc.Scan()
	a := []byte(sc.Text())
	sc.Scan()
	b := []byte(sc.Text())
	reverse := func(s []byte) {
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		}
	}
	reverse(a)
	reverse(b)
	c := strings.Compare(string(a), string(b))
	if c < 0 {
		ans = string(b)
	} else {
		ans = string(a)
	}
	wr.WriteString(ans)
}
