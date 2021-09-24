package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"unicode"
)

func main() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	c := 0
	s, _ := ioutil.ReadAll(os.Stdin)
	if !unicode.IsSpace(rune(s[0])) {
		c = 1
	}
	for i := 1; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i-1])) && !unicode.IsSpace(rune(s[i])) {
			c++
		}
	}
	wr.WriteString(strconv.FormatInt(int64(c), 10))
}
