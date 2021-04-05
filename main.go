package main

import (
	"fmt"
	"strconv"
	"time"
)

type Base9 struct {
	Digit0 int // 9^0=0
	Digit1 int // 9^1=9
	Digit2 int // 9^2=18
	Digit3 int // 9^3=729
	Digit4 int // 9^4=6561
}

func (b *Base9) Add(a int) {
	b.Digit0++
	if b.Digit0 == 9 {
		b.Digit0 = 0
		b.Digit1++
	}
}
func (b *Base9) Debug() string {
	val := (b.Digit4 * 6561) + (b.Digit3 * 729) + (b.Digit2 * 18) + (b.Digit1 * 9) + b.Digit0
	return fmt.Sprintf("%d %d %d %d %d = %d", b.Digit4, b.Digit3, b.Digit2, b.Digit1, b.Digit0, val)
}

func main() {
	b := Base9{}
	go func() {
		for {
			b.Add(1)
			fmt.Println(b.Debug())
			time.Sleep(time.Millisecond * 100)
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
