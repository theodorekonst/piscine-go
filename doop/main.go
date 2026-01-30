package main

import "os"

const (
	MaxInt64 = int64(1<<63 - 1)
	MinInt64 = int64(-1 << 63)
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoiSafe(os.Args[1])
	b, ok2 := atoiSafe(os.Args[3])
	if !ok1 || !ok2 {
		return
	}
	op := os.Args[2]

	switch op {
	case "+":
		res := a + b
		// overflow if signs indicate wrap
		if (b > 0 && res < a) || (b < 0 && res > a) {
			return
		}
		printNbr(res)
	case "-":
		res := a - b
		if (b > 0 && res > a) || (b < 0 && res < a) {
			return
		}
		printNbr(res)
	case "*":
		if a == 0 || b == 0 {
			printNbr(0)
			return
		}
		res := a * b
		if res/b != a {
			return
		}
		printNbr(res)
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		printNbr(a / b)
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		printNbr(a % b)
	default:
		return
	}
}

func atoiSafe(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := int64(1)
	i := 0
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}
	if i == len(s) {
		return 0, false
	}

	var n int64
	if sign == 1 {
		for ; i < len(s); i++ {
			c := s[i]
			if c < '0' || c > '9' {
				return 0, false
			}
			d := int64(c - '0')
			if n > (MaxInt64-d)/10 {
				return 0, false
			}
			n = n*10 + d
		}
		return n, true
	}
	// build negative directly to avoid -MinInt64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int64(c - '0')
		if n < (MinInt64+d)/10 {
			return 0, false
		}
		n = n*10 - d
	}
	return n, true
}

func printNbr(n int64) {
	os.Stdout.Write([]byte(itoa(n)))
	os.Stdout.Write([]byte("\n"))
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
	}
	// Build digits without negating n (handles MinInt64)
	var buf [20]byte
	i := len(buf)
	x := n
	for x != 0 {
		var digit byte
		if x > 0 {
			digit = byte(x % 10)
			x /= 10
		} else {
			digit = byte(-(x % 10))
			x /= 10
		}
		i--
		buf[i] = '0' + digit
	}
	if sign != "" {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
