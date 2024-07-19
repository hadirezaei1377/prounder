package matnkhan

import (
	"fmt"
	"regexp"
	"strconv"
)

func sumNumbersInText(text string) int {

	re := regexp.MustCompile(`\d+`)

	numStrings := re.FindAllString(text, -1)

	sum := 0
	for _, numStr := range numStrings {
		num, _ := strconv.Atoi(numStr)
		sum += num
	}

	return sum
}

func main() {
	text := "سلام من هادی هستم 25 ساله از اصفهان من 2 سال است که برنامه نویسی کار میکنم و 7 سال پیش وارد دانشگاه شدم."

	total := sumNumbersInText(text)
	fmt.Printf("مجموع اعداد موجود در متن: %d\n", total)
}
