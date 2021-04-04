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
		for {
			line := fmt.Sprintf("%d", happy)
			fmt.Println("forgiveness", line, AsciiByteToBase9(line))
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
