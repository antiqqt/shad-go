//go:build !solution

// Package speller implements a simple number speller
package speller

import (
	"fmt"
	"strings"
)

var numHmap = map[int64]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

type bigNum struct {
	value int64
	name  string
}

var bigNums = []bigNum{
	{value: 1000000000000000000, name: "quadrillion"},
	{value: 1000000000000, name: "trillion"},
	{value: 1000000000, name: "billion"},
	{value: 1000000, name: "million"},
	{value: 1000, name: "thousand"},
}

func Spell(n int64) string {
	var numStrs []string
	curr := n

	if curr == 0 {
		return convertUnderThousand(curr)
	}

	if curr < 0 {
		curr = curr * -1
		numStrs = append(numStrs, "minus")
	}

	for _, bn := range bigNums {
		if curr < bn.value {
			continue
		}

		units := curr / bn.value
		numStrs = append(numStrs, convertUnderThousand(units), bn.name)

		curr = curr % bn.value
	}

	if curr > 0 {
		numStrs = append(numStrs, convertUnderThousand(curr))
	}

	return strings.Join(numStrs, " ")
}

func convertUnderThousand(num int64) string {
	if num > 999 {
		panic("num is greater than 999")
	}

	var numStrs []string
	curr := num

	if curr >= 100 {
		hundreds := curr / 100
		numStrs = append(numStrs, fmt.Sprintf("%v hundred", numHmap[hundreds]))
		curr = curr % 100
	}

	if curr >= 20 {
		tens := curr / 10
		digits := curr % 10

		if digits == 0 {
			numStrs = append(numStrs, numHmap[tens*10])
		} else {
			numStrs = append(numStrs, fmt.Sprintf("%v-%v", numHmap[tens*10], numHmap[digits]))
		}

		return strings.Join(numStrs, " ")
	}

	if curr > 0 {
		numStrs = append(numStrs, numHmap[curr])
	}

	if curr == 0 && len(numStrs) == 0 {
		numStrs = append(numStrs, numHmap[curr])
	}

	return strings.Join(numStrs, " ")
}

// 0 - 19 - put into map as is
// 20 - 100 - combine tens and units with dash
// 100 - 1000 - important to implement well, thousand is the smallest "unit" in this whole task
// 1000 - end - recursively subtract from current unit by unit(their count will be less than 1000) + big-unit-name
// negatives?
