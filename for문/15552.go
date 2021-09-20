package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1<<22)
	w := bufio.NewWriterSize(os.Stdout, 1<<22)
	r.ReadLine()
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		sum := 0
		tmp := 0
		for _, c := range line {
			if c == ' ' {
				sum = tmp
				tmp = 0
			} else {
				tmp = tmp*10 + int(c-'0')
			}
		}
		w.WriteString(strconv.Itoa(sum + tmp))
		w.WriteByte('\n')
	}
	w.Flush()
}
