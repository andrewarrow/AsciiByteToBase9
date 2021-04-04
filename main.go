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
		otherToggleCount := 0
		otherToggle := true
		last3 := ""
		last6 := ""
		lastAnswer := 0
		prevAnswer := 0
		buff := []string{}
		for {
			line := fmt.Sprintf("%d", happy)
			lastOfLine := line[len(line)-1:]

			prev = last
			last = fmt.Sprintf("%d", AsciiByteToBase9WithWork(line))
			g := ""

			if prev == "9" && last == "1" {
				toggle = !toggle
			}
			if last == "1" || last == "2" || last == "4" {
				g = " "
			} else if last == "5" || last == "7" || last == "8" {
				g = " "
			} else if last == "3" {
				last3 = lastOfLine
				l3, _ := strconv.Atoi(last3)
				l6, _ := strconv.Atoi(last6)
				if prevAnswer == 6 || lastAnswer == 13 {
				}
				prevAnswer = lastAnswer
				lastAnswer = l3 + l6
				g = last3 + last6 + "=" + fmt.Sprintf("%d", lastAnswer)
			} else if last == "6" {
				last6 = lastOfLine
				l3, _ := strconv.Atoi(last3)
				l6, _ := strconv.Atoi(last6)
				if prevAnswer == 3 || lastAnswer == 15 {
				}
				prevAnswer = lastAnswer
				lastAnswer = l3 + l6
				g = last3 + last6 + "=" + fmt.Sprintf("%d", lastAnswer)
			} else if last == "9" {
				otherToggleCount += 1
				if otherToggleCount == 10 {
					otherToggleCount = 0
					otherToggle = !otherToggle
					buff = append(buff, line)
				}
				if otherToggle == false && lastOfLine != "0" {
					g = "   -" + lastOfLine
				} else {
					g = "    " + lastOfLine
				}

				g = g + "   " + deltas(buff)
			}
			if toggle {
				fmt.Println("forgiveness", line, "-"+last, g)
			} else {
				fmt.Println("forgiveness", line, " "+last, g)
			}
			happy++
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		time.Sleep(time.Second)
	}

}

func deltas(buff []string) string {
	i := 0
	prev := 0

	for _, s := range buff {

		prev = i
		i, _ = strconv.Atoi(s)
		//fmt.Println(i, prev, i-prev)
	}

	return fmt.Sprintf("%d,%d,%d", i-prev, i, prev)
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
