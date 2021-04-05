package main

import (
	"fmt"
	"strconv"
	"time"
)

type Base9 struct {
	Digit0 byte // 9^0=0
	Digit1 byte // 9^1=9
	Digit2 byte // 9^2=18
	Digit3 byte // 9^3=729
	Digit4 byte // 9^4=6561
}

func (b *Base9) Add(a int) {
}
func (b *Base9) Debug() string {
	return "hi"
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
