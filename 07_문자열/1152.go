package main

import (
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	sc.Buffer([]byte{}, 1000001)
	defer wr.Flush()

	sc.Scan()
	s := sc.Text()
	ab := [26]int{}
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			ab[c-'A']++
		} else if c >= 'a' && c <= 'z' {
			ab[c-'a']++
		}
	}
	mx := 0
	ans := '?'
	for i := 0; i < 26; i++ {
		if mx < ab[i] {
			mx, ans = ab[i], 'A'+rune(i)
		} else if mx == ab[i] {
			ans = '?'
		}
	}
	wr.WriteRune(ans)
}
