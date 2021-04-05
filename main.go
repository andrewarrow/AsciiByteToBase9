package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Base9 struct {
	Digit0 int // 9^0=0
	Digit1 int // 9^1=9
	Digit2 int // 9^2=81
	Digit3 int // 9^3=729
	Digit4 int // 9^4=6561
}
type Base3 struct {
	Digit0 int // 3^0=0
	Digit1 int // 3^1=3
	Digit2 int // 3^2=27
	Digit3 int // 3^3=81
	Digit4 int // 3^4=243
	Digit5 int // 3^4=729
}

func (b *Base9) Add(a int) {
	b.Digit0++
	if b.Digit0 == 9 {
		b.Digit0 = 0
		b.Digit1++
	}
	if b.Digit1 == 9 {
		b.Digit1 = 0
		b.Digit2++
	}
	if b.Digit2 == 9 {
		b.Digit2 = 0
		b.Digit3++
	}
	if b.Digit3 == 9 {
		b.Digit3 = 0
		b.Digit4++
	}
}
func (b *Base3) Add(a int) {
	b.Digit0++
	if b.Digit0 == 3 {
		b.Digit0 = 0
		b.Digit1++
	}
	if b.Digit1 == 3 {
		b.Digit1 = 0
		b.Digit2++
	}
	if b.Digit2 == 3 {
		b.Digit2 = 0
		b.Digit3++
	}
	if b.Digit3 == 3 {
		b.Digit3 = 0
		b.Digit4++
	}
	if b.Digit4 == 3 {
		b.Digit4 = 0
		b.Digit5++
	}
	if b.Digit5 == 3 {
		os.Exit(1)
	}
}
func (b *Base9) Debug() string {
	val := (b.Digit4 * 6561) + (b.Digit3 * 729) + (b.Digit2 * 81) + (b.Digit1 * 9) + b.Digit0
	return fmt.Sprintf("%d %d %d %d %d = %d", b.Digit4, b.Digit3, b.Digit2, b.Digit1, b.Digit0, val)
}
func (b *Base3) Debug() string {
	val := (b.Digit5 * 729) + (b.Digit4 * 243) + (b.Digit3 * 81) + (b.Digit2 * 27) + (b.Digit1 * 3) + b.Digit0
	return fmt.Sprintf("%d %d %d %d %d %d = %d", b.Digit5, b.Digit4, b.Digit3, b.Digit2, b.Digit1, b.Digit0, val)
}

func main() {
	b := Base3{}
	go func() {
		for {
			fmt.Println(b.Debug())
			time.Sleep(time.Millisecond * 100)
			b.Add(1)
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
