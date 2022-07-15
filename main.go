package main

import "fmt"

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

func main() {
	fmt.Println("wdotm")
	printWeekdaysOnly(calculate(FRI, 31))
}

func calculate(day Days, totalDay int) map[Days]int {
	ret := make(map[Days]int)
	for i := 0; i < totalDay; i++ {
		ret[(Days(i)+day)%7]++
	}
	return ret
}

func printWeekdaysOnly(data map[Days]int) {
	for i := MON; i < SAT; i++ {
		fmt.Printf("%10s:%3d\n", i, data[i])
	}
}
