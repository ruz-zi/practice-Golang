package main

import (
	"bufio"
	"os"
	"strconv"
)

// fastIO
var _sc = bufio.NewScanner(os.Stdin)
var _wt = bufio.NewWriter(os.Stdout)

func scan(x ...interface{}) {
	for _, v := range x {
		_sc.Scan()
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
}

func print(x ...interface{}) {
	for _, v := range x {
		switch f := v.(type) {
		case int:
			_wt.WriteString(strconv.FormatInt(int64(f), 10))
		case int64:
			_wt.WriteString(strconv.FormatInt(f, 10))
		case float64:
			_wt.WriteString(strconv.FormatFloat(f, 'g', 20, 64))
		case string:
			_wt.WriteString(f)
		}
	}
}

func main() {
	_sc.Split(bufio.ScanWords)
	defer _wt.Flush()

	var n, a, b int
	scan(&n)
	for i := 0; i < n; i++ {
		scan(&a, &b)
		print("Case #", i+1, ": ", a, " + ", b, " = ", a+b, "\n")
	}
}
