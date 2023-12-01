package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func reverse(source string) string {
	var result string
	for _, v := range source {
		result = string(v) + result
	}
	return result
}

func cleanCalibration(source string) string {
	r1 := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	r2 := regexp.MustCompile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]") // Yes it's dirty
	var number []string
	var last string
	first := r1.FindAllString(source, 1)[0]
	if len(r2.FindAllString(source, 1)) > 0 {
		last = r2.FindAllString(reverse(source), 1)[0]
	} else {
		last = first
	}

	number = append(number, first)
	number = append(number, last)

	var dest string
	for i, v := range number {
		if v == "one" || v == reverse("one") {
			number[i] = "1"
		} else if v == "two" || v == reverse("two") {
			number[i] = "2"
		} else if v == "three" || v == reverse("three") {
			number[i] = "3"
		} else if v == "four" || v == reverse("four") {
			number[i] = "4"
		} else if v == "five" || v == reverse("five") {
			number[i] = "5"
		} else if v == "six" || v == reverse("six") {
			number[i] = "6"
		} else if v == "seven" || v == reverse("seven") {
			number[i] = "7"
		} else if v == "eight" || v == reverse("eight") {
			number[i] = "8"
		} else if v == "nine" || v == reverse("nine") {
			number[i] = "9"
		}
	}
	fmt.Println(number)
	dest = number[0] + number[len(number)-1]
	return dest
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Panicln(err)
	}
	fScan := bufio.NewScanner(f)
	var result int
	var i int
	for fScan.Scan() {
		i, err = strconv.Atoi(cleanCalibration(fScan.Text()))
		if err != nil {
			log.Panicln(err)
		}
		result += i
	}
	fmt.Println(result)
}
