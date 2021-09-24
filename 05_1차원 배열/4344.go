package main

import (
	"bufio"
	"os"
	"strconv"
)

// fastIO
var _sc = bufio.NewScanner(os.Stdin)
var _wt = bufio.NewWriter(os.Stdout)

func scan(x ...interface{}) bool {
	for _, v := range x {
		if !_sc.Scan() {
			return false
		}
		switch f := v.(type) {
		case *int:
			i, _ := strconv.ParseInt(_sc.Text(), 10, strconv.IntSize)
			*f = int(i)
		case *int64:
			i, _ := strconv.ParseInt(_sc.Text(), 10, 64)
			*f = i
		case *float64:
			i, _ := strconv.ParseFloat(_sc.Text(), 64)
			*f = i
		case *string:
			*f = _sc.Text()
		}
	}
	return true
}

func print(x ...interface{}) {
	for _, v := range x {
		switch f := v.(type) {
		case int:
			_wt.WriteString(strconv.FormatInt(int64(f), 10))
		case int64:
			_wt.WriteString(strconv.FormatInt(f, 10))
		case float64:
			_wt.WriteString(strconv.FormatFloat(f, 'f', 3, 64))
		case string:
			_wt.WriteString(f)
		}
	}
}

func main() {
	_sc.Split(bufio.ScanWords)
	defer _wt.Flush()

	var t, n, s, c int
	scan(&t)
	for t > 0 {
		t--
		scan(&n)
		s = 0
		c = 0
		a := make([]int, n)
		for i := range a {
			scan(&a[i])
			s += a[i]
		}
		avg := float64(s) / float64(n)

		for i := range a {
			if float64(a[i]) > avg {
				c++
			}
		}
		print(float64(c*100)/float64(n), "%\n")
	}
}
