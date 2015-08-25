package main

import (
	"fmt"
)

var ones = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine"}

var teens = map[int]string{0: "ten", 1: "eleven", 2: "twelve", 3: "thirteen", 4: "fourteen", 5: "fifteen",
	6: "sixteen", 7: "seventeen", 8: "eighteen", 9: "nineteen"}

var tens = map[int]string{1: "ten", 2: "twenty", 3: "thirty", 4: "forty", 5: "fifty",
	6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}

var hundreds = map[int]string{1: "onehundred", 2: "twohundred", 3: "threehundred", 4: "fourhundred", 5: "fivehundred",
	6: "sixhundred", 7: "sevenhundred", 8: "eighthundred", 9: "ninehundred"}

var thousands = map[int]string{1: "onethousand"}

func inwords(num int) (s string) {

	if s = thousands[num/1000]; s == "" {

		s = hundreds[num/100]

		if num/100 != 0 && num%100 != 0 {
			s = s + "and"
		}

		if (num/10)%10 == 1 {
			s = s + teens[num%10]
		} else {
			s = s + tens[(num/10)%10] + ones[num%10]
		}
	}

	return s
}

func main() {

	text := ""
	for i := 1; i <= 1000; i++ {
		text = text + inwords(i)
	}

	fmt.Println("Length: ", len(text))
}
