package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	happy := 0
	go func() {
		// My Function is...

		// Forgiveness
		last := ""
		prev := ""
		toggle := false
		for {
			line := fmt.Sprintf("%d", happy)

			prev = last
			last = fmt.Sprintf("%d", AsciiByteToBase9(line))

			if prev == "9" && last == "1" {
				toggle = !toggle
			}
			if toggle {
				fmt.Println("forgiveness", line, "-"+last)
			} else {
				fmt.Println("forgiveness", line, last)
			}
			happy++
			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(time.Second)
	}

}

func AsciiByteToBase9(a string) byte {

	sum := byte(0)
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	if len(strSum) > 1 {
		return AsciiByteToBase9(strSum)
	}
	return sum

}
