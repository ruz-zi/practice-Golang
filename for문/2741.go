package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	n := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		w.WriteString(strconv.Itoa(i))
		w.WriteByte('\n')
	}
	w.Flush()
}
