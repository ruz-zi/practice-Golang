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
	n, _ := strconv.Atoi(sc.Text())
	c := 0
	is_gword := func(s string) bool {
		ab := [26]int{}
		for i, c := range s {
			if i > 0 && s[i] != s[i-1] && ab[c-'a'] != 0 {
				return false
			}
			ab[c-'a']++
		}
		return true
	}
	for n > 0 {
		n--
		sc.Scan()
		s := sc.Text()
		if is_gword(s) {
			c++
		}
	}
	ans = strconv.Itoa(c)
	wr.WriteString(ans)
}
