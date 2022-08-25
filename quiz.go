package greeting

import (
	"time"
)

/*
Problem 1:
         Define a function named Welcome that takes no arguments,
         and returns a string "Welcome to Community Carnival".
*/

func Welcome() string {
	return ""
}

/*
Problem 2:
         Define a function that takes your joining date
         and returns "Part of big celebration" if you joined before IPO date(16th September 2021)
         else it returns "Missed the big celebration".
*/

func Celebration(joiningDate time.Time) string {
	// IpoDate := time.Date(2021, time.September, 16, 0, 0, 0, 0, time.UTC)
	return ""
}

/*
Problem 3:
         Find number of visitors visited Go community stalls
*/

func NumberOfVisitors(time int) int {
	//visitorsPerMinutes := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	return 0
}

/*
Problem 4:
         Find the number of thoughtworkers joined after the year you joined
*/

func NumberOfThoughtworker(joiningDate time.Time) int {
	//twkersJoinedByYear := map[int]int{
	//	2018: 1000, 2019: 1100, 2020: 1200, 2021: 1400, 2022: 800,
	//}
	return 0
}

/*
Problem 5:
         Define a Employee struct with the following int type fields: id, casualLeave, annualLeave
         a> Allow creating a new employee by defaulting the casualLeave and annualLeave to 6 and 18 respectively
         b> Allow employees to take casualLeave and annualLeave
         c> Return json response for LeaveBalance
*/

type Employee struct {
	id          int
	casualLeave int
	annualLeave int
}

func (employee Employee) TakeAnnualLeave() {}

func (employee Employee) TakeCasualLeave() {}

func (employee Employee) LeaveBalance() map[string]int {
	return map[string]int{}
}

/*
Problem 6 [Bonus Question - No test]:
		Given a list of employees working in project find the total profession services fess charged to a client for the project.
		In a project there are two categories of contract -
           a. Time & material - Professional fee is calculated based on number of days worked (Rate per day * No of days)
           b. Fixed bid - Professional fee is same as total charge

		Here is the list of employees working on a project.
		Employee ID		Contract type			Total charge(USD)		Rate per day(USD)		No of days
		1				Time & Material									100						10
		2				Time & Material									150						5
		3				Fixed Bid				1500
		4				Fixed Bid				2000

		Example calculation
			For employee 1: 100 * 10 = 1,000 USD
			For employee 2: 150 * 5  = 750 USD
			For employee 3: 1,500 USD
			For employee 4: 2,000 USD
		------------------------------------
					Total: 5,250 USD
*/
