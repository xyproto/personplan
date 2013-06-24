package personplan

import (
	"fmt"
	"testing"
	"time"
)

// TODO: Actually check the return values
func Test(t *testing.T) {
	ppAlexander := NewPersonPlan("Alexander")
	ppAlexander.AddWorkday(time.Monday, 8, 15, "KNH")     // monday, from 8, up to 15
	ppAlexander.AddWorkday(time.Wednesday, 12, 17, "KOH") // wednesday, from 12, up to 17

	fmt.Println(ppAlexander.String())

	ppBob := NewPersonPlan("Bob")
	ppBob.AddWorkday(time.Monday, 9, 11, "KOH")   // monday, from 9, up to 11
	ppBob.AddWorkday(time.Thursday, 8, 10, "KNH") // wednesday, from 8, up to 10

	fmt.Println(ppBob.String())

	periodplan := NewPeriodPlan(2013, 1, 8)
	periodplan.AddPersonPlan(ppAlexander)
	periodplan.AddPersonPlan(ppBob)

	fmt.Println(periodplan.String())

	allPlans := NewPlans()
	allPlans.AddPeriodPlan(periodplan)

	fmt.Println("Info for all plans:")
	date := time.Date(2013, 3, 4, 10, 32, 0, 0, time.UTC)

	for i, pp := range allPlans.all {
		fmt.Printf("Plan %d\n", i)
		fmt.Println(pp.ViewDay(date))
	}

	allPlans.HourInfo(date)

	allPlans.HourInfo(time.Date(2013, 3, 7, 9, 14, 0, 0, time.UTC))
}
