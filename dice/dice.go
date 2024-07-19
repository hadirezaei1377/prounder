package dice

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	const rolls = 10000
	counts := make(map[int]int)

	for i := 0; i < rolls; i++ {
		num := rand.Intn(6) + 1
		counts[num]++
	}

	fmt.Printf("تعداد دفعات تکرار هر عدد در %d بار پرتاب تاس:\n", rolls)
	for i := 1; i <= 6; i++ {
		fmt.Printf("عدد %d: %d بار\n", i, counts[i])
	}
}
