package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUnitlCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUnitlCarRental = 25 - age
	if yearsUnitlCarRental < 0 {
		yearsUnitlCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDrinking, yearsUnitlCarRental

}

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUnitDrinking, yearsUnitlCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You are an adult in", yearsUnitDrinking, "years")
	fmt.Println("You can rent a car in", yearsUnitlCarRental, "years")
	fmt.Println("===========================")
}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}
