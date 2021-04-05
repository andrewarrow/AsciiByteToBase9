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
			//lastOfLine := line[len(line)-1:]

			prev = last
			last = fmt.Sprintf("%d", AsciiByteToBase9(line))

			if prev == "9" && last == "1" {
				toggle = !toggle
			}
			if last == "1" || last == "2" || last == "4" {
				fmt.Println("   -", last, "            ", line)
			} else if last == "5" || last == "7" || last == "8" {
				fmt.Println("    ", last, "            ", line)
			} else if last == "3" {
				fmt.Println(" -", line)
			} else if last == "6" {
				fmt.Println("  ", line)
			} else if last == "9" {
				fmt.Println(line)
			}
			happy++
			time.Sleep(time.Millisecond * 25)
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
func AsciiByteToBase9WithWork(a string) byte {

	sum := byte(0)
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)
		fmt.Printf("%d+", t)

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	if len(strSum) > 1 {
		return AsciiByteToBase9(strSum)
	}
	return sum

}
