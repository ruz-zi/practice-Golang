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
			_wt.WriteString(strconv.FormatFloat(f, 'f', 9, 64))
		case string:
			_wt.WriteString(f)
		}
	}
}

func main() {
	_sc.Split(bufio.ScanWords)
	defer _wt.Flush()

	var n int
	ishansu := func(x int) bool {
		if x == 1000 {
			return false
		}
		if x < 100 {
			return true
		}
		for d := -9; d < 10; d++ {
			if x/100+d == x/10%10 && x/100+d+d == x%10 {
				return true
			}
		}
		return false
	}
	scan(&n)
	var c int
	for i := 1; i <= n; i++ {
		if ishansu(i) {
			c++
		}
	}
	print(c)
}
