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
			g := ""

			if prev == "9" && last == "1" {
				toggle = !toggle
			}
			if last == "1" || last == "2" || last == "4" {
				g = " m"
			} else if last == "5" || last == "7" || last == "8" {
				g = " f"
			} else if last == "3" || last == "6" {
				g = " ."
			} else {
				g = ".."
			}
			if toggle {
				fmt.Println("forgiveness", line, "-"+last, g)
			} else {
				fmt.Println("forgiveness", line, " "+last, g)
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
