package main

import (
	"fmt"
)

const (
	MON Days = iota
	TUE
	WED
	THU
	FRI
	SAT
	SUN
)

type Days int

func (d Days) String() string {
	W := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return W[d]
}

func (d Days) Less(i, j Days) bool {
	return i < j
}

func Calculate(day Days, totalDay int) map[Days]int {
	ret := make(map[Days]int)
	for i := 0; i < totalDay; i++ {
		ret[(Days(i)+day)%7]++
	}
	return ret
}

func PrintWeekdaysOnly(data map[Days]int) {
	printDays(data, MON, FRI)
}

func PrintAllDays(data map[Days]int) {
	printDays(data, MON, SUN)
}

func printDays(data map[Days]int, start, end Days) {
	for i := start; i <= end; i++ {
		fmt.Printf("%10s:%3d\n", i, data[i])
	}
}
