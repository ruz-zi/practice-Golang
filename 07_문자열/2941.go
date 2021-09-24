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
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'c' && i+1 < len(s) && (s[i+1] == '=' || s[i+1] == '-') {
			i++
		} else if s[i] == 'd' {
			if i+2 < len(s) && s[i+1] == 'z' && s[i+2] == '=' {
				i += 2
			} else if i+1 < len(s) && s[i+1] == '-' {
				i++
			}
		} else if s[i] == 'l' && i+1 < len(s) && s[i+1] == 'j' {
			i++
		} else if s[i] == 'n' && i+1 < len(s) && s[i+1] == 'j' {
			i++
		} else if s[i] == 's' && i+1 < len(s) && s[i+1] == '=' {
			i++
		} else if s[i] == 'z' && i+1 < len(s) && s[i+1] == '=' {
			i++
		}
		c++
	}
	ans = strconv.FormatInt(int64(c), 10)
	wr.WriteString(ans)
}
