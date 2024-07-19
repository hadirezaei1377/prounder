package gavsandoogh

import (
	"fmt"
)

func main() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						if (e+c == 14) && ((2*b - 1) == a) && (d == b+1) && (b+c == 10) && (a+b+c+d+e == 30) {
							fmt.Printf("sinas pin code: %d%d%d%d%d\n", a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}
